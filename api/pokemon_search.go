package api

import (
	"car-api/internal/database"
	"car-api/internal/logs"
)

type PokemonFilter struct {
	Name string `json:"name,omitempty"`
	Type string `json:"typ e,omitempty"`
}

type Pokemon struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Order    uint   `json:"order"`
	Height   uint   `json:"height"`
	Weight   uint   `json:"weight"`
	Category string `json:"category"`
}

type PokemonSearch interface {
	Search(filter PokemonFilter) ([]Pokemon, error)
}

type PokemonService struct {
	*database.PostgresClient
}

func (p *PokemonService) Search(filter PokemonFilter) ([]Pokemon, error) {

	tx, err := p.Begin()

	if err != nil {
		logs.Error("cannot create transaction")
		return nil, err
	}

	rows, err := tx.Query(getPokemonsQuery(filter))

	if err != nil {
		logs.Error("cannot read pokemons " + err.Error())
		_ = tx.Rollback()
		return nil, err
	}

	var _pokemons []Pokemon
	for rows.Next() {
		var pokemon Pokemon
		err := rows.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Category)
		if err != nil {
			logs.Error("cannot read pokemons " + err.Error())
		}
		_pokemons = append(_pokemons, pokemon)
	}

	_ = tx.Commit()

	return _pokemons, nil
}
