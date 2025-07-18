package bucketrepo

import (
	"context"

	"github.com/gkarman/anti_bruteforce/internal/config"
)

type RedisBucketRepo struct {
}

func NewRedisBucketRepo(_ config.MemoryRepo) (*RedisBucketRepo, error) {
	return &RedisBucketRepo{}, nil
}

func (s *RedisBucketRepo) X(ctx context.Context) error {
	return nil
}
