package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"
)

func (store *internStore) MapInternSkill(internID string, info *model.MapInternSkill) error {
	if errInternIDNotExist := checkInternIDExists(store, internID); errInternIDNotExist != nil {
		return errInternIDNotExist
	}

	if len(info.Skills) != len(info.SkillLevel) {
		return fmt.Errorf("invalid-request: skills and skill-level arrays must have the same length")
	}

	tx, err := store.db.Begin()
	if err != nil {
		return fmt.Errorf("error in MapInternSkill transaction: %v", err)
	}

	deleteQuery := `delete from intern_skill where intern_id = ?`
	if _, err := tx.Exec(deleteQuery, internID); err != nil {
		tx.Rollback()
		return fmt.Errorf("error in MapInternSkill transaction-delete mapping: %v", err)
	}

	if len(info.Skills) == 0 {
		return nil	
	}

	values := ""
	for i := range info.Skills {
		if i > 0 {
			values += ","
		}
		values += fmt.Sprintf("('%s', %d, '%s')", internID, info.Skills[i], info.SkillLevel[i])
	}

	updateQuery := fmt.Sprintf("insert into intern_skill (intern_id, technical_id, skill_level) values %s ON DUPLICATE KEY UPDATE skill_level=values(skill_level)", values)
	if _, err := tx.Exec(updateQuery); err != nil {
		tx.Rollback()
		return fmt.Errorf("error in MapInternSkill transaction-update mapping: %v", err)

	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("error in MapInternSkill transaction-commit: %v", err)

	}
	return nil
}
