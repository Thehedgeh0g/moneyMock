package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"moneymock/api"
	"moneymock/pkg/app/provider"
	"moneymock/pkg/app/service"
	"moneymock/pkg/infrastructure/transport"
)

func main() {
	var port = flag.Int("port", 8080, "Port for mock HTTP server")
	flag.Parse()

	rateProvider := provider.NewStaticRateProvider()
	currencyService := service.NewCurrencyService(rateProvider)

	petStore := transport.NewPublicWeb(currencyService)

	petStoreStrictHandler := api.NewStrictHandler(petStore, nil)

	r := chi.NewRouter()

	api.HandlerFromMux(petStoreStrictHandler, r)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	log.Fatal(s.ListenAndServe())
}
