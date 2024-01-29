/**
 * file: web/handler/product.go
 * description: file responsible for the handler layer of the application.
 * data: 01/28/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/glaucia86/fc-studies-hexagonal-architecture-go/application"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/products/{id}", n.With(
		negroni.Wrap(getProducts(service)),
	)).Methods("GET", "OPTIONS")
}

func getProducts(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}