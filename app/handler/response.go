package handler

import (
	"example.com/m/model"
)

type categoryResponse struct {
	Category_id   int    `json:"categoryId"`
	Category_name string `json:"categoryName"`
}

type categoryListResponse struct {
	CategoryList []categoryResponse `json:"categoryList"`
}

func getCategoryListResponse(data *model.CategoryList) *categoryListResponse {
	cl := &categoryListResponse{}

	for _, category := range data.CategoryList {
		cr := &categoryResponse{Category_id: category.CategoryId, Category_name: category.CategoryName}
		cl.CategoryList = append(cl.CategoryList, *cr)
	}

	return cl
}

func getCategoryWithSubCategoryListResponse(data *model.CategoryList) *categoryListResponse {
	cl := &categoryListResponse{}

	for _, category := range data.CategoryList {
		cr := &categoryResponse{Category_id: category.CategoryId, Category_name: category.CategoryName}
		cl.CategoryList = append(cl.CategoryList, *cr)
	}

	return cl
}
