package router

import (
	"github.com/engchina/golang-oracle-demo/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func registerRouter(r *gin.Engine) {
	r.GET("/", handler.IndexHandler())
	r.POST("/insertorupdate", handler.InsertOrUpdateHandler)
	r.POST("/updatewithoptimisticlock", handler.UpdateWithOptimisticLockHandler)
	r.POST("/updatewithpessimisticlock", handler.UpdateWithPessimisticLockHandler)
}

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// load html files
	r.Static("../static", "static")
	r.LoadHTMLGlob("templates/*")

	// add routers
	registerRouter(r)

	// start gin
	err := r.Run(":3000")
	if err != nil {
		log.Fatalln("error in start gin server", err)
	}
}
