package accountquery

var table = `account`
var table_name = ` ` + table + ` `
var table_field = ` id, user_name, email, password, role `

func QueryCreateNewAccount() string {
	return `insert into` + table_name + `(id, user_name, email, password, role, created_at) values (?,?,?,?,?,?)`
}

func QueryCheckExistEmail() string {
	return `select exists(select 1 from` + table_name + `where email = ?)`
}

func QueryGetAccountByEmailForLogin() string {
	return `select` + table_field + `from` + table_name + `where email = ? and deleted_at is null`
}
