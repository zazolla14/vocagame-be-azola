package main

import (
	"github.com/quangdangfit/gocommon/logger"
	"github.com/quangdangfit/gocommon/validation"

	orderModel "vocagame-be-azola/internal/order/model"
	productModel "vocagame-be-azola/internal/product/model"
	httpServer "vocagame-be-azola/internal/server/http"
	userModel "vocagame-be-azola/internal/user/model"
	"vocagame-be-azola/pkg/config"
	"vocagame-be-azola/pkg/dbs"
	"vocagame-be-azola/pkg/redis"
)

//	@title				vocagame-be-azola Swagger API
//	@version				1.0
//	@description		Swagger API for vocagame-be-azola.
//	@termsOfService	http://swagger.io/terms/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in									header
//	@name									Authorization

//	@BasePath	/api/v1

func main() {
	cfg := config.LoadConfig()
	logger.Initialize(cfg.Environment)

	db, err := dbs.NewDatabase(cfg.DatabaseURI)
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}

	err = db.AutoMigrate(&userModel.User{}, &productModel.Product{}, orderModel.Order{}, orderModel.OrderLine{})
	if err != nil {
		logger.Fatal("Database migration fail", err)
	}

	validator := validation.New()

	cache := redis.New(redis.Config{
		Address:  cfg.RedisURI,
		Password: cfg.RedisPassword,
		Database: cfg.RedisDB,
	})

	go func() {
		httpSvr := httpServer.NewServer(validator, db, cache)
		if err = httpSvr.Run(); err != nil {
			logger.Fatal(err)
		}
	}()
}
