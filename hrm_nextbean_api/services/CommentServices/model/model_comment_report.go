package model

type Comment struct {
	Id        string  `json:"id"`
	UserName  string  `json:"user-name"`
	Avatar    *string `json:"avatar,omitempty"`
	Type      string  `json:"type"`
	Content   string  `json:"content"`
	CreatedAt string  `json:"created-at"`
	IsOwner   bool    `json:"is-owner"`
}
