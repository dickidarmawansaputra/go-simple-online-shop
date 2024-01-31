package transaction

import (
	"encoding/json"
	"time"

	"github.com/dickidarmawansaputra/go-simple-online-shop/infra/response"
)

type TransactionStatus uint8

const (
	TransactionStatusCreated    TransactionStatus = 1
	TransactionStatusProgress   TransactionStatus = 10
	TransactionStatusInDelivery TransactionStatus = 15
	TransactionStatusCompleted  TransactionStatus = 20

	TRX_CREATED      string = "CREATED"
	TRX_IN_PROGRESS  string = "IN_PROGRESS"
	TRX_IN_DELIVERY  string = "IN_DELIVERY"
	TRX_IN_COMPLETED string = "COMPLETED"
	TRX_UNKNOWN      string = "UNKNOWN"
)

var (
	MappingTransactionStatus = map[TransactionStatus]string{
		TransactionStatusCreated:    TRX_CREATED,
		TransactionStatusProgress:   TRX_IN_PROGRESS,
		TransactionStatusInDelivery: TRX_IN_DELIVERY,
		TransactionStatusCompleted:  TRX_IN_COMPLETED,
	}
)

type Transaction struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	ProductId    uint   `json:"product_id"`
	ProductPrice uint   `json:"product_price"`
	Amount       uint8  `json:"amount"`
	// price x amount = sub total
	SubTotal    uint `json:"sub_total"`
	PlatformFee uint `json:"platform_fee"`
	// sub total x addtional fee (platform fee) = grand total
	GrandTotal  uint              `json:"grand_total"`
	Status      TransactionStatus `json:"status"`
	ProductJSON json.RawMessage   `json:"product_snapshot"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

func NewTransaction(email string) Transaction {
	return Transaction{
		Email:     email,
		Status:    TransactionStatusCreated,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewTransactionFromCreateTransactionRequest(req CreateTransactionRequestPayload) Transaction {
	return Transaction{
		Email:     req.Email,
		Status:    TransactionStatusCreated,
		Amount:    req.Amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t *Transaction) Validate() (err error) {
	if t.Amount == 0 {
		return response.ErrAmountInvalid
	}

	return
}

func (t Transaction) ValidateStock(productStock uint8) (err error) {
	if t.Amount > productStock {
		return response.ErrAmountGreaterThanStock
	}
	return
}

func (t *Transaction) SetSubTotal() {
	if t.SubTotal == 0 {
		t.SubTotal = t.ProductPrice * uint(t.Amount)
	}
}

func (t *Transaction) SetPlatformFee(fee uint) *Transaction {
	t.PlatformFee = fee

	return t
}

func (t *Transaction) SetGrandTotal() {
	if t.GrandTotal == 0 {
		t.SetSubTotal()

		t.GrandTotal = t.SubTotal + t.PlatformFee
	}

	return
}

func (t *Transaction) FromProduct(product Product) *Transaction {
	t.ProductId = uint(product.Id)
	t.ProductPrice = uint(product.Price)

	return t
}

// set product id, price, & json
func (t *Transaction) SetProductJSON(product Product) (err error) {
	productJSON, err := json.Marshal(product)
	if err != nil {
		return
	}

	t.ProductJSON = productJSON

	return
}

func (t *Transaction) GetProduct() (product Product, err error) {
	err = json.Unmarshal(t.ProductJSON, &product)
	return
}

func (t *Transaction) GetStatus() string {
	status, ok := MappingTransactionStatus[t.Status]
	if !ok {
		return TRX_UNKNOWN
	}

	return status
}
