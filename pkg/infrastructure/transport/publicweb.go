package transport

import (
	"context"
	"moneymock/pkg/app/service"

	publicapi "moneymock/api"
)

func NewPublicWeb(
	currencyService service.CurrencyService,
) publicapi.StrictServerInterface {
	return &publicWeb{
		currencyService: currencyService,
	}
}

type publicWeb struct {
	currencyService service.CurrencyService
}

func (p publicWeb) GetApiRate(_ context.Context, request publicapi.GetApiRateRequestObject) (publicapi.GetApiRateResponseObject, error) {
	target := request.Params.Target
	base := "USD"
	rate, err := p.currencyService.GetRate(string(target))
	return publicapi.GetApiRate200JSONResponse{
		Base:   base,
		Rate:   rate,
		Target: string(request.Params.Target),
	}, err
}
