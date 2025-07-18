package repository

import (
	"context"
)

type BuckerRepo interface {
	X(ctx context.Context) error
}
