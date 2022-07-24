package models

type Brand struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type Product struct {
	ID        int64  `json:"id"`
	BrandId   int32  `json:"brand_id"`
	Name      string `json:"name"`
	Price     int16  `json:"price"`
	CreatedAt string `json:"created_at"`
}

type Customer struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
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
	CustomerId    int64  `json:"customer_id"`
	TransactionId int64  `json:"transaction_id"`
	CreatedAt     string `json:"created_at"`
}
