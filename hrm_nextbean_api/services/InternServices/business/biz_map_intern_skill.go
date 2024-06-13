package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"

type mapInternSkillStorage interface {
	MapInternSkill(internID string, info *model.MapInternSkill) error
}

type mapInternSkillBiz struct {
	store mapInternSkillStorage
}

func NewMapInternSkillBiz(store mapInternSkillStorage) *mapInternSkillBiz {
	return &mapInternSkillBiz{store: store}
}

func (biz *mapInternSkillBiz) MapInternSkillBiz(internID string, info *model.MapInternSkill) error {
	if err_query := biz.store.MapInternSkill(internID, info); err_query != nil {
		return err_query
	}
	return nil
}
