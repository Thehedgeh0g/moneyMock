package provider

import "fmt"

type RateProvider interface {
	GetRate(target string) (float64, error)
}

type StaticRateProvider struct {
	rates map[string]float64
}

func NewStaticRateProvider() *StaticRateProvider {
	return &StaticRateProvider{
		rates: map[string]float64{
			"EUR": 0.92,
			"JPY": 153.25,
			"RUB": 91.35,
		},
	}
}

func (rp *StaticRateProvider) GetRate(target string) (float64, error) {
	rate, ok := rp.rates[target]
	if !ok {
		return 0, fmt.Errorf("rate for %s not found", target)
	}
	return rate, nil
}
