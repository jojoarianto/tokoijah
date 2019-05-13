package web

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jojoarianto/tokoijah/config"
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/infrastructure/sqlite3"
	"github.com/jojoarianto/tokoijah/service"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/go-playground/validator.v9"
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

	RespondWithJSON(w, http.StatusCreated, product)
}

func importCsvProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	csvData, err := os.Open("_csv/import_products.csv")
	if err != nil {
		log.Print(err.Error())
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}

	reader := csv.NewReader(csvData)
	var products []model.Product
	for {
		line, err := reader.Read()
		if err == io.EOF { // end of line
			break
		} else if err != nil {
			RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
			return
		}

		products = append(products, model.Product{
			Sku:    line[0],
			Name:   line[1],
			Stocks: 0,
		})
	}

	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	productsvc := service.NewProductService(sqlite3.NewProductRepo(db))
	err = productsvc.AddMany(products)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}

	RespondWithJSON(w, http.StatusAccepted, products)
}

func exportProductToCSV(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	productsvc := service.NewProductService(sqlite3.NewProductRepo(db))
	products, err := productsvc.GetAll()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}

	csvData, err := os.Create("_csv/export_products.csv")
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}
	defer csvData.Close()

	writer := csv.NewWriter(csvData)

	var record []string
	record = append(record, "SKU")
	record = append(record, "Nama Item")
	record = append(record, "Jumlah Sekarang")
	writer.Write(record)

	for _, worker := range products {
		var record []string
		record = append(record, worker.Sku)
		record = append(record, worker.Name)
		record = append(record, strconv.Itoa(worker.Stocks))
		writer.Write(record)
	}

	writer.Flush()
	RespondWithJSON(w, http.StatusCreated, products)
}
