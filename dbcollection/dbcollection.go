package dbcollection

import (
	"fmt"

	"github.com/kolson4282/tdd-bible-api/graph/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBCollection struct {
	DB     *gorm.DB
	DBVars DBVars
}

type DBVars struct {
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_NAME     string
}

func NewDBCollection(DBVars DBVars) *DBCollection {
	dc := &DBCollection{}
	dc.DBVars = DBVars
	dc.initialize()
	dc.createTables()
	// dc.fillTables()
	return dc
}

func (dc *DBCollection) initialize() {
	db, err := gorm.Open(mysql.Open(dc.getDBURL()), &gorm.Config{})
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

// func (dc *DBCollection) fillTables() {
// 	dc.DB.Create(&model.Character{
// 		ID:          1,
// 		Name:        "God",
// 		Description: "God",
// 	})
// }

func (dc DBCollection) getDBURL() string {

	DB_USER := dc.DBVars.DB_USER
	DB_PASSWORD := dc.DBVars.DB_PASSWORD
	DB_HOST := dc.DBVars.DB_HOST
	DB_NAME := dc.DBVars.DB_NAME
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
}
