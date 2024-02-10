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
	var result []string

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
			result = append(result, *category_name)
			fmt.Println(*category_id, *category_name)
		}
		// if err := rows.Scan(&category_id); err != nil {
		// 	fmt.Println("id_error")
		// }
		// if err := rows.Scan(&category_name); err != nil {
		// 	fmt.Println("name_error")
		// }
		// fmt.Println(*category_id, *category_name)
		// fmt.Println(*category_id)
	}

	return c.JSON(http.StatusOK, result)
}
