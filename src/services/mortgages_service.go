package services

import (
	"math"
	"strings"

	"github.com/ferralucho/mortgage-calculator-canada/src/domain/mortgages"
	"github.com/ferralucho/mortgage-calculator-canada/src/rest_errors"
)

var (
	MortgagesService mortgagesServiceInterface = &mortgagesService{}
)

type mortgagesService struct{}

type mortgagesServiceInterface interface {
	GetCalculation(mortgages.CalculationInput) (*mortgages.CalculationOutput, rest_errors.RestErr)
}

func (s *mortgagesService) GetCalculation(input mortgages.CalculationInput) (*mortgages.CalculationOutput, rest_errors.RestErr) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	paymentSchedulePeriod := 0

	switch mortgages.PaymentSchedule(strings.ToLower(input.PaymentSchedule)) {
	case mortgages.AcceleratedBiWeekly, mortgages.BiWeekly:
		paymentSchedulePeriod = 2
	case mortgages.Monthly:
		paymentSchedulePeriod = 1
	default:
		return nil, rest_errors.NewBadRequestError("invalid payment schedule")
	}

	principal := input.PropertyPrice - input.DownPayment
	differenceRatio := (principal / input.PropertyPrice) * 100
	paymentScheduleResult := getPaymentSchedule(input.AnnualInterestRate, uint64(paymentSchedulePeriod), principal)

	output := &mortgages.CalculationOutput{
		TotalMortgageTotal: principal,
		MortgagePayment:    paymentScheduleResult,
		DifferenceRatio:    differenceRatio,
	}

	return output, nil
}

func getPaymentSchedule(annualInterestRate float64, paymentSchedulePeriod uint64, principal float64) float64 {
	monthlyInterest := annualInterestRate / 100 / 12
	periods := paymentSchedulePeriod * 12
	return principal * ((monthlyInterest * math.Pow(1+monthlyInterest, float64(periods))) / math.Pow(1+monthlyInterest, float64(periods)))
}
