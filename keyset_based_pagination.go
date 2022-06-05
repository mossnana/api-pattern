package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type KeysetBasedPaginationRequest struct {
	Limit   int `json:"limit"`
	SinceID int `json:"sinceId"` // since id
}

type KeysetBasedPaginationResponse struct {
	Data    []Number `json:"data"`
	Limit   int      `json:"limit"`
	SinceID int      `json:"sinceId"` // since id
}

func KeysetBasedPaginationHandler(c echo.Context, db *gorm.DB) error {
	numbers := []Number{}
	request := KeysetBasedPaginationRequest{}
	c.Bind(&request)

	db.
		Where("id > ?", request.SinceID).
		Order("id").
		Limit(request.Limit).
		Find(&numbers)

	response := KeysetBasedPaginationResponse{
		Data:    numbers,
		Limit:   request.Limit,
		SinceID: request.SinceID,
	}

	return c.JSON(200, response)
}
