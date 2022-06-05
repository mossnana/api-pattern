package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := GetDB()

	e.GET("/migrate", func(c echo.Context) error {
		Migrate(db)
		return c.NoContent(200)
	})

	e.GET("/insert", func(c echo.Context) error {
		numbers := []Number{}
		for i := range make([]int, 100) {
			numbers = append(numbers, Number{
				ID: i + 101,
			})
		}

		db.Create(&numbers)

		return c.NoContent(200)
	})

	e.GET("/page-based-pagination", func(c echo.Context) error {
		return PageBasedPaginationHandler(c, db)
	})

	e.GET("/keyset-based-pagination", func(c echo.Context) error {
		return KeysetBasedPaginationHandler(c, db)
	})

	e.GET("/cursor-based-pagination", func(c echo.Context) error {
		return CursorBasedPaginationHandler(c, db)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
