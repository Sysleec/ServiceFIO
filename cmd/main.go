package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	pplApi "github.com/Sysleec/ServiceFIO/internal/api/people"
	pplRepository "github.com/Sysleec/ServiceFIO/internal/repository/people"
	"github.com/Sysleec/ServiceFIO/internal/repository/people/sqlc"
	pplService "github.com/Sysleec/ServiceFIO/internal/service/people"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"

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
	logLevel := os.Getenv("LOG_LEVEL")

	conn, err := sql.Open("postgres", myDSN)
	if err != nil {
		log.Fatalf("Can't connect to DB: %v", err)
	}

	logLVL, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("Can't parse log level: %v", err)
	}

	zerolog.SetGlobalLevel(logLVL)

	db := sqlc.New(conn)

	pplRepo := pplRepository.NewRepo(db)
	pplServ := pplService.NewService(pplRepo)
	pAPI := pplApi.NewImplementation(pplServ)

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

	v1.Post("/people", pAPI.Create)
	v1.Delete("/people/{id}", pAPI.Delete)
	v1.Get("/people", pAPI.Get)
	v1.Patch("/people/{id}", pAPI.Update)

	app.Mount("/v1", v1)

	serv := &http.Server{
		Addr:              fmt.Sprintf("%v:%v", myURL, port),
		Handler:           app,
		ReadHeaderTimeout: 10 * time.Second,
	}

	fmt.Printf("Server serving on port %v...\n", port)

	log.Fatal(serv.ListenAndServe())

}
