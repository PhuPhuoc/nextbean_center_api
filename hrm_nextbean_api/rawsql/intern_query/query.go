package internquery

var table = `intern`
var table_name = ` ` + table + ` `

func QueryCreateNewAccount() string {
	return `insert into` + table_name + `(id, user_name, email, password, role, created_at) values (?,?,?,?,?,?)`
}

func QueryCreateNewIntern() string {
	return `insert into` + table_name + `(id,account_id,ojt_id,avatar,gender,dateofbirth,phone_number,address) values (?,?,?,?,?,?,?,?)`
}

func QueryCheckDulicateDataInIntern() string {
	start := `select`
	part1 := ` case when exists (select 1 from account where email = ?) then 'email' end as email_exists,`
	part2 := ` case when exists (select 1 from intern where id = ?) then 'id' end as studentcode_exists,`
	part3 := ` case when exists (select 1 from intern where phone_number = ?) then 'phone_number' end as phonenumber_exists`
	end := `from DUAL;`
	return start + ` ` + part1 + ` ` + part2 + ` ` + part3 + ` ` + end
}
