package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance: 110000000000,
			Active: true,
			PAN: "5003 xxxx xxxx 0000",
		},
		{
			Balance: 3,
			Active: true,
			PAN: "5001 xxxx xxxx 0000",
		},
		{
			Balance: 10,
			Active: false,
			PAN: "5003 xxxx xxxx 0000",
		},
	}

	payments := PaymentSources(cards)
	for _, payment := range payments {
		fmt.Println(payment.Number)
	}

	fmt.Println(payments)

	// Output:
	// 5001 xxxx xxxx 0000
	// [{ 5001 xxxx xxxx 0000 0}]
}