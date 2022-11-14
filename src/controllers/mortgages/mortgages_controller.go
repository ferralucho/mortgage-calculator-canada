package mortgages

import (
	"net/http"

	"github.com/ferralucho/mortgage-calculator-canada/src/domain/mortgages"
	"github.com/ferralucho/mortgage-calculator-canada/src/services"
	"github.com/gin-gonic/gin"
)

func GetCalculation(c *gin.Context) {
	propertyPrice := c.Query("property_price")
	downPayment := c.Query("down_payment")
	annualInterestRate := c.Query("annual_interest_rate")
	amortizationPeriod := c.Query("amortization_period")
	paymentSchedule := c.Query("payment_schedule")

	input := new(mortgages.CalculationInput)
	input.PropertyPrice = c.GetFloat64(propertyPrice)
	input.DownPayment = c.GetFloat64(downPayment)
	input.AnnualInterestRate = c.GetFloat64(annualInterestRate)
	input.AmortizationPeriod = c.GetUint64(amortizationPeriod)
	input.PaymentSchedule = c.Query(paymentSchedule)

	calculation, err := services.MortgagesService.GetCalculation(*input)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, calculation)
}
