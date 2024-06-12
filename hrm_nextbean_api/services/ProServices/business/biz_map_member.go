package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"

type mapProMemStorage interface {
	MapProMem(mapInfo *model.MapProMem) error
}

type mapProMemBiz struct {
	store mapProMemStorage
}

func NewMapProMemBiz(store mapProMemStorage) *mapProMemBiz {
	return &mapProMemBiz{store: store}
}

func (biz *mapProMemBiz) MapProMemBiz(info *model.MapProMem) error {
	if err_query := biz.store.MapProMem(info); err_query != nil {
		return err_query
	}
	return nil
}
