package serverinterceptors

import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/CoderI421/gframework/pkg/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// UnaryTimeoutInterceptor returns a func that sets timeout to incoming unary requests.
func UnaryTimeoutInterceptor(timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		var resp interface{}
		var err error
		var lock sync.Mutex
		done := make(chan struct{})
		// create channel with buffer size 1 to avoid goroutine leak
		panicChan := make(chan interface{}, 1)
		go func() {
			defer func() {
				if p := recover(); p != nil {
					// attach call stack to avoid missing in different goroutine
					panicChan <- fmt.Sprintf("%+v\n\n%s", p, strings.TrimSpace(string(debug.Stack())))
				}
			}()

			lock.Lock()
			defer lock.Unlock()

			resp, err = handler(ctx, req)
			close(done) //正常结束
		}()

		select {
		case p := <-panicChan:
			panic(p)
		case <-done:
			lock.Lock()
			defer lock.Unlock()
			return resp, err
		//超时会拿到cancel 内部会ctx.Done()<-下面会接受到
		case <-ctx.Done():
			err := ctx.Err()
			if err == context.Canceled {
				//我们之前说过我们把error统一了， grpc的error我们也可以统一, 自己完成
				err = errors.WithCode(int(codes.Canceled), err.Error())
				//err = status.Error(codes.Canceled, err.Error())
			} else if err == context.DeadlineExceeded {
				err = errors.WithCode(int(codes.Canceled), err.Error())
				//err = status.Error(codes.DeadlineExceeded, err.Error())
			}
			return nil, err
		}
	}
}
