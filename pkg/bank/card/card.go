package card

import (
	"bank/pkg/bank/types"
)

func PaymentSource(cards []types.Card) []types.PaymentSource {
	// TODO: Code

	paymentSource := []types.PaymentSource{{Number: ""},}

	for _, card := range cards {
		
		if card.Balance < 0 {
			continue
		}
		if !card.Active {
			continue
		}

		paymentSource = []types.PaymentSource{{Number: string(card.PAN)},}
	}

	return paymentSource

}
// экземпляр PaymentSource и доб в пустой слайсс с PaymentSource