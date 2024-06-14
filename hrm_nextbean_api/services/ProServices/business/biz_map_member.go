package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"

type mapProMemStorage interface {
	MapProMem(proid string, mapInfo *model.MapProMem) error
}

type mapProMemBiz struct {
	store mapProMemStorage
}

func NewMapProMemBiz(store mapProMemStorage) *mapProMemBiz {
	return &mapProMemBiz{store: store}
}

func (biz *mapProMemBiz) MapProMemBiz(proid string, info *model.MapProMem) error {
	if err_query := biz.store.MapProMem(proid, info); err_query != nil {
		return err_query
	}
	return nil
}
