package domain

// Adapter
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	return CustomerRepositoryStub{
		customers: []Customer{
			{Id: "101", Name: "Alex", City: "Moscow",
				Zipcode: "111121", BirthDate: "2002-12-08", Status: "1"},
			{Id: "102", Name: "Maria", City: "Kazan",
				Zipcode: "110055", BirthDate: "2008-11-08", Status: "1"},
		}}
}
