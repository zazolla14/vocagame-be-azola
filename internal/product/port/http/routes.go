package http

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/validation"

	"vocagame-be-azola/internal/product/repository"
	"vocagame-be-azola/internal/product/service"
	"vocagame-be-azola/pkg/dbs"
	"vocagame-be-azola/pkg/middleware"
	"vocagame-be-azola/pkg/redis"
)

func Routes(r *gin.RouterGroup, db dbs.IDatabase, validator validation.Validation, cache redis.IRedis) {
	productRepo := repository.NewProductRepository(db)
	productSvc := service.NewProductService(validator, productRepo)
	productHandler := NewProductHandler(cache, productSvc)

	authMiddleware := middleware.JWTAuth()

	productRoute := r.Group("/products")
	{
		productRoute.GET("", productHandler.ListProducts)
		productRoute.POST("", authMiddleware, productHandler.CreateProduct)
		productRoute.PUT("/:id", authMiddleware, productHandler.UpdateProduct)
		productRoute.GET("/:id", productHandler.GetProductByID)
	}
}
