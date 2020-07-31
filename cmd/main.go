package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/medhub-hex/pkg/database"
	"github.com/medhub-hex/pkg/fhir"
	"github.com/medhub-hex/pkg/http/rest"
)

func main() {
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	sq := database.NewSquirrel()
	repo := fhir.NewPostresRepository(dbpool, sq)
	service := fhir.NewService(repo)
	fhirHandler := fhir.NewHandler(service).(*fhir.Handler)

	r := chi.NewRouter()
	r.Use(rest.Recoverer)
	r.Use(rest.ContentTypeJson)
	r.Use(rest.Cors)
	r.Route("/api/baseR4/{resourceType}", func(r chi.Router) {
		r.Mount("/", fhirHandler.Routes())
	})

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  160 * time.Second,
		WriteTimeout: 160 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port: ", server.Addr)
		errs <- server.ListenAndServe()
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated cause: %s", <-errs)
}
