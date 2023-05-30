package mock

import (
	"context"

	dv1 "github.com/coderi421/gframework/app/user/srv/data/v1"
	metav1 "github.com/coderi421/gframework/pkg/common/meta/v1"
)

type users struct {
	//users []*dv1.UserDO   模拟数据库   （进行insert之后测试）
}

func NewUsers() *users {
	return &users{}
}
func (u *users) List(ctx context.Context, opts metav1.ListMeta) (*dv1.UserDOList, error) {
	//实现gorm查询
	return &dv1.UserDOList{
		TotalCount: 1,
		Items: []*dv1.UserDO{
			{"jzin"},
		},
	}, nil
}
