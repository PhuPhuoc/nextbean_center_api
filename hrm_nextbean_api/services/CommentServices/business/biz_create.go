package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/CommentServices/model"

type createCommentStorage interface {
	CreateComment(taskid, role, accID, internID string, info *model.CommentCreation) error
}

type createCommentBiz struct {
	store createCommentStorage
}

func NewCreateCommentBiz(store createCommentStorage) *createCommentBiz {
	return &createCommentBiz{store: store}
}

func (biz *createCommentBiz) CreateCommentBiz(taskid, role, accID, internID string, info *model.CommentCreation) error {
	if err_query := biz.store.CreateComment(taskid, role, accID, internID, info); err_query != nil {
		return err_query
	}
	return nil
}
