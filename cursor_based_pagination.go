package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var defaultCursor = "limit,0,sinceId,0"

type CursorBasedPaginationRequest struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

type CursorBasedPaginationResponse struct {
	Data   []Number `json:"data"`
	Cursor string   `json:"cursor"`
	Limit  int      `json:"limit"`
}

func encodeCursor(cursor string) string {
	encodedCursor := base64.StdEncoding.EncodeToString([]byte(cursor))
	return string(encodedCursor)
}

// limit, since_id
func decodeCursor(cursor string) (int, int) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(cursor)))
	n, _ := base64.StdEncoding.Decode(dst, []byte(cursor))
	dst = dst[:n]
	formatCursor := strings.Split(string(dst), ",")
	if len(formatCursor) < 4 {
		return 0, 0
	}

	limit, _ := strconv.Atoi(formatCursor[1])
	sinceID, _ := strconv.Atoi(formatCursor[3])

	return limit, sinceID
}

func CursorBasedPaginationHandler(c echo.Context, db *gorm.DB) error {
	numbers := []Number{}
	request := CursorBasedPaginationRequest{}
	if request.Cursor == "" {
		request.Cursor = encodeCursor(defaultCursor)
	}

	c.Bind(&request)

	limit, sinceID := decodeCursor(request.Cursor)

	if sinceID == 0 {
		limit = request.Limit
	}

	db.
		Where("id > ?", sinceID).
		Order("id").
		Limit(limit).
		Find(&numbers)

	nextCursor := fmt.Sprintf("limit,%v,sinceId,%v", limit, sinceID+limit)

	response := CursorBasedPaginationResponse{
		Data:   numbers,
		Limit:  limit,
		Cursor: encodeCursor(nextCursor),
	}

	return c.JSON(200, response)
}
