package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type brandRepo struct {
	db *pgxpool.Pool
}

func NewBrandRepo(db *pgxpool.Pool) *brandRepo {
	return &brandRepo{
		db: db,
	}
}

func (r *brandRepo) Create(ctx context.Context, req *models.CreateBrand) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO brands(
			brand_id, 
			brand_name 
		)
		VALUES (
			(
				SELECT MAX(brand_id) + 1 FROM brands
			),
			$1) RETURNING brand_id
	`
	err := r.db.QueryRow(ctx, query,
		req.BrandName,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *brandRepo) GetByID(ctx context.Context, req *models.BrandPrimaryKey) (*models.Brand, error) {

	var (
		query string
		brand models.Brand
	)

	query = `
		SELECT
			brand_id, 
			brand_name 
		FROM brands
		WHERE brand_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.BrandId).Scan(
		&brand.BrandId,
		&brand.BrandName,
	)
	if err != nil {
		return nil, err
	}

	return &brand, nil
}

func (r *brandRepo) GetList(ctx context.Context, req *models.GetListBrandRequest) (resp *models.GetListBrandResponse, err error) {

	resp = &models.GetListBrandResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			brand_id,
			brand_name
		FROM brands
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
		var brand models.Brand
		err = rows.Scan(
			&resp.Count,
			&brand.BrandId,
			&brand.BrandName,
		)
		if err != nil {
			return nil, err
		}

		resp.Brands = append(resp.Brands, &brand)
	}

	return resp, nil
}

func (r *brandRepo) Update(ctx context.Context, req *models.UpdateBrand) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		brands
		SET
			brand_id = :brand_id, 
			brand_name = :brand_name
		WHERE brand_id = :brand_id
	`

	params = map[string]interface{}{
		"brand_id":   req.BrandId,
		"brand_name": req.BrandName,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *brandRepo) Delete(ctx context.Context, req *models.BrandPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM brands
		WHERE brand_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.BrandId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
