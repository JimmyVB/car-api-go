package api

import (
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

func SetupPokemonsRoutes(app *fiber.App) {
	s := start()
	grp := app.Group("/pokemons", s.SearchPokemonHandler)
	grp.Get("/", s.SearchPokemonHandler)
}

func SetupUsersRouters(app *fiber.App) {
	s := start()
	grp := app.Group("/users")
	grp.Post("/", s.CreateUserHandler)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("mysecretkey"),
	}))

	grp.Get("/wishlist", s.WishListHandler)
}
