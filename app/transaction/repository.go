package transaction

import (
	"context"
	"database/sql"
	"log"
)

func newRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}

type repository struct {
	db *sql.DB
}

func (r repository) Begin(ctx context.Context) (tx *sql.Tx, err error) {
	tx, err = r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}

	return
}

func (r repository) Commit(ctx context.Context, tx *sql.Tx) (err error) {
	return tx.Commit()
}

func (r repository) Rollback(ctx context.Context, tx *sql.Tx) (err error) {
	return tx.Rollback()
}

func (r repository) CreateTransactionWithTx(ctx context.Context, tx *sql.Tx, trx Transaction) (err error) {
	query := "INSERT INTO transactions (email, product_id, product_price, amount, sub_total, platform_fee, grand_total, status, product_snapshot, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	log.Println(trx)
	_, err = stmt.ExecContext(
		ctx,
		trx.Email,
		trx.ProductId,
		trx.ProductPrice,
		trx.Amount,
		trx.SubTotal,
		trx.PlatformFee,
		trx.GrandTotal,
		trx.Status,
		trx.ProductJSON,
		trx.CreatedAt,
		trx.UpdatedAt,
	)
	if err != nil {
		log.Println("ERROR QUERY DISINI")
		log.Println(err.Error())
		return
	}
	log.Println("TIDAK ERROR QUERY DISINI")

	return
}

func (r repository) GetProductBySKU(ctx context.Context, productSKU string) (product Product, err error) {
	query := "SELECT id, sku, name, stock, price FROM products WHERE sku = ?"

	row := r.db.QueryRowContext(ctx, query, productSKU)
	if row.Err() != nil {
		return
	}

	err = row.Scan(
		&product.Id,
		&product.SKU,
		&product.Name,
		&product.Stock,
		&product.Price,
	)
	if err != nil {
		return
	}

	return
}

func (r repository) UpdateProductStockWithTx(ctx context.Context, tx *sql.Tx, product Product) (err error) {
	query := "UPDATE products SET stock = ? WHERE id = ?"

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, product.Stock, product.Id)
	if err != nil {
		return
	}

	return
}
