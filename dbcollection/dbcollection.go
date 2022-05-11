package dbcollection

import (
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
	return dc
}

func (dc *DBCollection) initialize() {
	db, err := gorm.Open(mysql.Open(getDBURL()), &gorm.Config{})
	dc.DB = db
	if err != nil {
		panic("Cannot connect to Database" + err.Error())
	}
	err = dc.DB.AutoMigrate(&model.Character{})
	if err != nil {
		panic("error creating tables" + err.Error())
	}
	db.Create(&model.Character{
		ID: 1,
	})
}

func getDBURL() string {
	return "user:dbpass@tcp(localhost:3305)/bible_api"
}
