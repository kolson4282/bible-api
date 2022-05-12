package main

import (
	"os"

	"github.com/kolson4282/tdd-bible-api/dbcollection"
	"github.com/kolson4282/tdd-bible-api/graph/resolver"
	"github.com/kolson4282/tdd-bible-api/server"
	"github.com/kolson4282/tdd-bible-api/utils"
)

const defaultPort = "8080"

func init() {
	utils.LoadEnv()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")

	collection := dbcollection.NewDBCollection(dbcollection.DBVars{
		DB_USER:     DB_USER,
		DB_PASSWORD: DB_PASSWORD,
		DB_HOST:     DB_HOST,
		DB_NAME:     DB_NAME,
	})
	server.RunServer(port, &resolver.Resolver{
		Collection: collection,
	})
}
