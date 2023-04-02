package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)

type stockRepo struct {
	db *pgxpool.Pool
}

func NewStockRepo(db *pgxpool.Pool) *stockRepo {
	return &stockRepo{
		db: db,
	}
}

func (r *stockRepo) Create(ctx context.Context, req *models.CreateStock) (int, int, error) {
	var (
		query     string
		storeId   int
		productId int
	)

	query = `
		INSERT INTO stocks(
			store_id,
			product_id,
			quantity
		)
		VALUES ($1, $2, $3) RETURNING store_id, product_id
	`
	err := r.db.QueryRow(ctx, query,
		req.StoreId,
		req.ProductId,
		req.Quantity,
	).Scan(&storeId, &productId)
	if err != nil {
		return 0, 0, err
	}

	return storeId, productId, nil
}

func (r *stockRepo) GetByID(ctx context.Context, req *models.StockPrimaryKey) (resp *models.GetStock, err error) {

	resp = &models.GetStock{}

	var (
		storeId  sql.NullInt64
		quantity sql.NullInt64
		products pgtype.JSONB
	)

	query := `
		SELECT
			s.store_id,
			SUM(s.quantity),
			JSONB_AGG (
				JSONB_BUILD_OBJECT (
					'product_id', p.product_id,
					'product_name', p.product_name,
					'brand_id', p.brand_id,
					'category_id', p.category_id,
					'model_year', p.model_year,
					'list_price', p.list_price,
					'quantity', s.quantity
				)
			) AS product_data
		FROM stocks AS s
		LEFT JOIN products AS p ON p.product_id = s.product_id
		WHERE s.store_id = $1
		GROUP BY s.store_id
	`
	err = r.db.QueryRow(ctx, query, req.StoreId).Scan(
		&storeId,
		&quantity,
		&products,
	)
	if err != nil {
		return nil, err
	}

	resp.StoreId = int(storeId.Int64)
	resp.Quantity = int(quantity.Int64)

	products.AssignTo(&resp.Products)

	return resp, nil
}

func (r *stockRepo) GetList(ctx context.Context, req *models.GetListStockRequest) (resp *models.GetListStockResponse, err error) {

	resp = &models.GetListStockResponse{}

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
			ARRAY_AGG(product_id),
			ARRAY_AGG(quantity)
		FROM stocks
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

	query += filter + " GROUP BY store_id " + offset + limit
	fmt.Println(query)

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			stock      models.GetStock
			productIds []sql.NullInt64
			amounts    []sql.NullInt64
		)

		err = rows.Scan(
			&resp.Count,
			&stock.StoreId,
			pq.Array(&productIds),
			pq.Array(&amounts),
		)
		if err != nil {
			return nil, err
		}

		for i, id := range productIds {
			data := models.ProductData{
				ProductId: int(id.Int64),
				Quantity:  int(amounts[i].Int64),
			}
			stock.Products = append(stock.Products, &data)
		}

		resp.Stocks = append(resp.Stocks, &stock)
	}

	return resp, nil
}

func (r *stockRepo) Update(ctx context.Context, req *models.UpdateStock) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		stocks
		SET
			quantity = :quantity
		WHERE store_id = :store_id AND product_id = :product_id
	`

	params = map[string]interface{}{
		"store_id":   req.StoreId,
		"product_id": req.ProductId,
		"quantity":   req.Quantity,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *stockRepo) Delete(ctx context.Context, req *models.StockPrimaryKey) (int64, error) {

	var (
		storeId string
	)
	if req.StoreId > 0 {
		storeId = fmt.Sprintf(" store_id = %d ", req.StoreId)
	}

	query := `
		DELETE
		FROM stocks
		WHERE 
	` + storeId

	result, err := r.db.Exec(ctx, query, req.StoreId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *stockRepo) SendProduct(ctx context.Context, req *models.SendProduct) error {
	var (
		senderStock int
	)

	if req.Quantity <= 0 {
		return errors.New("Invalid quantity")
	}

	err := r.db.QueryRow(ctx,
		`SELECT quantity FROM stocks WHERE store_id = $1 AND product_id = $2`,
		req.SenderId,
		req.ProductId,
	).Scan(&senderStock)
	if err != nil {
		return err
	}

	if senderStock < req.Quantity {
		return errors.New("Sender doesn't have enough of this product")
	}

	_, err = r.db.Exec(ctx,
		`UPDATE stocks SET quantity = quantity - $1 WHERE store_id = $2 AND product_id = $3`,
		req.Quantity,
		req.SenderId,
		req.ProductId,
	)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx,
		`UPDATE stocks SET quantity = quantity + $1 WHERE store_id = $2 AND product_id = $3`,
		req.Quantity,
		req.ReceiverId,
		req.ProductId,
	)
	if err != nil {
		return err
	}

	return nil
}
