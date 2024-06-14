package business

type RemoveMapProMemStorage interface {
	RemoveMemInPro(proid string, memid string) error
}

type removeMapProMemBiz struct {
	store RemoveMapProMemStorage
}

func NewRemoveProMemBiz(store RemoveMapProMemStorage) *removeMapProMemBiz {
	return &removeMapProMemBiz{store: store}
}

func (biz *removeMapProMemBiz) RemoveProMemBiz(proid string, memid string) error {
	if err_query := biz.store.RemoveMemInPro(proid, memid); err_query != nil {
		return err_query
	}
	return nil
}
