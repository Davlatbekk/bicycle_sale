package models

type Category struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
type CategoryPrimaryKey struct {
	CategoryId int `json:"category_id"`
}

type CreateCategory struct {
	CategoryName string `json:"category_name"`
}

type UpdateCategory struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type GetListCategoryRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListCategoryResponse struct {
	Count      int         `json:"count"`
	Categories []*Category `json:"categories"`
}
