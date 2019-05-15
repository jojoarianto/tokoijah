package web

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jojoarianto/tokoijah/config"
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/infrastructure/sqlite3"
	"github.com/jojoarianto/tokoijah/service"
	"github.com/julienschmidt/httprouter"
)

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
			Sku:  line[0],
			Name: line[1],
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

	RespondWithJSON(w, http.StatusAccepted, response{
		Message:    MsgDataSuccessCreated,
		StatusCode: http.StatusCreated,
	})
}
