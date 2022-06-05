package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PageBasedPaginationRequest struct {
	Page  int `json:"page"`  // current's page
	Limit int `json:"limit"` // limit is page's size
}

type PageBasedPaginationResponse struct {
	Data  []Number `json:"data"`
	Page  int      `json:"page"`
	Limit int      `json:"limit"`
}

func (c *PageBasedPaginationRequest) Offset() int {
	page := c.Page
	if page > 0 {
		page -= 1
	}

	return c.Limit * page
}

func PageBasedPaginationHandler(c echo.Context, db *gorm.DB) error {
	numbers := []Number{}
	request := PageBasedPaginationRequest{}
	c.Bind(&request)

	db.Limit(request.Limit).Offset(request.Offset()).Find(&numbers)

	response := PageBasedPaginationResponse{
		Data:  numbers,
		Page:  request.Page,
		Limit: request.Limit,
	}

	return c.JSON(200, response)
}
