package overtime

import (
	"context"
	"payslips/internal/presentations"
)

type Overtime interface {
	Create(ctx context.Context, input presentations.Overtime) error
}
