package services

import (
	"reflect"
	"testing"

	"github.com/ferralucho/mortgage-calculator-canada/src/domain/mortgages"
	"github.com/ferralucho/mortgage-calculator-canada/src/rest_errors"
)

var (
	testService mortgagesServiceInterface = &mortgagesService{}
)

func Test_mortgagesService_GetCalculation(t *testing.T) {
	tests := []struct {
		name      string
		args      mortgages.CalculationInput
		want      *mortgages.CalculationOutput
		wantError rest_errors.RestErr
	}{
		{
			"down payment not enough, should throw an error",
			mortgages.CalculationInput{
				PropertyPrice:      900000,
				DownPayment:        65000,
				AnnualInterestRate: 4,
				AmortizationPeriod: 20,
				PaymentSchedule:    string(mortgages.Monthly),
			},
			nil,
			rest_errors.NewBadRequestError("invalid down payment"),
		},
		{
			"correct inputs, should not throw an error",
			mortgages.CalculationInput{
				PropertyPrice:      900000,
				DownPayment:        94000,
				AnnualInterestRate: 4.94,
				AmortizationPeriod: 20,
				PaymentSchedule:    string(mortgages.Monthly),
			},
			&mortgages.CalculationOutput{
				MortgageTotal:           806000,
				MortgagePaymentSchedule: 5292.56,
				DifferenceRatio:         89.56,
				MortgageBeforeChmc:      806000,
				ChmcInsuranceTotal:      0,
			},
			nil,
		},
		{
			"",
			mortgages.CalculationInput{
				PropertyPrice:      900000,
				DownPayment:        180000,
				AnnualInterestRate: 5.19,
				AmortizationPeriod: 25,
				PaymentSchedule:    string(mortgages.Monthly),
			},
			&mortgages.CalculationOutput{
				MortgageTotal:           720000,
				MortgagePaymentSchedule: 4289.13,
				DifferenceRatio:         80,
				MortgageBeforeChmc:      720000,
				ChmcInsuranceTotal:      0,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := testService.GetCalculation(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mortgagesService.GetCalculation() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantError) {
				t.Errorf("mortgagesService.GetCalculation() got1 = %v, want %v", got1, tt.wantError)
			}

		})
	}
}
