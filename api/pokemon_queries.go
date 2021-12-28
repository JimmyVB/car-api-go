package api

import "strings"

func getPokemonsQuery(filter PokemonFilter) string {
	var (
		d, g, t string
		clause  = false
		q       = "select id, name, category from pokemon"
		b       = strings.Builder{}
	)

	b.WriteString(q)

	if filter.Name != "" {
		d = "name like '%" + filter.Name + "%'"
		clause = true
	}

	if filter.Type != "" {
		g = "type like '%" + filter.Type + "%'"
		clause = true
	}

	if clause {
		var i int
		b.WriteString(" where ")
		if d != "" {
			b.WriteString(d)
			i = 1
		}

		if g != "" {
			if i == 1 {
				b.WriteString(" or ")
			}
			b.WriteString(g)
			i = 2
		}

		if t != "" {
			if i == 1 || i == 2 {
				b.WriteString(" or ")
			}
			b.WriteString(t)
		}

		return b.String()
	} else {
		return b.String()
	}
}

func CreateUserQuery() string {
	return "insert into users (id, username, password) values ($1, $2, $3)"
}
