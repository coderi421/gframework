package v1

import (
	"context"

	metav1 "github.com/CoderI421/gmicro/pkg/common/meta/v1"
)

type UserDO struct {
	Name string
}
type UserDOList struct {
	TotalCount int64     `json:"totalCount,omitempty"`
	Items      []*UserDO `json:"data"`
}

type UserStore interface {
	List(ctx context.Context, opts metav1.ListMeta) (*UserDOList, error)
}
