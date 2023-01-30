package services

type CurrencyService interface {
	ConvertCurrencyValue(amount float64, sourceCurrencyCode string, destinationCurrencyCode string) (float64, error)
}
