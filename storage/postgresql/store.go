package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type storeRepo struct {
	db *pgxpool.Pool
}

func NewStoreRepo(db *pgxpool.Pool) *storeRepo {
	return &storeRepo{
		db: db,
	}
}

func (r *storeRepo) Create(ctx context.Context, req *models.CreateStore) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO stores(
			store_id, 
			store_name,
			phone,
			email,
			street,
			city,
			state,
			zip_code
		)
		VALUES (
			(
				SELECT MAX(store_id) + 1 FROM stores
			),
			$1, $2, $3, $4, $5, $6, $7) RETURNING store_id
	`
	err := r.db.QueryRow(ctx, query,
		req.StoreName,
		helper.NewNullString(req.Phone),
		helper.NewNullString(req.Email),
		helper.NewNullString(req.Street),
		helper.NewNullString(req.City),
		helper.NewNullString(req.State),
		helper.NewNullString(req.ZipCode),
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *storeRepo) GetByID(ctx context.Context, req *models.StorePrimaryKey) (*models.Store, error) {

	var (
		query string
		store models.Store
	)

	query = `
		SELECT
			store_id, 
			store_name,
			COASLESCE(phone, ''),
			COASLESCE(email, ''),
			COASLESCE(street, ''),
			COASLESCE(city, ''),
			COASLESCE(state, ''),
			COASLESCE(zip_code, '')
		FROM stores
		WHERE store_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.StoreId).Scan(
		&store.StoreId,
		&store.StoreName,
		&store.Phone,
		&store.Email,
		&store.Street,
		&store.City,
		&store.State,
		&store.ZipCode,
	)
	if err != nil {
		return nil, err
	}

	return &store, nil
}

func (r *storeRepo) GetList(ctx context.Context, req *models.GetListStoreRequest) (resp *models.GetListStoreResponse, err error) {

	resp = &models.GetListStoreResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			store_id, 
			store_name,
			COALESCE(phone, ''),
			COALESCE(email, ''),
			COALESCE(street, ''),
			COALESCE(city, ''),
			COALESCE(state, ''),
			COALESCE(zip_code, '')
		FROM stores
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
		var store models.Store
		err = rows.Scan(
			&resp.Count,
			&store.StoreId,
			&store.StoreName,
			&store.Phone,
			&store.Email,
			&store.Street,
			&store.City,
			&store.State,
			&store.ZipCode,
		)
		if err != nil {
			return nil, err
		}

		resp.Stores = append(resp.Stores, &store)
	}

	return resp, nil
}

func (r *storeRepo) UpdatePut(ctx context.Context, req *models.UpdateStore) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		stores
		SET
			store_id = :store_id, 
			store_name = :store_name,
			phone = :phone,
			email = :email,
			street = :street,
			city = :city,
			state = :state,
			zip_code = :zip_code
		WHERE store_id = :store_id
	`

	params = map[string]interface{}{
		"store_id":   req.StoreId,
		"store_name": req.StoreName,
		"phone":      helper.NewNullString(req.Phone),
		"email":      helper.NewNullString(req.Email),
		"street":     helper.NewNullString(req.Street),
		"city":       helper.NewNullString(req.City),
		"state":      helper.NewNullString(req.State),
		"zip_code":   helper.NewNullString(req.ZipCode),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *storeRepo) UpdatePatch(ctx context.Context, req *models.PatchRequest) (int64, error) {
	var (
		query string
		set   string
	)

	if len(req.Fields) <= 0 {
		return 0, errors.New("no fields")
	}

	i := 0
	for key := range req.Fields {
		if i == len(req.Fields)-1 {
			set += fmt.Sprintf(" %s = :%s ", key, key)
		} else {
			set += fmt.Sprintf(" %s = :%s, ", key, key)
		}
		i++
	}

	query = `
		UPDATE
		stores
		SET
		` + set + `
		WHERE store_id = :store_id
	`

	req.Fields["store_id"] = req.ID

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	fmt.Println(query)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *storeRepo) Delete(ctx context.Context, req *models.StorePrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM stores
		WHERE store_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.StoreId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
