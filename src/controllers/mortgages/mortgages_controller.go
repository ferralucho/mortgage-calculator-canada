package mortgages

import (
	"net/http"
	"strconv"

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
	var errParsing error

	input.PropertyPrice, errParsing = strconv.ParseFloat(propertyPrice, 64)

	if errParsing != nil {
		c.JSON(400, errParsing)
		return
	}
	input.DownPayment, errParsing = strconv.ParseFloat(downPayment, 64)

	if errParsing != nil {
		c.JSON(400, errParsing)
		return
	}
	input.AnnualInterestRate, errParsing = strconv.ParseFloat(annualInterestRate, 64)

	if errParsing != nil {
		c.JSON(400, errParsing)
		return
	}
	input.AmortizationPeriod, errParsing = strconv.ParseUint(amortizationPeriod, 10, 64)
	input.PaymentSchedule = paymentSchedule

	if errParsing != nil {
		c.JSON(400, errParsing)
		return
	}

	calculation, err := services.MortgagesService.GetCalculation(*input)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, calculation)
}
