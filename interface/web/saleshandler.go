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

func addSales(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sales := model.Sales{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&sales); err != nil {
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

	salessvc := service.NewSalesService(sqlite3.NewSalesRepo(db))
	err = salessvc.Add(sales)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, response{
		Message:    "Created success",
		StatusCode: http.StatusCreated,
	})
}
