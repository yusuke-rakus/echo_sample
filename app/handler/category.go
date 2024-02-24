package handler

import (
	"net/http"
	"fmt"

	"example.com/m/store"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetCategoryList(c echo.Context) error {
	cs := store.CategoryStore{}
	result := cs.GetCategoryList()
	result_list := getCategoryListResponse(result)

	return c.JSON(http.StatusOK, *result_list)
}

func (h *Handler) GetCategoryWithSubCategoryList(c echo.Context) error {
	cs := store.CategoryStore{}
	result := cs.GetCategoryWithSubCategoryList()
	for _, c := range result.category_with_sub_category_list{
		fmt.Println(*c)
	}
	// result_list := getCategoryListResponse(result)

	// return c.JSON(http.StatusOK, *result_list)
	return c.JSON(http.StatusOK, [3]string{"ok", "ok", "ng"})
}
