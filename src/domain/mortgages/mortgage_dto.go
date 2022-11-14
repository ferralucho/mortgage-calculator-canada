package mortgages

import (
	"strings"

	"github.com/ferralucho/mortgage-calculator-canada/src/rest_errors"
)

type CalculationInput struct {
	PropertyPrice      float64 `json:"property_price"`
	DownPayment        float64 `json:"down_payment"`
	AnnualInterestRate float64 `json:"annual_interest_rate"`
	AmortizationPeriod uint64  `json:"amortization_period"`
	PaymentSchedule    string  `json:"payment_schedule"`
}

type CalculationOutput struct {
	TotalMortgageTotal float64 `json:"total_mortgage_total"`
	MortgagePayment    float64 `json:"mortgage_payment"`
}

func (input *CalculationInput) Validate() rest_errors.RestErr {
	input.PaymentSchedule = strings.TrimSpace(input.PaymentSchedule)

	input.PaymentSchedule = strings.TrimSpace(strings.ToLower(input.PaymentSchedule))
	if input.PaymentSchedule == "" {
		return rest_errors.NewBadRequestError("invalid payment schedule")
	}

	return nil
}
