package mortgages

import (
	"reflect"
	"testing"

	"github.com/ferralucho/mortgage-calculator-canada/src/rest_errors"
)

func TestCalculationInput_Validate(t *testing.T) {
	tests := []struct {
		name  string
		input *CalculationInput
		want  rest_errors.RestErr
	}{
		{
			"no validation error",
			&CalculationInput{
				PropertyPrice:      900000,
				DownPayment:        180000,
				AnnualInterestRate: 5.19,
				AmortizationPeriod: 25,
				PaymentSchedule:    string(Monthly),
			},
			nil,
		},
		{
			"little downpayment no validation error",
			&CalculationInput{
				PropertyPrice:      100000,
				DownPayment:        100,
				AnnualInterestRate: 1,
				AmortizationPeriod: 10,
				PaymentSchedule:    string(Monthly),
			},
			nil,
		},
		{
			"down payment is bigger than property price, should throw an error",
			&CalculationInput{
				PropertyPrice:      100,
				DownPayment:        100000,
				AnnualInterestRate: 1,
				AmortizationPeriod: 10,
				PaymentSchedule:    string(Monthly),
			},
			rest_errors.NewBadRequestError("invalid down payment"),
		},
		{
			"down payment is less than 5 percent, should throw an error",
			&CalculationInput{
				PropertyPrice:      100,
				DownPayment:        1,
				AnnualInterestRate: 1,
				AmortizationPeriod: 10,
				PaymentSchedule:    string(Monthly),
			},
			rest_errors.NewBadRequestError("invalid down payment"),
		},
		{
			"annual interest rate is 0, should throw an error",
			&CalculationInput{
				PropertyPrice:      100000,
				DownPayment:        4300,
				AnnualInterestRate: 0,
				AmortizationPeriod: 10,
				PaymentSchedule:    string(Monthly),
			},
			rest_errors.NewBadRequestError("invalid annual interest rate"),
		},
		{
			"amortization period less than 5, should throw an error",
			&CalculationInput{
				PropertyPrice:      100000,
				DownPayment:        4300,
				AnnualInterestRate: 4,
				AmortizationPeriod: 1,
				PaymentSchedule:    string(Monthly),
			},
			rest_errors.NewBadRequestError("invalid amortization period"),
		},
		{
			"amortization period is not divided by 5, should throw an error",
			&CalculationInput{
				PropertyPrice:      100000,
				DownPayment:        4300,
				AnnualInterestRate: 4,
				AmortizationPeriod: 7,
				PaymentSchedule:    string(Monthly),
			},
			rest_errors.NewBadRequestError("invalid amortization period"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculationInput.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
