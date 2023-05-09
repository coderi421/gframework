package db

import (
	"context"

	"gorm.io/gorm"

	dv1 "github.com/CoderI421/gmicro/app/user/srv/data/v1"
	metav1 "github.com/CoderI421/gmicro/pkg/common/meta/v1"
)

type user struct {
	db *gorm.DB
}

var _ dv1.UserStore = (*user)(nil)

func NewUser(db *gorm.DB) *user {
	return &user{db: db}
}

func (u *user) List(ctx context.Context, opts metav1.ListMeta) (*dv1.UserDOList, error) {
	//实现gorm查询
	return nil, nil
}
