package store

import (
	"example.com/m/db"
	"example.com/m/model"
	"gorm.io/gorm"
)

type CategoryStore struct {
	db *gorm.DB
}

func (cs *CategoryStore) GetCategoryList() *model.CategoryList {
	var category_list model.CategoryList

	db := db.New()
	db.Table("category").Select("category_id, category_name").Find(&category_list.CategoryList)
	return &category_list
}

func NewCategoryStore(db *gorm.DB) *CategoryStore {
	return &CategoryStore{db: db}
}
