package model

type CommentFilter struct {
	TaskId string `json:"-"`
	AccId  string `json:"-"`
	Type   string `json:"type,omitempty"`
}
