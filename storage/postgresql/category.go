package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) Create(ctx context.Context, req *models.CreateCategory) (int, error) {
	var (
		query string
		id    int
	)
	
	query = `
		INSERT INTO categories(
			category_id, 
			category_name 
		)
		VALUES (
			(
				SELECT MAX(category_id) + 1 FROM categories
			),
			$1) RETURNING category_id
	`
	err := r.db.QueryRow(ctx, query,
		req.CategoryName,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *categoryRepo) GetByID(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error) {

	var (
		query    string
		category models.Category
	)

	query = `
		SELECT
			category_id,
			category_name
		FROM categories
		WHERE category_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.CategoryId).Scan(
		&category.CategoryId,
		&category.CategoryName,
	)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (resp *models.GetListCategoryResponse, err error) {

	resp = &models.GetListCategoryResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			category_id,
			category_name
		FROM categories
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
		var category models.Category
		err = rows.Scan(
			&resp.Count,
			&category.CategoryId,
			&category.CategoryName,
		)
		if err != nil {
			return nil, err
		}

		resp.Categories = append(resp.Categories, &category)
	}

	return resp, nil
}

func (r *categoryRepo) Update(ctx context.Context, req *models.UpdateCategory) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		categories
		SET
			category_id = :category_id, 
			category_name = :category_name
		WHERE category_id = :category_id
	`

	params = map[string]interface{}{
		"category_id":   req.CategoryId,
		"category_name": req.CategoryName,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *categoryRepo) Delete(ctx context.Context, req *models.CategoryPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM categories
		WHERE category_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.CategoryId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
