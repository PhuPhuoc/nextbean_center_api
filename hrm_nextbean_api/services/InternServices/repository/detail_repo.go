package repository

import (
	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/intern_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"
)

func (store *internStore) GetDetailIntern(int_id string) (*model.InternDetailInfo, error) {
	if err_check_accID_exist := checkInternIDExist(store, int_id); err_check_accID_exist != nil {
		return nil, err_check_accID_exist
	}

	i_info := new(model.InternDetailInfo)

	infoQuery := query.QueryGetInternDetailInfo()
	skillQuery := query.QueryGetInternSkill()

	if err_query_info := store.db.QueryRow(infoQuery, int_id).Scan(&i_info.Id, &i_info.UserName, &i_info.Email, &i_info.StudentCode, &i_info.Avatar, &i_info.Gender, &i_info.DateOfBirth, &i_info.PhoneNumber, &i_info.Address); err_query_info != nil {
		return nil, err_query_info
	}

	rows, err_query := store.db.Query(skillQuery, i_info.Id)
	if err_query != nil {
		return nil, err_query
	}
	defer rows.Close()

	for rows.Next() {
		s := new(model.DetailSkill)
		if err_scan := rows.Scan(&s.TechnicalSkill, &s.SkillLevel); err_scan != nil {
			return nil, err_scan
		}
		i_info.InternSkill = append(i_info.InternSkill, *s)
	}

	return i_info, nil
}
