package middleware

import (
	"database/sql"
	"fmt"
)

func checkPMInProject(db *sql.DB, proID string, pmID string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from project_manager pm join account acc on pm.account_id=acc.id where pm.project_id = ? and pm.account_id = ? and acc.deleted_at is null)`
	if err_query := db.QueryRow(rawsql, proID, pmID).Scan(&flag); err_query != nil {
		return err_query
	}
	if !flag {
		return fmt.Errorf("pm (id: %v) is not part of the project or the pm's account has been deleted", pmID)
	}
	return nil
}

func checkMemInProject(db *sql.DB, proID string, inid string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from project_intern pin join intern i on pin.intern_id=i.id join account a on i.account_id=a.id where pin.project_id=? and pin.intern_id=? and a.deleted_at is null)`
	if err_query := db.QueryRow(rawsql, proID, inid).Scan(&flag); err_query != nil {
		return err_query
	}
	if !flag {
		return fmt.Errorf("member is not part of the project or the member's account has been deleted")
	}
	return nil
}

func checkInternIDBelongToTimeTable(db *sql.DB, tid string, inid string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from timetable where id=? and intern_id=? and deleted_at is null)`
	if err_query := db.QueryRow(rawsql, tid, inid).Scan(&flag); err_query != nil {
		return err_query
	}
	if !flag {
		return fmt.Errorf("this timetable does not belong to you")
	}
	return nil
}

func checkPMInTask(db *sql.DB, taskid string, pmID string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from task t join project p on t.project_id=p.id join project_manager pm on pm.project_id=p.id where t.id=? and pm.account_id=? and t.deleted_at is NULL)`
	if err_query := db.QueryRow(rawsql, taskid, pmID).Scan(&flag); err_query != nil {
		return err_query
	}
	if !flag {
		return fmt.Errorf("this pm (id: %v) does not belong to the project this task belongs to", pmID)
	}
	return nil
}

func checkMemInTask(db *sql.DB, taskid string, inid string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from task t where t.id=? and t.assigned_to=? and t.deleted_at is null)`
	if err_query := db.QueryRow(rawsql, taskid, inid).Scan(&flag); err_query != nil {
		return err_query
	}
	if !flag {
		return fmt.Errorf("this user is not part of this task so cannot comment")
	}
	return nil
}
