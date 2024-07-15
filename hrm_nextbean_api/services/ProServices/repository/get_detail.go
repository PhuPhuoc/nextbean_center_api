package repository

import (
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) GetDetail(pro_id string) (model.ProjectDetail, error) {
	obj := model.ProjectDetail{}
	rawsql := rawQueryForGetProjectDetail()
	if err_query := store.db.QueryRow(rawsql, pro_id).Scan(&obj.Id, &obj.Name, &obj.Status, &obj.Description, &obj.EstStartTime, &obj.EstCompletionTime, &obj.TotalMember, &obj.TotalTask, &obj.TotalTaskCompleted); err_query != nil {
		return obj, err_query
	}
	return obj, nil
}

func rawQueryForGetProjectDetail() string {
	var query strings.Builder
	query.WriteString(`select p.id, p.name, p.status, p.description, p.est_start_time, p.est_completion_time, `)
	query.WriteString(`count(DISTINCT mem.intern_id) as total_member, `)
	query.WriteString(`count(DISTINCT t.id) as total_task, `)
	query.WriteString(`count(DISTINCT case when t.status = 'completed' then t.id end) AS total_completed_tasks `)
	query.WriteString(`from project p join task t on t.project_id=p.id join project_intern mem on mem.project_id=p.id `)
	query.WriteString(`where p.id=? `)
	query.WriteString(`group by p.id, p.name, p.status, p.description, p.est_start_time, p.est_completion_time`)
	return query.String()
}
