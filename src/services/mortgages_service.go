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

const ChmcRate = 3.10

type mortgagesService struct{}

type mortgagesServiceInterface interface {
	GetCalculation(mortgages.CalculationInput) (*mortgages.CalculationOutput, rest_errors.RestErr)
}

func (s *mortgagesService) GetCalculation(input mortgages.CalculationInput) (*mortgages.CalculationOutput, rest_errors.RestErr) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	var paymentSchedulePeriod uint64

	switch mortgages.PaymentSchedule(strings.ToLower(input.PaymentSchedule)) {
	case mortgages.AcceleratedBiWeekly, mortgages.BiWeekly:
		paymentSchedulePeriod = 84
	case mortgages.Monthly:
		paymentSchedulePeriod = 12
	default:
		return nil, rest_errors.NewBadRequestError("invalid payment schedule")
	}

	principal := input.PropertyPrice - input.DownPayment
	differenceRatio := (principal / input.PropertyPrice) * 100
	paymentScheduleResult := getPaymentSchedule(input.AnnualInterestRate, uint64(input.AmortizationPeriod), paymentSchedulePeriod, principal)
	chmcInsurance := 0.0

	if isEligibleForChmcInsurance(input) {
		chmcInsurance = calculateChmcInsurance(principal)
	}

	output := &mortgages.CalculationOutput{
		MortgageTotal:           math.Round((principal+chmcInsurance)*100) / 100,
		MortgageBeforeChmc:      math.Round(principal*100) / 100,
		MortgagePaymentSchedule: math.Round(paymentScheduleResult*100) / 100,
		DifferenceRatio:         math.Round(differenceRatio*100) / 100,
		ChmcInsuranceTotal:      chmcInsurance,
	}

	return output, nil
}

func getPaymentSchedule(annualInterestRate float64, amortizationPeriod uint64, paymentSchedulePeriod uint64, principal float64) float64 {
	monthlyInterest := annualInterestRate / 100 / 12
	periods := amortizationPeriod * paymentSchedulePeriod
	return principal * ((monthlyInterest * math.Pow(1+monthlyInterest, float64(periods))) / (math.Pow(1+monthlyInterest, float64(periods)) - 1))
}

func isEligibleForChmcInsurance(input mortgages.CalculationInput) bool {
	differenceRatio := (input.DownPayment * 100) / input.PropertyPrice

	return !(differenceRatio >= 20 || input.PropertyPrice > 1000000 || input.AmortizationPeriod > 25)
}

func calculateChmcInsurance(principal float64) float64 {
	return (principal * ChmcRate) / 100
}
