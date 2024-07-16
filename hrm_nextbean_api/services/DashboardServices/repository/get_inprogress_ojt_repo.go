package repository

import "github.com/PhuPhuoc/hrm_nextbean_api/services/DashboardServices/model"

func (store *dashboardStore) GetDashboardInpgrogressOJT() ([]model.DashboardOJTInProgress, error) {
	data := []model.DashboardOJTInProgress{}
	rawsql := `select o.id, o.semester, o.university, o.status, (select count(*) from intern where ojt_id=o.id) as total_intern from ojt o where o.status='in_progress'`
	rows, err := store.db.Query(rawsql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := model.DashboardOJTInProgress{}
		if err := rows.Scan(&obj.Id, &obj.Semester, &obj.University, &obj.Status, &obj.TotalIntern); err != nil {
			return nil, err
		}
		data = append(data, obj)
	}

	return data, nil
}
