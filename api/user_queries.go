package api

func CreateUserQuery() string {
	return "insert into users (id, username, password) values ($1, $2, $3)"
}

func GetLoginQuery() string {
	return "select id from users where username = $1 and password = $2"
}

func GetAddWishPokemonQuery() string {
	return "insert into wish_list (user_id, pokemon_id, comment) values ($1, $2, $3)"
}
