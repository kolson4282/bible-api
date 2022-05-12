package dbcollection

import (
	"github.com/kolson4282/tdd-bible-api/graph/model"
	"gorm.io/gorm"
)

type DBCollection struct {
	DB *gorm.DB
}

func NewDBCollection(db *gorm.DB) *DBCollection {
	dc := &DBCollection{}
	dc.DB = db
	dc.createTables()
	return dc
}

func (dc *DBCollection) createTables() {
	err := dc.DB.AutoMigrate(&model.Character{})
	if err != nil {
		panic("error creating tables" + err.Error())
	}
}
