package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"

type RemoveMapProMemStorage interface {
	RemoveMemInPro(mapInfo *model.MapProMem) error
}

type removeMapProMemBiz struct {
	store RemoveMapProMemStorage
}

func NewRemoveProMemBiz(store RemoveMapProMemStorage) *removeMapProMemBiz {
	return &removeMapProMemBiz{store: store}
}

func (biz *removeMapProMemBiz) RemoveProMemBiz(info *model.MapProMem) error {
	if err_query := biz.store.RemoveMemInPro(info); err_query != nil {
		return err_query
	}
	return nil
}
