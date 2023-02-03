package services_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/rodrigoikari/pact-poc-consumer-go/models"
	"github.com/rodrigoikari/pact-poc-consumer-go/services"
	"github.com/stretchr/testify/assert"
)

func Test_Currency_Service_ConvertCurrencyValue(t *testing.T) {
	value := 15.0
	sourceCurrencyCode := "USD"
	destinationCurrencyCode := "BRL"
	exchangeRate := 5.34
	convertedValue := 80.1 // 15.0 * 5.34

	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, req.URL.String(), fmt.Sprintf("/exchange/convert"))

			converted_amount_response, _ := json.Marshal(models.ConvertCurrencyValueResponse{
				Value:                   value,
				ConvertedValue:          convertedValue,
				SourceCurrencyCode:      sourceCurrencyCode,
				DestinationCurrencyCode: destinationCurrencyCode,
				ExchangeRate:            exchangeRate,
			})

			rw.Write([]byte(converted_amount_response))
		}))

	defer server.Close()

	u, _ := url.Parse(server.URL)

	currency_service := services.CurrencyServiceImpl{
		BaseUrl: u,
	}

	ca, err := currency_service.ConvertCurrencyValue(value, sourceCurrencyCode, destinationCurrencyCode)
	assert.NoError(t, err)
	assert.Equal(t, ca, convertedValue)

}
