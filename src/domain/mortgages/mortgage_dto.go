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
	if input.PropertyPrice <= 0 {
		return rest_errors.NewBadRequestError("invalid property price")
	}

	if input.DownPayment <= 0 {
		return rest_errors.NewBadRequestError("invalid down payment")
	}

	if input.AnnualInterestRate <= 0 {
		return rest_errors.NewBadRequestError("invalid annual interest rate")
	}

	if input.AmortizationPeriod == 0 {
		return rest_errors.NewBadRequestError("invalid amortization period")
	}

	input.PaymentSchedule = strings.TrimSpace(strings.ToLower(input.PaymentSchedule))
	if input.PaymentSchedule == "" {
		return rest_errors.NewBadRequestError("invalid payment schedule")
	}

	return nil
}
