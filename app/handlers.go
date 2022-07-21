package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/deltaneverhood/bank/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

type CusomerHandlers struct {
	service service.CustomerService
}

func (ch *CusomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Alex", "Moscow", "111121"},
	// 	{"Maria", "Kazan", "110055"},
	// }

	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		// XML encoding
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		// JSON encoding
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello")
// }

// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])
// }
// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post request received")
// }
