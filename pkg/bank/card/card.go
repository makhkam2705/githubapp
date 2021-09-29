package card

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {
	// TODO: Code
	paymentSources := []types.PaymentSource{}

	for _, i := range cards {
		if i.Balance > 0 && i.Active {
			pCrds := types.PaymentSource{
				Type: "card",
				Number: string(i.PAN),
				Balance: i.Balance,
			}
			paymentSources = append(paymentSources, pCrds)
		}

	}
	return paymentSources

}
