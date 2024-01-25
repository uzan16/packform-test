package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uzan16/packform-test/dto"
	"github.com/uzan16/packform-test/models"
	"gorm.io/gorm"
)

type OrdersController struct {
	db     *gorm.DB
	router *gin.Engine
}

func NewOrdersController(db *gorm.DB, router *gin.Engine) *OrdersController {
	return &OrdersController{db: db, router: router}
}

func (r *OrdersController) Initialize() {
	orderRoute := r.router.Group("/api/orders")
	orderRoute.GET("/", r.getOrders)
}

func (r *OrdersController) getOrders(ctx *gin.Context) {
	request := dto.NewPaginationRequest(ctx)

	startRange := ctx.DefaultQuery("startDate", "")
	endRange := ctx.DefaultQuery("endDate", "")

	if t, err := time.Parse(time.RFC3339Nano, startRange); err != nil {
		startRange = ""
	} else {
		startRange = strconv.FormatInt(t.Unix(), 10)
	}
	if t, err := time.Parse(time.RFC3339Nano, endRange); err != nil {
		endRange = ""
	} else {
		endRange = strconv.FormatInt(t.Unix(), 10)
	}

	var data []dto.OrderResponse
	var totalRows int64

	query := models.BuildGetOrderQuery(r.db)

	if len(request.Query) > 0 {
		query.Where(
			"LOWER(orders.order_name) LIKE ? OR LOWER(order_items.product) LIKE ?",
			fmt.Sprintf("%%%s%%", request.Query),
			fmt.Sprintf("%%%s%%", request.Query),
		)
	}

	if len(startRange) > 0 && len(endRange) > 0 {
		query.Where(
			"DATE_PART('EPOCH', orders.created_at) BETWEEN ? AND ?",
			startRange,
			endRange,
		)
	}

	query.Count(&totalRows)

	if totalRows <= 0 {
		ctx.JSON(http.StatusOK, dto.PaginatedResponse{
			Response: dto.Response{
				Success: true,
				Message: "No Data Found",
				Data:    []interface{}{},
			},
			TotalRows: totalRows,
		})
		return
	}

	query.Limit(request.PageSize).Offset(request.Offset()).Order(request.Order()).Find(&data)

	ctx.JSON(http.StatusOK, dto.PaginatedResponse{
		Response: dto.Response{
			Success: true,
			Message: "Sucess retreiving data",
			Data:    data,
		},
		TotalRows: totalRows,
	})
}
