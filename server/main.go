package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/handlers"
	"server/middlewares"
	"server/services"

	_ "server/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title						LinkVault API
//	@version					1.0
//	@description				Docs for LinkVault API
//	@host						localhost:9000
//	@BasePath					/
//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
func main() {
	l := log.New(os.Stdout, "", log.LstdFlags)

	if err := godotenv.Load(); err != nil {
		l.Fatalln("Failed loading .env file: ", err)
	}

	connstring := fmt.Sprintf(
		"user='%s' password='%s' dbname='%s' host='%s' sslmode='disable'",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("DB_HOST"),
	)

	ctx := context.Background()
	pg, dbConnErr := sqlx.Open("postgres", connstring)
	if dbConnErr != nil {
		l.Fatal("Failed to connect to database", dbConnErr)
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))

	r.Get("/docs/*", httpSwagger.WrapHandler)

	linkService := services.NewLinkService(ctx, l, pg)

	userHandler := handlers.NewUserHandler(ctx, l, pg)
	authHandler := handlers.NewAuthHandler(ctx, l, pg)
	linkHanlder := handlers.NewLinkHandler(ctx, l, *linkService)

	// Public
	r.Group(func(r chi.Router) {
		r.Post("/auth", authHandler.Login)

		r.Get("/users", userHandler.GetManyUser)
		r.Post("/users", userHandler.CreateUser)
	})

	// Protected
	r.Group(func(r chi.Router) {
		r.Use(middlewares.JwtAuth(ctx, l, pg))

		r.Get("/users/{userId}", userHandler.GetOneUserById)
		r.Patch("/users/{userId}", userHandler.UpdateOneUserById)
		r.Delete("/users/{userId}", userHandler.DeleteOneUserById)

		r.Post("/folders/{folderId}/links", linkHanlder.CreateLink)
		r.Get("/links", linkHanlder.GetManyLinks)
		r.Get("/folders/{folderId}/links", linkHanlder.GetManyLinksInFolder)
		r.Patch("/links/{linkId}", linkHanlder.UpdateLink)
	})
	http.ListenAndServe(":9000", r)
}
