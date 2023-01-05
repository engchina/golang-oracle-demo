package router

import (
	"github.com/engchina/golang-oracle-demo/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

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
