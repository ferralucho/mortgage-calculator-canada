# mortgage-calculator-canada

Mortage calculator for Canada

CMHC insurance must be considered. Guidelines for the calculation, and restrictions, can be found here:
https://www.ratehub.ca/cmhc-insurance-british-columbia.

https://www.ratehub.ca/mortgage-payment-calculator

## Input

● property price

● down payment

● annual interest rate

● amortization period (5 year increments between 5 and 30 years)

● payment schedule (accelerated bi-weekly, bi-weekly, monthly)

```
curl --location --request GET 'localhost:8082/mortgage/formula?property_price=900000&down_payment=65700&amortization_period=25&annual_interest_rate=5&payment_schedule=monthly'
```

## Expected Output

● payment per payment schedule
● an error if the inputs are not valid. This includes cases where the down payment is not large enough
