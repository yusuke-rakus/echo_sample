package handler

import (
	"net/http"

	"example.com/m/store"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetCategoryList(c echo.Context) error {
	cs := store.CategoryStore{}
	result := cs.GetCategoryList()
	result_list := getCategoryListResponse(result)

	return c.JSON(http.StatusOK, *result_list)
}
