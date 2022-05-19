package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/kolson4282/tdd-bible-api/graph/generated"
	"github.com/kolson4282/tdd-bible-api/graph/resolver"
	"github.com/rs/cors"
)

func RunServer(port string, resolver *resolver.Resolver) {
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	router.Handle("/", playground.Handler("Bible API", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
