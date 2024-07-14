package model

type CommentCreation struct {
	Type    string `json:"type" validate:"required,type=enum(report or comment)"`
	Content string `json:"content" validate:"required,type=string"`
}
