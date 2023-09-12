package main

import (
	"net/http"

	"github.com/Jhon-Henkel/go_lang_api_example/tree/main/configs"
	"github.com/Jhon-Henkel/go_lang_api_example/tree/main/internal/entity"
	"github.com/Jhon-Henkel/go_lang_api_example/tree/main/internal/infra/database"
	"github.com/Jhon-Henkel/go_lang_api_example/tree/main/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// esse middleware é para recuperar o panic, não deixando o servidor cair
	router.Use(middleware.Recoverer)
	router.Use(middleware.WithValue("jwt", config.TokenAuthKey))
	router.Use(middleware.WithValue("jwtExpires", config.JWTExpiresIn))

	router.Route("/products", func(router chi.Router) {
		router.Use(jwtauth.Verifier(config.TokenAuthKey))
		router.Use(jwtauth.Authenticator)
		router.Post("/", productHandler.CreateProduct)
		router.Get("/", productHandler.GetProducts)
		router.Get("/{id}", productHandler.GetProduct)
		router.Put("/{id}", productHandler.UpdateProduct)
		router.Delete("/{id}", productHandler.DeleteProduct)
	})

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	router.Post("/users", userHandler.CreateUser)
	router.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8000", router)
}
