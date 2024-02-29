package handler

import (
	

	"io"
	"net/http"
	"os"


	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

// NewRouter creates a new HTTP router
func NewRouter(RegHandler RegHandler) (*Router, error) {
	// Disable debug mode and write logs to file in production
	env := os.Getenv("APP_ENV")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)

		logFile, _ := os.Create("gin.log")
		gin.DefaultWriter = io.Writer(logFile)
	}

	// CORS
	config := cors.DefaultConfig()

	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"*"}
	config.AllowBrowserExtensions = true
	config.AllowMethods = []string{"*"}

	router := gin.New()
	router.RedirectTrailingSlash = false
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": []string{"Invalid Path"},
			"errorno": []string{"INV1"},
		})

	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"success": false,
			"message": []string{"Method not Allowed"},
			"errorno": []string{"MD01"},
		})

	})

	//r.Use( ValidateContentType( []string{"application/json", "application/xml"}   )   )

	// if env == "production" {
	// 	router.Use(gin.LoggerWithFormatter(customLogger), gin.Recovery(), cors.New(config), ValidateContentType([]string{"application/json"}))
	// }
	// if env == "test" {
	// 	router.Use(gin.LoggerWithFormatter(customLogger), gin.Recovery(), cors.New(config), ValidateContentType([]string{"application/json", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8", "text/css,*/*;q=0.1", "application/json,*/*", "*/*"}))

	// }router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")
	{
		v1.POST("register", RegHandler.CreateReg)
	}
	//}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}

// customLogger is a custom Gin logger
// func customLogger(param gin.LogFormatterParams) string {
// 	return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s [%s]\"\n",
// 		param.TimeStamp.Format(time.RFC1123),
// 		param.ClientIP,
// 		param.Method,
// 		param.Path,
// 		param.Request.Proto,
// 		param.StatusCode,
// 		param.Latency.Round(time.Millisecond),
// 		param.Request.UserAgent(),
// 	)
// }
