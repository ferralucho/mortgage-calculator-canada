package services

import (
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

func (s *mortgagesService) GetCalculation(mortgage mortgages.CalculationInput) (*mortgages.CalculationOutput, rest_errors.RestErr) {
	if err := mortgage.Validate(); err != nil {
		return nil, err
	}

	output := new(mortgages.CalculationOutput)

	return output, nil
}
