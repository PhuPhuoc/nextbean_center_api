package repository

import (
	"fmt"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/intern_query"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"
)

func (store *internStore) MapInternSkill(info *model.MapInternSkill) error {
	if errInternIDNotExist := checkInternIDExist(store, info.InternId); errInternIDNotExist != nil {
		return errInternIDNotExist
	}

	if len(info.Skills) != len(info.SkillLevel) {
		return fmt.Errorf("skills and skill-level arrays must have the same length")
	}

	tx, err := store.db.Begin()
	if err != nil {
		return fmt.Errorf("error when MapInternSkill (start transaction) in store: %v", err)
	}

	deleteQuery := query.QueryDeleteMapInternSkill()
	if _, err := tx.Exec(deleteQuery, info.InternId); err != nil {
		tx.Rollback()
		return fmt.Errorf("error when MapInternSkill (delete mapping transaction) in store: %v", err)
	}

	values := ""
	for i := range info.Skills {
		if i > 0 {
			values += ","
		}
		values += fmt.Sprintf("('%s', %d, '%s')", info.InternId, info.Skills[i], info.SkillLevel[i])
	}

	updateQuery := query.QueryMapInternSkill(values)
	if _, err := tx.Exec(updateQuery); err != nil {
		tx.Rollback()
		return fmt.Errorf("error when MapInternSkill (update mapping transaction) in store: %v", err)

	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("error when MapInternSkill (commit transaction) in store: %v", err)

	}
	return nil
}

func checkInternIDExist(store *internStore, intern_id string) error {
	var flag bool
	query := query.QueryCheckExistInternID()
	err := store.db.QueryRow(query, intern_id).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error when check exist student-code : %v", err)
	}
	if !flag {
		return fmt.Errorf("student-code does not exist")
	}
	return nil
}
