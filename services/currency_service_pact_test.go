package services_test

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/rodrigoikari/pact-poc-consumer-go/services"
)

func TestMain(m *testing.M) {
	var exitCode int

	// Setup Pact and related test stuff
	setup()

	// Run all the tests
	exitCode = m.Run()

	// Shutdown the Mock Service and Write pact files to disk
	if err := pact.WritePact(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pact.Teardown()
	os.Exit(exitCode)
}

// Common test data
var pact dsl.Pact

// Aliases
var term = dsl.Term

type request = dsl.Request

var u *url.URL
var currency_service *services.CurrencyServiceImpl

func setup() {
	pact = createPact()

	// Proactively start service to get access to the port
	pact.Setup(true)

	u, _ = url.Parse(fmt.Sprintf("http://localhost:%d", pact.Server.Port))

	currency_service = &services.CurrencyServiceImpl{
		BaseUrl: u,
	}

}

func createPact() dsl.Pact {
	return dsl.Pact{
		Consumer: os.Getenv("CONSUMER_NAME"),
		Provider: os.Getenv("PROVIDER_NAME"),
		LogDir:   os.Getenv("LOG_DIR"),
		PactDir:  os.Getenv("PACT_DIR"),
		LogLevel: "INFO",
	}
}
