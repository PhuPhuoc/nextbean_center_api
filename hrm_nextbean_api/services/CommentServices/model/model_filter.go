package model

type CommentFilter struct {
	AccID string `json:"-"`
	Type  string `json:"type,omitempty"`
}
