package data

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrRecordNotFound = fmt.Errorf("movie not found")

type Models struct {
	Movies MovieModel
}

func NewModels(db *pgxpool.Pool) *Models {
	return &Models{
		Movies: MovieModel{DB: db},
	}
}
