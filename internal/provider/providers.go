package provider

import (
	"payslips/bootstrap"
)

type Provider struct {
}

func NewProvider(cfg bootstrap.Providers) Provider {
	return Provider{}
}
