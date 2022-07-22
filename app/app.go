package app

import (
	"log"
	"net/http"

	"github.com/deltaneverhood/bank/domain"
	"github.com/deltaneverhood/bank/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	// ch := CusomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CusomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
