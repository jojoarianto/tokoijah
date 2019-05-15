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
		record = append(record, strconv.Itoa(worker.Stock))
		writer.Write(record)
	}

	writer.Flush()
	RespondWithJSON(w, http.StatusCreated, products)
	RespondWithJSON(w, http.StatusCreated, response{
		Message:    "Export data products to csv success check your export file at _csv/export_products.csv",
		StatusCode: 200,
	})
}

func exportPurchaseToCSV(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	purchasesvc := service.NewPurchaseService(sqlite3.NewPurchaseRepo(db), sqlite3.NewProductRepo(db))
	purchases, err := purchasesvc.GetAll()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	csvData, err := os.Create("_csv/export_purchases.csv")
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	defer csvData.Close()

	writer := csv.NewWriter(csvData)

	var record []string
	record = append(record, "Waktu")
	record = append(record, "SKU")
	record = append(record, "Nama Barang")
	record = append(record, "Jumlah Pemesanan")
	record = append(record, "Jumlah Diterima")
	record = append(record, "Harga Beli")
	record = append(record, "Total")
	record = append(record, "Nomer Kwitansi")
	record = append(record, "Catatan")
	writer.Write(record)

	for _, worker := range purchases {
		var record []string
		record = append(record, worker.PurchaseTime.Format("2006/01/02 15:04"))
		record = append(record, worker.Product.Sku)
		record = append(record, worker.Product.Name)
		record = append(record, strconv.Itoa(worker.OrderQty))
		record = append(record, strconv.Itoa(worker.ReceivedQty))
		record = append(record, strconv.Itoa(worker.Price))
		record = append(record, strconv.Itoa(worker.TotalPrice))
		record = append(record, worker.Receipt)

		progress := worker.StockIn
		note := ""
		for _, childProgress := range progress {
			note += childProgress.StockInTime.Format("2006/01/02")
			note += " terima " + strconv.Itoa(childProgress.Qty) + "; "
		}

		if worker.StausInCode == 0 {
			note += "Masih menunggu"
		}
		record = append(record, note)

		writer.Write(record)
	}
	writer.Flush()

	RespondWithJSON(w, http.StatusCreated, response{
		Message:    "Export data purchases to csv success check your export file at _csv/export_purchases.csv",
		StatusCode: 200,
	})
}

func cxportStockOutToCSV(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, model.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	stockout := []model.StockOut{}
	db.Preload("Product").Find(&stockout)

	csvData, err := os.Create("_csv/export_stockouts.csv")
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	defer csvData.Close()

	writer := csv.NewWriter(csvData)

	var record []string
	record = append(record, "Waktu")
	record = append(record, "SKU")
	record = append(record, "Nama Barang")
	record = append(record, "Jumlah Keluar")
	record = append(record, "Harga Jual")
	record = append(record, "Total")
	record = append(record, "Catatan")
	writer.Write(record)

	for _, worker := range stockout {
		var record []string
		record = append(record, worker.StockOutTime.Format("2006-01-02 15:04:05"))
		record = append(record, worker.Product.Sku)
		record = append(record, worker.Product.Name)
		record = append(record, strconv.Itoa(worker.Qty))
		record = append(record, strconv.Itoa(worker.SellPrice))
		record = append(record, strconv.Itoa(worker.TotalPrice))

		switch code := worker.StatusOutCode; code {
		case 1:
			str := "Pesanan ID-" + string(worker.SalesID)
			record = append(record, str)
		case 2:
			record = append(record, "Barang Hilang")
		case 3:
			record = append(record, "Barang Rusak")
		case 4:
			record = append(record, "Barang Rusak")
		default:
			record = append(record, "")
		}
		writer.Write(record)
	}
	writer.Flush()

	RespondWithJSON(w, http.StatusCreated, response{
		Message:    "Export data purchases to csv success check your export file at _csv/export_stockout.csv",
		StatusCode: 200,
	})
}
