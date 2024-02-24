package model

type Category struct {
	CategoryId   int    `json:category_id`
	CategoryName string `json:category_name`
}

type CategoryList struct {
	CategoryList []Category `json:"category_list"`
}

type CategoryWithSubCategoryList struct {
	CategoryWithSubCategoryList []CategoryWithSubCategory `json:"category_with_sub_category_list"`
}

type CategoryWithSubCategory struct {
	CategoryId      int           `json:category_id`
	CategoryName    string        `json:category_name`
	SubCategoryList []SubCategory `json:"sub_category_list"`
}

type SubCategory struct {
	SubCategoryId   int    `json:sub_category_id`
	SubCategoryName string `json:sub_category_name`
	Enable          bool   `json:enable`
}
