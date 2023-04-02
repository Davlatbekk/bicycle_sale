package models

type Stock struct {
	StoreId     int      `json:"store_id"`
	StoreData   *Store   `json:"store_data"`
	ProductId   int      `json:"product_id"`
	ProductData *Product `json:"product_data"`
	Quantity    int      `json:"quantity"`
}
type StockPrimaryKey struct {
	StoreId int `json:"store_id"`
}

type ProductData struct {
	ProductId    int       `json:"product_id"`
	ProductName  string    `json:"product_name"`
	BrandId      int       `json:"brand_id"`
	BrandData    *Brand    `json:"brand_data"`
	CategoryId   int       `json:"category_id"`
	CategoryData *Category `json:"category_data"`
	ModelYear    int       `json:"model_year"`
	ListPrice    float64   `json:"list_price"`
	Quantity     int       `json:"quantity"`
}

type GetStock struct {
	StoreId  int            `json:"store_id"`
	Quantity int            `json:"quantity"`
	Products []*ProductData `json:"products"`
}

type CreateStock struct {
	StoreId   int `json:"store_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type UpdateStock struct {
	StoreId   int `json:"store_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type GetListStockRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListStockResponse struct {
	Count  int         `json:"count"`
	Stocks []*GetStock `json:"stocks"`
}
