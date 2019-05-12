package web

import (
	"encoding/json"
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
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	RespondWithJSON(w, http.StatusOK, products)
}

func addProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		RespondWithError(w, http.StatusBadRequest, model.ErrBadParamInput.Error())
		return
	}
	defer r.Body.Close()

	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	productsvc := service.NewProductService(sqlite3.NewProductRepo(db))
	err = productsvc.Add(product)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, product)
}
