package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jojoarianto/tokoijah/config"
	"github.com/jojoarianto/tokoijah/domain/model"
)

func main() {

	conf := config.NewConfig("sqlite3", "tokoijah.sqlite3")
	db, _ := conf.ConnectDB()

	DBMigrate(db)
}

// DBMigrate will create and migrate the tables
func DBMigrate(db *gorm.DB) *gorm.DB {

	db.AutoMigrate(&model.Product{}, &model.Purchase{}, &model.StockIn{}, &model.Sales{}, &model.StockOut{})
	log.Println("Schema migration has been procceed")

	return db
}
