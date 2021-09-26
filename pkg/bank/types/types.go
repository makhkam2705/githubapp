package types

// Money представляет собой денежную сумму в минимальных единацах
type Money int64

// Currency представляет кол валюты
type Currency string

// Код валяты
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN представляет номер карты
type PAN string

// Card представляет информацию о платёжной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money // использовать Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}

// Payment представляет информацию о платеже.
type Payment struct {
	ID     int
	Amount Money
}

// type PaymentSource
type PaymentSource struct {
	Type string // 'card'
	Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
}