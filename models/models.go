package models

type Brand struct {
	ID        int64  `json:"id"`
	Name      string `json:"name" form:"name" validate:"required,max=255,min=3"`
	CreatedAt string `json:"created_at"`
}

type Product struct {
	ID        int64  `json:"id"`
	BrandId   int64  `json:"brandId" form:"brandId" validate:"required,numeric"`
	Name      string `json:"name" form:"name" validate:"required,max=255,min=3"`
	Price     int64  `json:"price" form:"price" validate:"required,numeric"`
	CreatedAt string `json:"created_at"`
}

type Customer struct {
	ID        int64  `json:"id"`
	Name      string `json:"name" form:"name" validate:"required,max=255,min=3"`
	CreatedAt string `json:"created_at"`
}

type Transaction struct {
	ID         int64  `json:"id"`
	CustomerId int64  `json:"customer_id"`
	Amount     string `json:"amount"`
	CreatedAt  string `json:"created_at"`
}

type Order struct {
	ID            int64  `json:"id"`
	ProductId     int64  `json:"product_id"`
	TransactionId int64  `json:"transaction_id"`
	CreatedAt     string `json:"created_at"`
}
