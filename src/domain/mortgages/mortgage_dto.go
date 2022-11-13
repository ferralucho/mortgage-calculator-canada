package mortgages

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
