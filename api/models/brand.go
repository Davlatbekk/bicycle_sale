package models

type Brand struct {
	BrandId   int    `json:"brand_id"`
	BrandName string `json:"brand_name"`
}
type BrandPrimaryKey struct {
	BrandId int `json:"brand_id"`
}

type BrandCategory struct {
	BrandId int `json:"brand_id"`
}

type CreateBrand struct {
	BrandName string `json:"brand_name"`
}

type UpdateBrand struct {
	BrandId   int    `json:"brand_id"`
	BrandName string `json:"brand_name"`
}

type GetListBrandRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListBrandResponse struct {
	Count  int      `json:"count"`
	Brands []*Brand `json:"brands"`
}
