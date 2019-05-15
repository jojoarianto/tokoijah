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
	"gopkg.in/go-playground/validator.v9"
)

func getProductByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	productID, err := strconv.Atoi(ps.ByName("product_id"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, model.ErrBadParamInput.Error())
		return
	}

	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	productsvc := service.NewProductService(sqlite3.NewProductRepo(db))
	product, err := productsvc.GetByID(productID)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, model.ErrNotFound.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, product)
}

func getAllProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	products := []model.Product{}

	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	productsvc := service.NewProductService(sqlite3.NewProductRepo(db))
	products, err = productsvc.GetAll()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}

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

	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		RespondWithError(w, http.StatusBadRequest, model.ErrBadParamInput.Error())
		return
	}

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

	RespondWithJSON(w, http.StatusCreated, response{
		Message:    MsgDataSuccessCreated,
		StatusCode: http.StatusCreated,
	})
}
