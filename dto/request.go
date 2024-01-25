package dto

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uzan16/packform-test/helpers"
)

type PaginationRequest struct {
	PageSize      int    `json:"pageSize"`
	PageNumber    int    `json:"pageNumber"`
	SortBy        string `json:"sortBy"`
	SortDirection string `json:"sortDirection"`
	Query         string `json:"q"`
}

func NewPaginationRequest(ctx *gin.Context) *PaginationRequest {
	var pageSize, pageNumber int
	var err error

	sortBy := ctx.DefaultQuery("sortBy", "id")
	sortDirection := ctx.DefaultQuery("sortDirection", "asc")
	query := strings.TrimSpace(strings.ToLower(ctx.DefaultQuery("q", "")))

	if pageSize, err = strconv.Atoi(ctx.DefaultQuery("pageSize", "5")); err != nil {
		pageSize = 5
	}
	if pageNumber, err = strconv.Atoi(ctx.DefaultQuery("pageNumber", "1")); err != nil {
		pageNumber = 1
	}
	if len(sortBy) <= 0 {
		sortBy = "id"
	} else {
		sortBy = helpers.ToSnakeCase(sortBy)
	}
	if len(sortDirection) <= 0 {
		sortDirection = "asc"
	}

	return &PaginationRequest{PageSize: pageSize, PageNumber: pageNumber, SortBy: sortBy, SortDirection: sortDirection, Query: query}
}

func (r *PaginationRequest) Offset() int {
	return (r.PageNumber - 1) * r.PageSize
}

func (r *PaginationRequest) Order() string {
	return fmt.Sprintf("%s %s", r.SortBy, r.SortDirection)
}
