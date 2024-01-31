package product

import (
	"context"
	"database/sql"

	"github.com/dickidarmawansaputra/go-simple-online-shop/infra/response"
)

type repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateProduct(ctx context.Context, model Product) (err error) {
	query := "INSERT INTO products (name, sku, stock, price, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = r.db.ExecContext(ctx, query, model.Name, model.SKU, model.Stock, model.Price, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return
	}

	return
}

func (r repository) GetAllProductsWithPaginationCursor(ctx context.Context, model ProductPagination) (products []Product, err error) {
	query := "SELECT id, name, sku, stock, price, created_at, updated_at FROM products WHERE id > ? ORDER BY id ASC LIMIT ?"

	rows, err := r.db.QueryContext(ctx, query, model.Cursor, model.Size)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}

		return
	}
	defer rows.Close()

	for rows.Next() {
		product := Product{}
		rows.Scan(&product.Id, &product.Name, &product.SKU, &product.Stock, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		products = append(products, product)
	}

	return
}

func (r repository) GetDetailProductBySKU(ctx context.Context, sku string) (product Product, err error) {
	query := "SELECT * FROM products WHERE sku = ?"

	row := r.db.QueryRowContext(ctx, query, sku)
	if row.Err() != nil {
		return
	}

	err = row.Scan(
		&product.Id,
		&product.Name,
		&product.SKU,
		&product.Stock,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return
	}

	return
}
