package service

import (
	appprovider "moneymock/pkg/app/provider"
)

type CurrencyService struct {
	provider appprovider.RateProvider
}

func NewCurrencyService(provider appprovider.RateProvider) CurrencyService {
	return CurrencyService{
		provider: provider,
	}
}

func (cs *CurrencyService) GetRate(target string) (int, error) {
	rate, err := cs.provider.GetRate(target)
	if err != nil {
		return 0, err
	}
	return int(rate + 0.5), nil // округление
}
