package routes

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"github.com/your-org/your-app/config"
	"github.com/your-org/your-app/contract"
	"github.com/your-org/your-app/handler"
)

func NewRouter(s *contract.Service) *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = false

	cfg := config.Load()

	var limitter int64
	if cfg.IsProduction == false {
		limitter = 1000
	} else {
		limitter = 100
	}

	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  limitter,
	}

	store := memory.NewStore()
	instance := limiter.New(store, rate)
	rateLimitter := mgin.NewMiddleware(instance)
	r.Use(rateLimitter)

	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowAllOrigins = true
	r.Use(cors.New(defaultConfig))

	healthController := &handler.HealthHandler{}
	healthController.InitService(s.HealthService)

	api := r.Group("/")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "Server is Online!",
			})
		})
		api.GET("/health", healthController.GetHealthStatus)
	}

	return r
}
