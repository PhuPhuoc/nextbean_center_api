package repository

import "github.com/PhuPhuoc/hrm_nextbean_api/services/DashboardServices/model"

func (store *dashboardStore) GetDashboardTotal() (model.DashboardTotalNumber, error) {
	data := model.DashboardTotalNumber{}
	rawsql := `SELECT (SELECT COUNT(*) FROM project WHERE status='in_progress') AS Total_Project,(SELECT COUNT(*) FROM intern WHERE ojt_id IN (SELECT id FROM ojt WHERE status='in_progress')) AS Total_intern_InProgress;`
	if err_query := store.db.QueryRow(rawsql).Scan(&data.TotalProjectInProgress, &data.TotalInternInProgress); err_query != nil {
		return model.DashboardTotalNumber{}, nil
	}
	return data, nil
}
