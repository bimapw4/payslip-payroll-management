package business

import (
	"payslips/internal/business/auth"
	"payslips/internal/repositories"
)

type Business struct {
	Auth auth.Contract
}

func NewBusiness(repo *repositories.Repository) Business {
	return Business{
		Auth: auth.NewBusiness(repo),
	}
}
