package service

import (
	"errors"

	"github.com/shopspring/decimal"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var exchangeRates = map[string]map[string]float64{
	"TWD": {
		"TWD": 1,
		"JPY": 3.669,
		"USD": 0.03281,
	},
	"JPY": {
		"TWD": 0.26956,
		"JPY": 1,
		"USD": 0.00885,
	},
	"USD": {
		"TWD": 30.444,
		"JPY": 111.801,
		"USD": 1,
	},
}

func Convert(source, target string, convert_amount float64) (string, error) {
	sourceRates, ok := exchangeRates[source]
	if !ok {
		return "", errors.New("source not found")
	}
	targetRate, ok := sourceRates[target]
	if !ok {
		return "", errors.New("target not found")
	}

	rate := decimal.NewFromFloat(targetRate)

	amount, _ := decimal.NewFromFloat(convert_amount).Mul(rate).Round(2).Float64()

	printer := message.NewPrinter(language.English)
	amountStr := printer.Sprintf("$%.2f", amount)

	return amountStr, nil
}
