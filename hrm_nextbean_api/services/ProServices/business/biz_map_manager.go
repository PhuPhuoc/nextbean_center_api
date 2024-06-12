package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"

type mapPMStorage interface {
	MapPM(info *model.MapProPM) error
}

type mapPMBiz struct {
	store mapPMStorage
}

func NewMapPMBiz(store mapPMStorage) *mapPMBiz {
	return &mapPMBiz{store: store}
}

func (biz *mapPMBiz) MapPMBiz(info *model.MapProPM) error {
	if err_query := biz.store.MapPM(info); err_query != nil {
		return err_query
	}
	return nil
}
