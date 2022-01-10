package queries

func CreateUserQuery() string {
	return "insert into users (username, password) values ($1, $2)"
}

func GetLoginQuery() string {
	return "select u.id, u.username, ro.nombre from users u " +
		"inner join usuarios_roles ur on u.id = ur.usuario_id " +
		"inner join roles ro on ro.id = ur.role_id where u.username = $1 and u.password = $2"
}
