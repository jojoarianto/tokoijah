package web

import (
	"encoding/csv"
	"net/http"
	"os"
	"strconv"

	"github.com/jojoarianto/tokoijah/config"
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/infrastructure/sqlite3"
	"github.com/jojoarianto/tokoijah/service"
	"github.com/julienschmidt/httprouter"
)

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
