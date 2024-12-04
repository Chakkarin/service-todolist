package server

import (
	"net/http"
	"sync"

	"github.com/Chakkarin/service-todolist/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type ginServer struct {
	app  *gin.Engine
	db   *gorm.DB
	conf *config.Config
}

var (
	server *ginServer
	once   sync.Once
)

func NewServer(conf *config.Config, db *gorm.DB) *ginServer {

	gin.SetMode(gin.DebugMode)

	ginApp := gin.Default()
	once.Do(func() {
		server = &ginServer{
			app:  ginApp,
			db:   db,
			conf: conf,
		}
	})

	return server
}

func (s *ginServer) Start() {

	s.app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.app.GET("/v1/health", HealthCheckHandler)

	s.initTodosRouter()

	s.app.Run(":4000")
}

// HealthCheckHandler godoc
// @summary Health Check
// @description Health checking for the service
// @id HealthCheckHandler
// @produce plain
// @response 200 {string} string "OK"
// @router /v1/health [get]
func HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
