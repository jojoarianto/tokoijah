package web

import (
	"encoding/json"
	"net/http"

	"github.com/jojoarianto/tokoijah/config"
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/julienschmidt/httprouter"
)

func addPurchase(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	purchase := model.Purchase{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&purchase); err != nil {
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

	// productsvc := service.NewProductService(sqlite3.NewProductRepo(db))
	// err = productsvc.Add(product)
	// if err != nil {
	// 	RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
	// 	return
	// }

	RespondWithJSON(w, http.StatusCreated, purchase)
}
