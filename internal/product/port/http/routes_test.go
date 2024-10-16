package http

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/validation"

	dbMocks "vocagame-be-azola/pkg/dbs/mocks"
	redisMocks "vocagame-be-azola/pkg/redis/mocks"
)

func TestRoutes(t *testing.T) {
	mockDB := dbMocks.NewIDatabase(t)
	mockRedis := redisMocks.NewIRedis(t)
	Routes(gin.New().Group("/"), mockDB, validation.New(), mockRedis)
}
