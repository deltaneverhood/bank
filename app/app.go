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
	ch := CusomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
