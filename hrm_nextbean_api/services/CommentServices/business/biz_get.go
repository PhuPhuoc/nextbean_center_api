package business

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/CommentServices/model"
)

type getCommentStore interface {
	GetComment(ownerid string, pagin *common.Pagination, filter *model.CommentFilter) ([]model.Comment, error)
}

type getCommentBiz struct {
	store getCommentStore
}

func NewGetCommentBiz(store getCommentStore) *getCommentBiz {
	return &getCommentBiz{
		store: store,
	}
}

func (biz *getCommentBiz) GetCommentBiz(ownerid string, pagin *common.Pagination, filter *model.CommentFilter) ([]model.Comment, error) {
	data, err_query := biz.store.GetComment(ownerid, pagin, filter)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
