package app

import (
	"github.com/gkarman/anti_bruteforce/internal/app/anti_bruteforce/domain/repository"
)

type AntiBruteForceApp struct {
	configRepo repository.ConfigRepo
	bucketRepo repository.BuckerRepo
}

func NewAntiBruteForceApp(configRepo repository.ConfigRepo, bucketRepo repository.BuckerRepo) *AntiBruteForceApp {
	return &AntiBruteForceApp{
		configRepo: configRepo,
		bucketRepo: bucketRepo,
	}
}
