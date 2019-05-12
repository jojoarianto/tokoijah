package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jojoarianto/tokoijah/config"
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/infrastructure/sqlite3"
	"github.com/jojoarianto/tokoijah/service"
	"github.com/julienschmidt/httprouter"
)

const (
	// URIDbConn database connection
	URIDbConn = "tokoijah.sqlite3"
	// Dialeg driver
	Dialeg = "sqlite3"
)

func getAllProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	products := []model.Product{}

	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		JSON(w, http.StatusInternalServerError, products)
	}
	defer db.Close()

	JSON(w, http.StatusOK, products)
}

func addProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		JSON(w, http.StatusBadRequest, product)
	}
	defer r.Body.Close()

	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		JSON(w, http.StatusInternalServerError, product)
	}
	defer db.Close()

	productsvc := service.NewProductService(sqlite3.NewProductRepo(db))
	err = productsvc.Add(product)
	if err != nil {
		log.Println(err)
		JSON(w, http.StatusInternalServerError, product)
	}

	JSON(w, http.StatusCreated, product)
}
