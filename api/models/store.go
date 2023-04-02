package models

type Store struct {
	StoreId   int    `json:"store_id"`
	StoreName string `json:"store_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
}

type StorePrimaryKey struct {
	StoreId int `json:"store_id"`
}

type CreateStore struct {
	StoreId   int    `json:"store_id"`
	StoreName string `json:"store_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
}

type UpdateStore struct {
	StoreId   int    `json:"store_id"`
	StoreName string `json:"store_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
}

type GetListStoreRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListStoreResponse struct {
	Count  int      `json:"count"`
	Stores []*Store `json:"stores"`
}
