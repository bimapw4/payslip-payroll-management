package business

import (
	"payslips/internal/provider"
	"payslips/internal/repositories"
)

type Business struct {
}

func NewBusiness(repo *repositories.Repository, provider provider.Provider) Business {
	return Business{}
}
