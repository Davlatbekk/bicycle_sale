package postgresql

import (
	"app/config"
	"app/storage"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db       *pgxpool.Pool
	brand    storage.BrandRepoI
	product  storage.ProductRepoI
	category storage.CategoryRepoI
	stock    storage.StockRepoI
	stores   storage.StoreRepoI
	customer storage.CustomerRepoI
	staff    storage.StaffRepoI
	order    storage.OrderRepoI
	code     storage.CodeRepoI
}

func NewConnectPostgresql(cfg *config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))
	if err != nil {
		return nil, err
	}

	pgpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:       pgpool,
		product:  NewProductRepo(pgpool),
		category: NewCategoryRepo(pgpool),
		brand:    NewBrandRepo(pgpool),
		stock:    NewStockRepo(pgpool),
		stores:   NewStoreRepo(pgpool),
		customer: NewCustomerRepo(pgpool),
		staff:    NewStaffRepo(pgpool),
		order:    NewOrderRepo(pgpool),
		code:     NewCodeRepo(pgpool),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Brand() storage.BrandRepoI {
	if s.brand == nil {
		s.brand = NewBrandRepo(s.db)
	}

	return s.brand
}

func (s *Store) Product() storage.ProductRepoI {
	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}

func (s *Store) Category() storage.CategoryRepoI {
	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}

func (s *Store) Stock() storage.StockRepoI {
	if s.stock == nil {
		s.stock = NewStockRepo(s.db)
	}

	return s.stock
}

func (s *Store) Store() storage.StoreRepoI {
	if s.stores == nil {
		s.stores = NewStoreRepo(s.db)
	}

	return s.stores
}

func (s *Store) Customer() storage.CustomerRepoI {
	if s.customer == nil {
		s.customer = NewCustomerRepo(s.db)
	}

	return s.customer
}

func (s *Store) Staff() storage.StaffRepoI {
	if s.staff == nil {
		s.staff = NewStaffRepo(s.db)
	}

	return s.staff
}

func (s *Store) Order() storage.OrderRepoI {
	if s.order == nil {
		s.order = NewOrderRepo(s.db)
	}

	return s.order
}

func (s *Store) Code() storage.CodeRepoI {
	if s.code == nil {
		s.code = NewCodeRepo(s.db)
	}

	return s.code
}
