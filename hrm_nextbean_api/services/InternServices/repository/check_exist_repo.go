package repository

import (
	"fmt"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/intern_query"
)

func checkAccountIDExist(store *internStore, accID string) error {
	var flag bool
	query := query.QueryCheckExistAccountID()
	err := store.db.QueryRow(query, accID).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error when check exist account-id : %v", err)
	}
	if !flag {
		return fmt.Errorf("account'id not exists")
	}
	return nil
}

func getInternIDByAccountID(store *internStore, accID string) (*string, error) {
	var stuID *string
	query := query.QueryGetCurrentInternIDByAccountID()
	err := store.db.QueryRow(query, accID).Scan(&stuID)
	if err != nil {
		return nil, fmt.Errorf("error when get current student-code: %v", err)
	}
	return stuID, nil
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

func getAccIDByInternID(store *internStore, int_id string) (*string, error) {
	var accID *string
	query := query.QueryGetAccIDByInternID()
	err := store.db.QueryRow(query, int_id).Scan(&accID)
	if err != nil {
		return nil, fmt.Errorf("error when get accID by internID: %v", err)
	}
	return accID, nil
}
