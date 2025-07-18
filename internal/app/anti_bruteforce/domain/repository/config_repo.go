package repository

import (
	"context"

	"github.com/gkarman/anti_bruteforce/internal/app/anti_bruteforce/domain/valueobject"
)

type ConfigRepo interface {
	AddCidrInBlacklist(ctx context.Context, cidr valueobject.CIDR) error
}
