package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSource() {
	fmt.Println(PaymentSource([]types.Card{
		{
			Balance: 1_000,
			Active: true,
			PAN: "5000 xxxx xxxx 5000",
		},
	}))
	fmt.Println(PaymentSource([]types.Card{
		{
			Balance: 1_000,
			Active: true,
			PAN: "5000 xxxx xxxx 5000",
		},
	}))
	fmt.Println(PaymentSource([]types.Card{
		{
			Balance: 1_000,
			Active: true,
			PAN: "5000 xxxx xxxx 5000",
		},
	}))


	// Output:
	// {0 5000 xxxx xxxx 5000 1000    true 0}
	// []
	// {0 5000 xxxx xxxx 5000 1000    true 0}
	// []
	// {0 5000 xxxx xxxx 5000 1000    true 0}
	// []
}