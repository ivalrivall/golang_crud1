package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ivalrivall/golang_crud1/middleware"

	"github.com/gorilla/mux"
)

func healthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("health middleware")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	healthchecks := router.PathPrefix("/health").Subrouter()
	healthchecks.Use(healthMiddleware)
	healthchecks.HandleFunc("/ready", handler)

	api := router.PathPrefix("/api").Subrouter()
	users := api.PathPrefix("/user").Subrouter()
	brands := api.PathPrefix("/brand").Subrouter()
	products := api.PathPrefix("/product").Subrouter()
	order := api.PathPrefix("/order").Subrouter()
	// users.Path("").Methods(http.MethodGet).HandlerFunc(middleware.GetAllUser)
	users.Path("").Methods(http.MethodPost).HandlerFunc(middleware.CreateUser)
	brands.Path("").Methods(http.MethodPost).HandlerFunc(middleware.CreateBrand)
	products.Path("").Methods(http.MethodPost).HandlerFunc(middleware.CreateProduct)
	products.Path("/{id}").Methods(http.MethodGet).HandlerFunc(middleware.GetProduct)
	products.Path("/brand/{id}").Methods(http.MethodGet).HandlerFunc(middleware.GetProductByBrand)
	order.Path("").Methods(http.MethodPost).HandlerFunc(middleware.CreateOrder)
	order.Path("/{id}").Methods(http.MethodGet).HandlerFunc(middleware.GetDetailTransactionById)
	// users.Path("/{id}").Methods(http.MethodPut).HandlerFunc(middleware.UpdateUser)
	// users.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(middleware.DeleteUser)

	return router
}
