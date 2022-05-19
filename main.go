package main

import (
	"fmt"
	"os"

	"github.com/kolson4282/bible-api/dbcollection"
	"github.com/kolson4282/bible-api/graph/resolver"
	"github.com/kolson4282/bible-api/server"
	"github.com/kolson4282/bible-api/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	collection := dbcollection.NewDBCollection(createDatabase())
	server.RunServer(port, &resolver.Resolver{
		Collection: collection,
	})
}

func createDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open(getDBURL()), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to Database" + err.Error())
	}
	return db
}

func getDBURL() string {

	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
}
