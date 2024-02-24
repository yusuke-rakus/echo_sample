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

func (cs *CategoryStore) GetCategoryWithSubCategoryList() *model.CategoryWithSubCategoryList {
	var category_with_sub_category_list model.CategoryWithSubCategoryList

	db := db.New()
	db.Table("category AS c").Select("c.category_id, c.category_name, sc.sub_category_id, sc.sub_category_name, hsc.sub_category_id IS NULL as enable").Joins("JOIN sub_category sc ON c.category_id = sc.category_id").Joins("LEFT JOIN hidden_sub_category hsc ON sc.sub_category_id = hsc.sub_category_id").Scan(&category_with_sub_category_list.CategoryWithSubCategoryList)
	return &category_with_sub_category_list
}

func NewCategoryStore(db *gorm.DB) *CategoryStore {
	return &CategoryStore{db: db}
}
