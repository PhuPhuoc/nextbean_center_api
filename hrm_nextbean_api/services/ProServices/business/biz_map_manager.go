package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"

type mapPMStorage interface {
	MapPM(proid string, info *model.MapProPM) error
}

type mapPMBiz struct {
	store mapPMStorage
}

func NewMapPMBiz(store mapPMStorage) *mapPMBiz {
	return &mapPMBiz{store: store}
}

func (biz *mapPMBiz) MapPMBiz(proid string, info *model.MapProPM) error {
	if err_query := biz.store.MapPM(proid, info); err_query != nil {
		return err_query
	}
	return nil
}
