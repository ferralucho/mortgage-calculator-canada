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
	MortgageTotal           float64 `json:"mortgage_total"`
	MortgageBeforeChmc      float64 `json:"mortgage_before_chmc"`
	MortgagePaymentSchedule float64 `json:"mortgage_payment_schedule"`
	DifferenceRatio         float64 `json:"difference_ratio"`
	ChmcInsuranceTotal      float64 `json:"chmc_insurance_total"`
}

type PaymentSchedule string

const (
	AcceleratedBiWeekly PaymentSchedule = "accelerated_bi_weekly"
	BiWeekly            PaymentSchedule = "bi_weekly"
	Monthly             PaymentSchedule = "monthly"
)

func (input *CalculationInput) Validate() rest_errors.RestErr {
	if input.PropertyPrice <= 0 || input.PropertyPrice >= 1000000000 {
		return rest_errors.NewBadRequestError("invalid property price")
	}

	isValidDownPayment, err := validateDownPayment(input)
	if !isValidDownPayment {
		return err
	}

	if input.AnnualInterestRate <= 0 || input.AnnualInterestRate > 1000 {
		return rest_errors.NewBadRequestError("invalid annual interest rate")
	}

	isValidPeriod, err := validateAmortizationPeriod(input.AmortizationPeriod)
	if !isValidPeriod {
		return err
	}

	input.PaymentSchedule = strings.TrimSpace(strings.ToLower(input.PaymentSchedule))
	if input.PaymentSchedule == "" {
		return rest_errors.NewBadRequestError("invalid payment schedule")
	}

	return nil
}

func validateDownPayment(input *CalculationInput) (bool, rest_errors.RestErr) {
	differenceRatio := (input.DownPayment * 100) / input.PropertyPrice
	if input.DownPayment <= 0 || input.DownPayment >= 1000000000 || differenceRatio < 5 || input.DownPayment >= input.PropertyPrice || (input.PropertyPrice > 500000 && differenceRatio < 10) {
		return false, rest_errors.NewBadRequestError("invalid down payment")
	}
	return true, nil
}

func validateAmortizationPeriod(amortizationPeriod uint64) (bool, rest_errors.RestErr) {
	if amortizationPeriod == 0 || amortizationPeriod <= 5 || amortizationPeriod >= 30 || amortizationPeriod%5 != 0 {
		return false, rest_errors.NewBadRequestError("invalid amortization period")
	}
	return true, nil
}
