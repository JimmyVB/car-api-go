package queries

func CreateUserQuery() string {
	return "insert into users (username, password) values ($1, $2)"
}

func GetLoginQuery() string {
	return "select id, username from users where username = $1 and password = $2"
}
