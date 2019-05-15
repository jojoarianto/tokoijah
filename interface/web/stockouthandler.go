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

func addStockOut(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stockout := model.StockOut{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stockout); err != nil {
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

	stockoutsvc := service.NewStockOutService(sqlite3.NewStockOutRepo(db))
	err = stockoutsvc.Add(stockout)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, response{
		Message:    MsgDataSuccessCreated,
		StatusCode: http.StatusCreated,
	})
}
