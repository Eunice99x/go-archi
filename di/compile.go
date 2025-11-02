package di

import (
	"fmt"

	"github.com/sarulabs/dingo/v4"
)


type Provider struct {
	dingo.BaseProvider
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Load() error {
	allDeps := [][]dingo.Def{
		getAppRepo(),
		getAppDB(),
		getAppService(),
		getAppHandler(),
	}

	for _, deps := range allDeps {
		if err := p.AddDefSlice(deps); err != nil {
			return fmt.Errorf("add def: %w", err)
		}
	}

	return nil
}