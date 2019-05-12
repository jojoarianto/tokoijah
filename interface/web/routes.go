package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
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
	r.POST("/products", addProduct)

	return r
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	JSON(w, http.StatusOK, "TokoIjah API")
}
