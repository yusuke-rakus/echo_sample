package handler

import (
	"fmt"
	"net/http"

	"example.com/m/db"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetCategoryList(c echo.Context) error {
	var (
		category_id   *int
		category_name *string
	)
	result := map[int]string{}
	var result_list []string

	db := db.New()
	rows, err := db.Query("SELECT category_id, category_name FROM category;")
	if err != nil {
		fmt.Println("error")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&category_id, &category_name)
		if err != nil {
			fmt.Println("id_name_error")
		} else {
			result_list = append(result_list, *category_name)
			result[*category_id] = *category_name
		}
	}

	// for k, v := range result {
	// 	fmt.Println(k, v)
	// }
	for i, v := range result_list {
		fmt.Println(i, ": ", v)
	}

	return c.JSON(http.StatusOK, result)
}
