package ojtquery

func QueryCreateNewOJT() string {
	return `insert into ojt(semester, university, start_at, end_at, created_at) values (?,?,?,?,?)`
}

func QueryCheckExistID() string {
	return `select exists(select 1 from ojt where id = ?)`
}

func QueryUpdateOJT() string {
	return `update ojt set semester = ?, university = ?, start_at = ?, end_at = ? where id = ?`
}

func QueryDeleteOJT() string {
	return `update ojt set deleted_at = ? where id = ?`
}
