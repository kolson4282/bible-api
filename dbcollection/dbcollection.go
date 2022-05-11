package dbcollection

import (
	"fmt"

	"github.com/kolson4282/tdd-bible-api/graph/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBCollection struct {
	DB *gorm.DB
}

func NewDBCollection() *DBCollection {
	dc := &DBCollection{}
	dc.initialize()
	dc.createTables()
	dc.fillTables()
	return dc
}

func (dc *DBCollection) initialize() {
	db, err := gorm.Open(mysql.Open(getDBURL()), &gorm.Config{})
	dc.DB = db
	if err != nil {
		panic("Cannot connect to Database" + err.Error())
	}

}
func (dc *DBCollection) createTables() {
	err := dc.DB.AutoMigrate(&model.Character{})
	if err != nil {
		panic("error creating tables" + err.Error())
	}
}

func (dc *DBCollection) fillTables() {
	dc.DB.Create(&model.Character{
		ID:          1,
		Name:        "God",
		Description: "God",
	})
}

func getDBURL() string {
	USER := "user"
	DB_PASSWORD := "dbpass"
	HOST := "localhost:3305"
	DB_NAME := "bible_api"
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, DB_PASSWORD, HOST, DB_NAME)
}
