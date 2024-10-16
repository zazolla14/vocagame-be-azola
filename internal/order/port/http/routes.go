package http

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/validation"

	"vocagame-be-azola/internal/order/repository"
	"vocagame-be-azola/internal/order/service"
	"vocagame-be-azola/pkg/dbs"
	"vocagame-be-azola/pkg/middleware"
)

func Routes(r *gin.RouterGroup, db dbs.IDatabase, validator validation.Validation) {
	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	productSvc := service.NewOrderService(validator, orderRepo, productRepo)
	orderHandler := NewOrderHandler(productSvc)

	authMiddleware := middleware.JWTAuth()

	orderRoute := r.Group("/orders", authMiddleware)
	{
		orderRoute.POST("", orderHandler.PlaceOrder)
		orderRoute.GET("/:id", orderHandler.GetOrderByID)
		orderRoute.GET("", orderHandler.GetOrders)
		orderRoute.PUT("/:id/cancel", orderHandler.CancelOrder)
	}
}
