package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance: -10,
			Active: true,
			PAN: "5002",
		},
		{
			Balance: 3,
			Active: true,
			PAN: "5001",
		},
		{
			Balance: 10,
			Active: true,
			PAN: "5003",
		},
	}

	payments := PaymentSources(cards)
	for _, payment := range payments {
		fmt.Println(payment.Number)
	}

	fmt.Println(payments)

	// Output:
	// 5001
	// 5003
	// [{card 5001 3} {card 5003 10}]
}