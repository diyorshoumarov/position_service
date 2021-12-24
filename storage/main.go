package storage

import (
	"bitbucket.org/Udevs/position_service/storage/postgres"
	"bitbucket.org/Udevs/position_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Profession() repo.ProfessionRepoI
}

type storagePG struct {
	profession repo.ProfessionRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return &storagePG{
		profession: postgres.NewProfessionRepo(db),
	}
}

func (s *storagePG) Profession() repo.ProfessionRepoI {
	return s.profession
}
