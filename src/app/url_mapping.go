package app

import (
	"github.com/ferralucho/mortgage-calculator-canada/src/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	//router.GET("/mortage/formula", mortgage.GetCalculation)
}
