package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	l := log.New(os.Stdout, "server ", log.LstdFlags)

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
	pg, dbConnErr := sql.Open("postgres", connstring)
	if dbConnErr != nil {
		l.Fatal("Failed to connect to database", dbConnErr)
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))

	userHandler := handlers.NewUserHandler(ctx, l, pg)
	r.Get("/users", userHandler.GetManyUser)
	r.Post("/users", userHandler.CreateUser)
	r.Get("/users/{userId}", userHandler.GetOneUserById)
	r.Patch("/users/{userId}", userHandler.UpdateOneUserById)
	r.Delete("/users/{userId}", userHandler.DeleteOneUserById)

	http.ListenAndServe(":9000", r)
}
