package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type codeRepo struct {
	db *pgxpool.Pool
}

func NewCodeRepo(db *pgxpool.Pool) *codeRepo {
	return &codeRepo{
		db: db,
	}
}

func (r *codeRepo) Create(ctx context.Context, req *models.CreateCode) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO promo_code(
			code_id, 
			code_name, 
			discount,
			discount_type,
			order_limit_price
		)
		VALUES (
			(
				SELECT COALESCE(MAX(code_id), 0) + 1 FROM promo_code
			)
			, $1, $2, $3, $4) RETURNING code_id
	`

	err := r.db.QueryRow(ctx, query,

		req.CodeName,
		req.Discount,
		req.DiscountType,
		req.OrderLimitPrice,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *codeRepo) GetByID(ctx context.Context, req *models.CodePrimaryKey) (*models.Code, error) {

	var (
		query string
		code  models.Code
	)

	query = `
		SELECT
			COALESCE(code_id, 0), 
			code_name,
			discount,
			discount_type,
			order_limit_price
		FROM promo_code
		WHERE code_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Code_Id).Scan(
		&code.Code_Id,
		&code.CodeName,
		&code.Discount,
		&code.DiscountType,
		&code.OrderLimitPrice,
	)
	if err != nil {
		return nil, err
	}

	return &code, nil
}

func (r *codeRepo) GetList(ctx context.Context, req *models.GetListCodeRequest) (resp *models.GetListCodeResponse, err error) {

	resp = &models.GetListCodeResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			code_id, 
			code_name,
			discount,
			discount_type,
			order_limit_price
		FROM promo_code
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
		var code models.Code
		err = rows.Scan(
			&resp.Count,
			&code.Code_Id,
			&code.CodeName,
			&code.Discount,
			&code.DiscountType,
			&code.OrderLimitPrice,
		)
		if err != nil {
			return nil, err
		}

		resp.Codes = append(resp.Codes, &code)
	}

	return resp, nil
}

func (r *codeRepo) Update(ctx context.Context, req *models.UpdateCode) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		promo_code
		SET
			code_id = :code_id, 
			code_name = :code_name,
			discount = :discount,
			discount_type = :discount_type,
			order_limit_price = :order_limit_price
		WHERE code_id = :code_id
	`

	params = map[string]interface{}{
		"code_id":           req.Code_Id,
		"code_name":         req.CodeName,
		"discount":          req.Discount,
		"discount_type":     req.DiscountType,
		"order_limit_price": req.OrderLimitPrice,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *codeRepo) Delete(ctx context.Context, req *models.CodePrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM promo_code
		WHERE code_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.Code_Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
