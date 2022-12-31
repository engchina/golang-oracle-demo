package router

import (
	docs "github.com/engchina/golang-oracle-demo/docs"
	"github.com/engchina/golang-oracle-demo/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// add swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:3000/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/ping", handler.PingHandler)
		}
	}

	// load html files
	r.Static("../static", "static")
	r.LoadHTMLGlob("templates/*")

	// add routers
	r.GET("/", handler.IndexHandler())
	r.POST("/insertorupdate", handler.InsertOrUpdateHandler)
	r.POST("/updatewithoptimisticlock", handler.UpdateWithOptimisticLockHandler)
	r.POST("/updatewithpessimisticlock", handler.UpdateWithPessimisticLockHandler)
	return r
}
