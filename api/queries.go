package api

func Insert() string {
	return "insert into cars (id, marca, model, price) values ($1, $2, $3, $4)"
}
