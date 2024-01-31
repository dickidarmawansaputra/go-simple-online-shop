package transaction

import (
	"context"
	"database/sql"
	"log"

	"github.com/dickidarmawansaputra/go-simple-online-shop/infra/response"
)

type Repository interface {
	TransactionDBRepository
	TransactionRepository
	ProductRepository
}

type TransactionDBRepository interface {
	Begin(ctx context.Context) (tx *sql.Tx, err error)
	Rollback(ctx context.Context, tx *sql.Tx) (err error)
	Commit(ctx context.Context, tx *sql.Tx) (err error)
}

type TransactionRepository interface {
	CreateTransactionWithTx(ctx context.Context, tx *sql.Tx, trx Transaction) (err error)
}

type ProductRepository interface {
	GetProductBySKU(ctx context.Context, productSKU string) (product Product, err error)
	UpdateProductStockWithTx(ctx context.Context, tx *sql.Tx, product Product) (err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) CreateTransaction(ctx context.Context, req CreateTransactionRequestPayload) (err error) {
	myProduct, err := s.repo.GetProductBySKU(ctx, req.ProductSKU)
	if err != nil {
		return
	}

	if !myProduct.IsExists() {
		err = response.ErrNotFound
		return
	}

	trx := NewTransactionFromCreateTransactionRequest(req)
	trx.FromProduct(myProduct).
		SetPlatformFee(5000).
		SetGrandTotal()

	// bisa cara gini juga
	// trx.SetGrandTotal()
	// trx.SetPlatformFee(5000)

	if err = trx.Validate(); err != nil {
		return
	}

	if err = trx.ValidateStock(uint8(myProduct.Stock)); err != nil {
		return
	}

	// start transaction
	tx, err := s.repo.Begin(ctx)
	if err != nil {
		return
	}

	// defer transaction if any error or after commit
	defer s.repo.Rollback(ctx, tx)

	if err = s.repo.CreateTransactionWithTx(ctx, tx, trx); err != nil {
		log.Println("ERROR DISINI")
		return
	}
	log.Println("TIDAK ERROR DISINI")

	// update current stock
	if err = myProduct.UpdateStockProduct(trx.Amount); err != nil {
		return
	}

	// update into database
	if err = s.repo.UpdateProductStockWithTx(ctx, tx, myProduct); err != nil {
		return
	}

	// commit to end transaction
	if err = s.repo.Commit(ctx, tx); err != nil {
		return
	}

	return
}
