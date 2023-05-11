package v1

import (
	"context"

	dv1 "github.com/CoderI421/gframework/app/user/srv/data/v1"
	metav1 "github.com/CoderI421/gframework/pkg/common/meta/v1"
)

type UserDTO struct {
	dv1.UserDO
}

type UserDTOList struct {
	TotalCount int64      `json:"totalCount,omitempty"`
	Items      []*UserDTO `json:"data"`
}

type UserSrv interface {
	List(ctx context.Context, opts metav1.ListMeta) (*UserDTOList, error)
}

type userService struct {
	userStore dv1.UserStore
}

var _ UserSrv = (*userService)(nil)

func NewUserService(userStore dv1.UserStore) *userService {
	return &userService{userStore: userStore}
}

func (u *userService) List(ctx context.Context, opts metav1.ListMeta) (*UserDTOList, error) {

	list, err := u.userStore.List(ctx, opts)
	if err != nil {
		return nil, err
	}

	var userDTOList UserDTOList
	for _, value := range list.Items {
		projectDTO := UserDTO{*value}
		userDTOList.Items = append(userDTOList.Items, &projectDTO)
	}

	return &userDTOList, nil
}
