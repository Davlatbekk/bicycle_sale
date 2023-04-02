package models

type Code struct {
	Code_Id         int     `json:"code_id"`
	CodeName        string  `json:"code_name"`
	Discount        float64 `json:"discount"`
	DiscountType    string  `json:"discount_type"`
	OrderLimitPrice float64 `json:"order_limit_price"`
}
type Promocode struct {
	PromocodeId     int     `json:"promocode_id"`
	PromocodeName   string  `json:"promocode_name"`
	Discount        float64 `json:"discount"`
	DiscountType    int     `json:"discount_type"`
	OrderLimitPrice float64 `json:"order_limit_price"`
}

type CodePrimaryKey struct {
	Code_Id int `json:"code_id"`
}

type CreateCode struct {
	CodeName        string  `json:"code_name"`
	Discount        float64 `json:"discount"`
	DiscountType    string  `json:"discount_type"`
	OrderLimitPrice float64 `json:"order_limit_price"`
}

type UpdateCode struct {
	Code_Id         int     `json:"code_id"`
	CodeName        string  `json:"code_name"`
	Discount        float64 `json:"discount"`
	DiscountType    string  `json:"discount_type"`
	OrderLimitPrice float64 `json:"order_limit_price"`
}

type GetListCodeRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListCodeResponse struct {
	Count int     `json:"count"`
	Codes []*Code `json:"codes"`
}
