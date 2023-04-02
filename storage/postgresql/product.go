package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(ctx context.Context, req *models.CreateProduct) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO products(
			product_id, 
			product_name, 
			brand_id,
			category_id,
			model_year,
			list_price
		)
		VALUES (
			(
				SELECT MAX(product_id) + 1 FROM products
			)
			, $1, $2, $3, $4, $5) RETURNING product_id
	`
	fmt.Println(query)

	err := r.db.QueryRow(ctx, query,
		req.ProductName,
		req.BrandId,
		req.CategoryId,
		req.ModelYear,
		req.ListPrice,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *productRepo) GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {

	var (
		query   string
		product models.Product
	)

	query = `
		SELECT
			p.product_id, 
			p.product_name, 
			p.brand_id,

			b.brand_id,
			b.brand_name,

			p.category_id,
			c.category_id,
			c.category_name,
			
			p.model_year,
			p.list_price
		FROM products AS p
		JOIN brands AS b ON b.brand_id = p.brand_id
		JOIN categories AS c ON c.category_id = p.category_id
		WHERE product_id = $1
	`
	product.BrandData = &models.Brand{}
	product.CategoryData = &models.Category{}

	err := r.db.QueryRow(ctx, query, req.ProductId).Scan(
		&product.ProductId,
		&product.ProductName,
		&product.BrandId,
		&product.BrandData.BrandId,
		&product.BrandData.BrandName,
		&product.CategoryId,
		&product.CategoryData.CategoryId,
		&product.CategoryData.CategoryName,
		&product.ModelYear,
		&product.ListPrice,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) GetList(ctx context.Context, req *models.GetListProductRequest) (resp *models.GetListProductResponse, err error) {

	resp = &models.GetListProductResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			p.product_id, 
			p.product_name, 
			p.brand_id,

			b.brand_id,
			b.brand_name,

			p.category_id,
			c.category_id,
			c.category_name,
			
			p.model_year,
			p.list_price
		FROM products AS p
		JOIN brands AS b ON b.brand_id = p.brand_id
		JOIN categories AS c ON c.category_id = p.category_id
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		product.BrandData = &models.Brand{}
		product.CategoryData = &models.Category{}
		err = rows.Scan(
			&resp.Count,
			&product.ProductId,
			&product.ProductName,
			&product.BrandId,
			&product.BrandData.BrandId,
			&product.BrandData.BrandName,
			&product.CategoryId,
			&product.CategoryData.CategoryId,
			&product.CategoryData.CategoryName,
			&product.ModelYear,
			&product.ListPrice,
		)
		if err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &product)
	}

	return resp, nil
}

func (r *productRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		products
		SET
			product_id = :product_id, 
			product_name = :product_name, 
			brand_id = :brand_id,
			category_id = :category_id,
			model_year = :model_year,
			list_price = :list_price
		WHERE product_id = :product_id
	`

	params = map[string]interface{}{
		"product_id":   req.ProductId,
		"product_name": req.ProductName,
		"brand_id":     req.BrandId,
		"category_id":  req.CategoryId,
		"model_year":   req.ModelYear,
		"list_price":   req.ListPrice,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *productRepo) Delete(ctx context.Context, req *models.ProductPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM products
		WHERE product_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.ProductId)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
