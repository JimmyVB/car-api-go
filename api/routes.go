package api

import (
	"car-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

//
//func SetupPokemonsRoutes(app *fiber.App, tokenKey string) {
//	s := cars.start(tokenKey)
//	grp := app.Group("/pokemons")
//	grp.Get("/", s.SearchPokemonHandler)
//}

func SetupUsersRouters(app *fiber.App, tokenKey string) {
	s := Start(tokenKey)
	grp := app.Group("/users")
	grp.Post("/", s.CreateUserHandler)
	grp.Post("/login", s.LoginHandler)
	grp.Use(middleware.JwtMiddleware(tokenKey)).Post("/wishlist", s.WishListHandler)
}

func SetupCarRouters(app *fiber.App, tokenKey string) {
	s := Start(tokenKey)
	grp := app.Group("/cars")
	grp.Use(middleware.JwtMiddleware(tokenKey)).Post("/create", s.CreateHandler)
}
