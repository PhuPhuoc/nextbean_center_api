package business

type deleteOJTStorage interface {
	DeleteOJT(id string) error
}

type deleteOJTBiz struct {
	store deleteOJTStorage
}

func NewDeleteOJTBiz(store deleteOJTStorage) *deleteOJTBiz {
	return &deleteOJTBiz{store: store}
}

func (biz *deleteOJTBiz) DeleteOJTBiz(id string) error {
	if err_query := biz.store.DeleteOJT(id); err_query != nil {
		return err_query
	}
	return nil
}
