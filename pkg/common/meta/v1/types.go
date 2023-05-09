package v1

// ListMeta describes metadata that synthetic resources must have, including lists and
type ListMeta struct {
	Page int `json:"totalCount,omitempty"`

	PageSize int `json:"offset,omitempty" form:"offset"`
}
