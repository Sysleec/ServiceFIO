package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	pplApi "github.com/Sysleec/ServiceFIO/internal/api/people"
	pplRepository "github.com/Sysleec/ServiceFIO/internal/repository/people"
	"github.com/Sysleec/ServiceFIO/internal/repository/people/sqlc"
	pplService "github.com/Sysleec/ServiceFIO/internal/service/people"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	port := os.Getenv("HTTP_PORT")
	myURL := os.Getenv("HTTP_HOST")
	myDSN := os.Getenv("PG_DSN")

	conn, err := sql.Open("postgres", myDSN)
	if err != nil {
		log.Fatalf("Can't connect to DB: %v", err)
	}

	db := sqlc.New(conn)

	pplRepo := pplRepository.NewRepo(db)
	pplServ := pplService.NewService(pplRepo)
	pApi := pplApi.NewImplementation(pplServ)

	app := chi.NewRouter()
	app.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1 := chi.NewRouter()

	v1.Post("/people", pApi.Create)

	app.Mount("/v1", v1)

	serv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", myURL, port),
		Handler: app,
	}

	fmt.Printf("Server serving on port %v...", port)

	log.Fatal(serv.ListenAndServe())

}
