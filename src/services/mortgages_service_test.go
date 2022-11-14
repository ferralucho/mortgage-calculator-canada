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
			"",
			mortgages.CalculationInput{
				PropertyPrice:      900000,
				DownPayment:        65000,
				AnnualInterestRate: 4,
				AmortizationPeriod: 20,
				PaymentSchedule:    string(mortgages.Monthly),
			},
			&mortgages.CalculationOutput{
				TotalMortgageTotal: 835000,
				MortgagePayment:    5059.94,
				DifferenceRatio:    92.78,
			},
			nil,
		},
		{
			"",
			mortgages.CalculationInput{
				PropertyPrice:      900000,
				DownPayment:        65000,
				AnnualInterestRate: 4.94,
				AmortizationPeriod: 20,
				PaymentSchedule:    string(mortgages.Monthly),
			},
			&mortgages.CalculationOutput{
				TotalMortgageTotal: 835000,
				MortgagePayment:    5482.99,
				DifferenceRatio:    92.78,
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
