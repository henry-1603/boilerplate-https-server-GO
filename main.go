package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/joho/godotenv"
)

func main() {

	// //////////////////////////////////////////////////////////////////

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// //////////////////////////////////////////////////////////////////
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT not found in environment variable")
	}

	// //////////////////////////////////////////////////////////////////

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*" , "https://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE" , "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	// new router
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz" , handlerReadiness)
	v1Router.Get("/err" , handlerErr)
	

	// mounting
	router.Mount("/v1" , v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Printf("Server Started on port %v" , portString)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}


	// //////////////////////////////////////////////////////////////////

}
