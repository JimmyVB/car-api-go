package api

import (
	"car-api/internal/middleware"
	"github.com/gofiber/fiber"
)

func SetupPokemonsRoutes(app *fiber.App, tokenKey string) {
	s := start(tokenKey)
	grp := app.Group("/pokemons", s.SearchPokemonHandler)
	grp.Get("/", s.SearchPokemonHandler)
}

func SetupUsersRouters(app *fiber.App, tokenKey string) {
	s := start(tokenKey)
	grp := app.Group("/users")
	grp.Post("/", s.CreateUserHandler)
	grp.Post("/login", s.LoginHandler)
	grp.Use(middleware.JwtMiddleware(tokenKey)).Post("/wishlist", s.WishListHandler)
}
