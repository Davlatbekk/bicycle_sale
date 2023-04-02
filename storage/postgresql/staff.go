package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type staffRepo struct {
	db *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) *staffRepo {
	return &staffRepo{
		db: db,
	}
}

func (r *staffRepo) Create(ctx context.Context, req *models.CreateStaff) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO staffs(
			staff_id, 
			first_name,
			last_name,
			email,
			phone,
			active,
			store_id,
			manager_id
		)
		VALUES (
			(
				SELECT MAX(staff_id) + 1 FROM staffs
			),
			$1, $2, $3, $4, $5, $6, $7) RETURNING staff_id
	`
	err := r.db.QueryRow(ctx, query,
		req.FirstName,
		req.LastName,
		req.Email,
		helper.NewNullString(req.Phone),
		req.Active,
		req.StoreId,
		helper.NewNullInt32(req.ManagerId),
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *staffRepo) GetByID(ctx context.Context, req *models.StaffPrimaryKey) (*models.Staff, error) {

	var (
		query string
		staff models.Staff
	)

	query = `
		SELECT
			s1.staff_id, 
			s1.first_name,
			s1.last_name,
			s1.email,
			COALESCE(s1.phone, ''),
			s1.active,
			s1.store_id,
			
			stores.store_id,
			stores.store_name,
			COALESCE(stores.phone, ''),
			COALESCE(stores.email, ''),
			COALESCE(stores.street, ''),
			COALESCE(stores.city, ''),
			COALESCE(stores.state, ''),
			COALESCE(stores.zip_code, ''),

			COALESCE(s1.manager_id, 0),

			s2.staff_id, 
			s2.first_name,
			s2.last_name,
			s2.email,

			COALESCE(s2.phone, ''),
			s2.active,
			s2.store_id,
			COALESCE(s2.manager_id, 0)

			FROM staffs AS s1
		JOIN stores ON stores.store_id = s1.store_id
		INNER JOIN staffs AS s2 ON COALESCE(s1.manager_id, s1.staff_id) = s2.staff_id
		WHERE s1.staff_id = $1;
	`

	staff.StoreData = &models.Store{}
	staff.ManagerData = &models.Staff{}

	err := r.db.QueryRow(ctx, query, req.StaffId).Scan(
		&staff.StaffId,
		&staff.FirstName,
		&staff.LastName,
		&staff.Email,
		&staff.Phone,
		&staff.Active,
		&staff.StoreId,

		&staff.StoreData.StoreId,
		&staff.StoreData.StoreName,
		&staff.StoreData.Phone,
		&staff.StoreData.Email,
		&staff.StoreData.Street,
		&staff.StoreData.City,
		&staff.StoreData.State,
		&staff.StoreData.ZipCode,
		&staff.ManagerId,

		&staff.ManagerData.StaffId,
		&staff.ManagerData.FirstName,
		&staff.ManagerData.LastName,
		&staff.ManagerData.Email,
		&staff.ManagerData.Phone,
		&staff.ManagerData.Active,
		&staff.ManagerData.StoreId,
		&staff.ManagerData.ManagerId,
	)
	if err != nil {
		return nil, err
	}

	return &staff, nil
}

func (r *staffRepo) GetList(ctx context.Context, req *models.GetListStaffRequest) (resp *models.GetListStaffResponse, err error) {

	resp = &models.GetListStaffResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)
	// bug if manager_id null then get dont work

	query = `
		SELECT
			COUNT(*) OVER(),
			s1.staff_id, 
			s1.first_name,
			s1.last_name,
			s1.email,
			COALESCE(s1.phone, ''),
			s1.active,
			s1.store_id,

			stores.store_id,
			stores.store_name,
			COALESCE(stores.phone, ''),
			COALESCE(stores.email, ''),
			COALESCE(stores.street, ''),
			COALESCE(stores.city, ''),
			COALESCE(stores.state, ''),
			COALESCE(stores.zip_code, ''),

			COALESCE(s1.manager_id, 0),

			s2.staff_id, 
			s2.first_name,
			s2.last_name,
			s2.email,

			COALESCE(s2.phone, ''),
			s2.active,
			s2.store_id,
			COALESCE(s2.manager_id, 0)
		FROM staffs AS s1
		JOIN stores ON stores.store_id = s1.store_id
		INNER JOIN staffs AS s2 ON COALESCE(s1.manager_id, s1.staff_id)= s2.staff_id
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
		var staff models.Staff
		staff.ManagerData = &models.Staff{}

		staff.StoreData = &models.Store{}
		err = rows.Scan(
			&resp.Count,
			&staff.StaffId,
			&staff.FirstName,
			&staff.LastName,
			&staff.Email,
			&staff.Phone,
			&staff.Active,
			&staff.StoreId,
			&staff.StoreData.StoreId,
			&staff.StoreData.StoreName,
			&staff.StoreData.Phone,
			&staff.StoreData.Email,
			&staff.StoreData.Street,
			&staff.StoreData.City,
			&staff.StoreData.State,
			&staff.StoreData.ZipCode,
			&staff.ManagerId,
			&staff.ManagerData.StaffId,
			&staff.ManagerData.FirstName,
			&staff.ManagerData.LastName,
			&staff.ManagerData.Email,
			&staff.ManagerData.Phone,
			&staff.ManagerData.Active,
			&staff.ManagerData.StoreId,
			&staff.ManagerData.ManagerId,
		)
		if err != nil {
			return nil, err
		}

		resp.Staffs = append(resp.Staffs, &staff)
	}

	return resp, nil
}

func (r *staffRepo) GetListReport(ctx context.Context, req *models.GetListReportStaffRequest) (resp *models.GetListReportStaffResponse, err error) {

	resp = &models.GetListReportStaffResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)
	// bug if manager_id null then get dont work

	query = `
		SELECT
		COUNT(*) OVER(),

			s1.first_name || ' ' || s1.last_name as full_name,
			c.category_name as category,
			p.product_name as product,
			oi.quantity as count,
			(oi.list_price * oi.quantity) as total_summ,
			CAST(o.order_date::timestamp AS VARCHAR)
			
		FROM staffs AS s1
		JOIN orders o  ON o.staff_id = s1.staff_id
		JOIN order_items oi  ON oi.order_id = o.order_id
		JOIN products p  ON p.product_id = oi.product_id
		JOIN categories c  ON c.category_id = p.category_id
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
		var report models.Report
		err = rows.Scan(
			&resp.Count,
			&report.FullName,
			&report.Category,
			&report.Product,
			&report.Count,
			&report.TotalSumm,
			&report.Date,
		)
		if err != nil {
			fmt.Println("ok")
			return nil, err
		}

		resp.Reports = append(resp.Reports, &report)
	}

	return resp, nil
}

func (r *staffRepo) UpdatePut(ctx context.Context, req *models.UpdateStaff) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		staffs
		SET
			staff_id = :staff_id, 
			first_name = :first_name,
			last_name = :last_name,
			email = :email,
			phone = :phone,
			active = :active,
			store_id = :store_id,
			manager_id = :manager_id
		WHERE staff_id = :staff_id
	`

	params = map[string]interface{}{
		"staff_id":   req.StaffId,
		"first_name": req.FirstName,
		"last_name":  req.LastName,
		"email":      req.Email,
		"phone":      helper.NewNullString(req.Phone),
		"active":     req.Active,
		"store_id":   req.StoreId,
		"manager_id": helper.NewNullInt32(req.ManagerId),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *staffRepo) UpdatePatch(ctx context.Context, req *models.PatchRequest) (int64, error) {
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
		staffs
		SET
		` + set + `
		WHERE staff_id = :staff_id
	`

	req.Fields["staff_id"] = req.ID

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	fmt.Println(query)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *staffRepo) Delete(ctx context.Context, req *models.StaffPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM staffs
		WHERE staff_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.StaffId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
