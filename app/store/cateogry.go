package store

import (
	"database/sql"
)

type CategoryStore struct {
	db *sql.DB
}

// GetCategoryList implements category.Store.
func (*CategoryStore) GetCategoryList() error {
	panic("unimplemented")
}

func NewCategoryStore(db *sql.DB) *CategoryStore {
	return &CategoryStore{db: db}
}
