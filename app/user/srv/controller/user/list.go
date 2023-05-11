package user

import (
	"context"

	upbv1 "github.com/CoderI421/gframework/api/user/v1"
	srvv1 "github.com/CoderI421/gframework/app/user/srv/service/v1"
	metav1 "github.com/CoderI421/gframework/pkg/common/meta/v1"
)

func DTOToResponse(userdto srvv1.UserDTO) upbv1.UserInfoResponse {
	return upbv1.UserInfoResponse{}
}

/*
controller层依赖了service层，service层依赖了data层：
contoller层能否直接依赖data层：可以的
contoller依赖service层，并不是直接依赖了具体的struct 而是依赖了interface，
但是底层是绝对不能依赖父层的！
*/

func (us userServer) GetUserList(ctx context.Context, info *upbv1.PageInfo) (*upbv1.UserListResponse, error) {
	srvOpts := metav1.ListMeta{
		Page:     int(info.Pn),
		PageSize: int(info.PSize),
	}

	dtoList, err := us.srv.List(ctx, srvOpts)
	if err != nil {
		return nil, err
	}
	var rsp upbv1.UserListResponse
	for _, value := range dtoList.Items {
		userRsp := DTOToResponse(*value)
		rsp.Data = append(rsp.Data, &userRsp)
	}
	return &rsp, nil
}
