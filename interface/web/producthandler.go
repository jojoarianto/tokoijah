package web

import (
	"net/http"

	"github.com/jojoarianto/tokoijah/config"
	"github.com/jojoarianto/tokoijah/domain/model"
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
