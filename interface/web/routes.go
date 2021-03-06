package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	// URIDbConn database connection
	URIDbConn = "tokoijah.sqlite3"
	// Dialeg driver
	Dialeg = "sqlite3"
	// MsgSuccessCreated msg for success created
	MsgDataSuccessCreated = "Created success"
)

// Run start server
func Run(port int) error {
	log.Printf("Server running at http://localhost:%d/", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
}

// Routes returns the initialized router
func Routes() *httprouter.Router {
	r := httprouter.New()

	// Index Route
	r.GET("/", index)

	// Route for Product
	r.GET("/products", getAllProduct)
	r.GET("/products/:product_id", getProductByID)
	r.POST("/products", addProduct)

	// Route for Purchase
	r.GET("/purchases", getAllPurchase)
	r.GET("/purchases/:purchase_id", getPurchaseByID)
	r.POST("/purchases", addPurchase)

	// Route for StockIn
	r.POST("/purchases/:purchase_id/stockin", addStockIn)

	// Route for StockOut
	r.POST("/stockout", addStockOut)

	// Route for Sales
	r.GET("/sales", getAllSales)
	r.POST("/sales", addSales)

	// Route for import data
	r.POST("/import/products", importCsvProduct)

	// Route for export data
	r.GET("/export/products", exportProductToCSV)
	r.GET("/export/purchases", exportPurchaseToCSV)
	r.GET("/export/stockout", cxportStockOutToCSV)
	r.GET("/export/sales", cxportSalesToCSV)

	return r
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RespondWithJSON(w, http.StatusOK, "TokoIjah API")
}
