package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type customerRepo struct {
	db *pgxpool.Pool
}

func NewCustomerRepo(db *pgxpool.Pool) *customerRepo {
	return &customerRepo{
		db: db,
	}
}

func (r *customerRepo) Create(ctx context.Context, req *models.CreateCustomer) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO customers(
			customer_id, 
			first_name,
			last_name,
			phone,
			email,
			street,
			city,
			state,
			zip_code
		)
		VALUES (
			(
				SELECT MAX(customer_id) + 1 FROM customers
			),
			$1, $2, $3, $4, $5, $6, $7, $8) RETURNING customer_id
	`
	err := r.db.QueryRow(ctx, query,
		req.FirstName,
		req.LastName,
		helper.NewNullString(req.Phone),
		req.Email,
		helper.NewNullString(req.Street),
		helper.NewNullString(req.City),
		helper.NewNullString(req.State),
		helper.NewNullInt32(req.ZipCode),
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *customerRepo) GetByID(ctx context.Context, req *models.CustomerPrimaryKey) (*models.Customer, error) {

	var (
		query    string
		customer models.Customer
	)

	query = `
		SELECT
			customer_id, 
			first_name,
			last_name,
			COALESCE(phone, ''),
			email,
			COALESCE(street, ''),
			COALESCE(city, ''),
			COALESCE(state, ''),
			COALESCE(zip_code, 0)
		FROM customers
		WHERE customer_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.CustomerId).Scan(
		&customer.CustomerId,
		&customer.FirstName,
		&customer.LastName,
		&customer.Phone,
		&customer.Email,
		&customer.Street,
		&customer.City,
		&customer.State,
		&customer.ZipCode,
	)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *customerRepo) GetList(ctx context.Context, req *models.GetListCustomerRequest) (resp *models.GetListCustomerResponse, err error) {

	resp = &models.GetListCustomerResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			customer_id, 
			first_name,
			last_name,
			COALESCE(phone, ''),
			email,
			COALESCE(street, ''),
			COALESCE(city, ''),
			COALESCE(state, ''),
			COALESCE(zip_code, 0)
		FROM customers
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
		var customer models.Customer
		err = rows.Scan(
			&resp.Count,
			&customer.CustomerId,
			&customer.FirstName,
			&customer.LastName,
			&customer.Phone,
			&customer.Email,
			&customer.Street,
			&customer.City,
			&customer.State,
			&customer.ZipCode,
		)
		if err != nil {
			return nil, err
		}

		resp.Customers = append(resp.Customers, &customer)
	}

	return resp, nil
}

func (r *customerRepo) UpdatePut(ctx context.Context, req *models.UpdateCustomer) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		customers
		SET
			customer_id =:customer_id, 
			first_name =:first_name,
			last_name =:last_name,
			phone = :phone,
			email = :email,
			street = :street,
			city = :city,
			state = :state,
			zip_code = :zip_code
		WHERE customer_id = :customer_id
	`

	params = map[string]interface{}{
		"customer_id": req.CustomerId,
		"first_name":  req.FirstName,
		"last_name":   req.LastName,
		"phone":       helper.NewNullString(req.Phone),
		"email":       req.Email,
		"street":      helper.NewNullString(req.Street),
		"city":        helper.NewNullString(req.City),
		"state":       helper.NewNullString(req.State),
		"zip_code":    helper.NewNullInt32(req.ZipCode),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *customerRepo) UpdatePatch(ctx context.Context, req *models.PatchRequest) (int64, error) {
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
		customers
		SET
		` + set + `
		WHERE customer_id = :customer_id
	`

	req.Fields["customer_id"] = req.ID

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	fmt.Println(query)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *customerRepo) Delete(ctx context.Context, req *models.CustomerPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM customers
		WHERE customer_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.CustomerId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
