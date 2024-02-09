package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetCategoryList(c echo.Context) error {
	// var (
	// 	category_id   *int
	// 	category_name *string
	// )

	// db := db.New()
	// rows, err := db.Query("SELECT category_id, category_name FROM category;")
	// if err != nil {
	// 	fmt.Println("error")
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	if err := rows.Scan(&category_id); err != nil {
	// 		fmt.Println("error")
	// 	}
	// 	fmt.Println(*category_id, *category_name)
	// }

	data2 := "data2"
	data3 := "data3"

	return c.JSON(http.StatusOK, [3]string{"data1", data2, data3})
}
