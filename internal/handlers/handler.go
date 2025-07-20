package handlers

import (
	"payslips/internal/business"
)

type Handlers struct {
}

func NewHandler(business business.Business) Handlers {
	return Handlers{}
}
