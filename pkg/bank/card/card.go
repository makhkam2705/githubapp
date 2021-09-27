package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func PaymentSource(cards []types.Card) []types.PaymentSource {
	// TODO: Code

	paymentSource := []types.PaymentSource{}
	for _, card := range cards {
		if !card.Active {
			continue
		}

		if card.Balance < 0 {
			continue
		}

		paymentSource = []types.PaymentSource{}
		fmt.Println(card)		
	}

	return paymentSource

}