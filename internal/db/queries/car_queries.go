package queries

func Insert() string {
	return "insert into cars (marca, model, price) values ($1, $2, $3)"
}

func GetAll() string {
	return "select id, marca, model, price from cars"
}

func GetOne() string {
	return "select id, marca, model, price from cars where id = $1"
}

func Update() string {
	return "update cars set marca = $2, model = $3, price = $4 where id = $1"
}

func Delete() string {
	return "delete from cars where id = $1"
}