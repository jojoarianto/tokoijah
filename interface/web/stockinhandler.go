package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jojoarianto/tokoijah/config"
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/infrastructure/sqlite3"
	"github.com/jojoarianto/tokoijah/service"
	"github.com/julienschmidt/httprouter"
)

func addStockIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	purchaseID, err := strconv.Atoi(ps.ByName("purchase_id"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, model.ErrBadParamInput.Error())
		return
	}

	stockin := model.StockIn{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stockin); err != nil {
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

	stockinsvc := service.NewStockInService(sqlite3.NewStockInRepo(db))
	err = stockinsvc.Add(purchaseID, stockin)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	purchasesvc := service.NewPurchaseService(sqlite3.NewPurchaseRepo(db), sqlite3.NewProductRepo(db))
	purchase, err := purchasesvc.GetByID(purchaseID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, purchase)
}
