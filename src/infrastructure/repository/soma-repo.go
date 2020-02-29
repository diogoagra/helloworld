package repository

import (
	"database/sql"
	"helloworld/src/domain/repository"
)

type somaRepository struct {
	db *sql.DB
}

// NewSomaRepository -
func NewSomaRepository(db *sql.DB) repository.SomaRepository {
	return &somaRepository{
		db: db,
	}
}

func (r *somaRepository) Find() {
	r.db.Exec("")
}
