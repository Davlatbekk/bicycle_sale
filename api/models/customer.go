package models

type Customer struct {
	CustomerId int    `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	ZipCode    string `json:"zip_code"`
}

type CustomerPrimaryKey struct {
	CustomerId int `json:"customer_id"`
}

type CreateCustomer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   int    `json:"zip_code"`
}

type UpdateCustomer struct {
	CustomerId int    `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	ZipCode    int    `json:"zip_code"`
}

type GetListCustomerRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListCustomerResponse struct {
	Count     int         `json:"count"`
	Customers []*Customer `json:"customers"`
}
