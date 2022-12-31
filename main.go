package main

import (
	"github.com/engchina/golang-oracle-demo/models"
	"github.com/engchina/golang-oracle-demo/router"
	"github.com/engchina/golang-oracle-demo/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
	"github.com/sirupsen/logrus"
	"net/http"
)

func initMyUserDBEngine() {
	MyUserDBEngine.Engine = utils.DBEngine
	// set context with timeout
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	//defer cancel()
	//MyUserDBEngine.Engine.SetDefaultContext(ctx)
	// create table
	err := MyUserDBEngine.Sync(new(models.MyUser))
	if err != nil {
		panic(err)
	}
}

var MyUserDBEngine models.MyUserEngine

func main() {

	logrus.SetLevel(logrus.InfoLevel)
	initMyUserDBEngine()

	//// insert
	//var newMyUser models.MyUser
	//newMyUser.UserId = "200"
	//newMyUser.Name = "first"
	//resp, err := MyUserDBEngine.Transaction(Insert, &newMyUser)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// update
	//newMyUser.Name = "fourth"
	//resp, err = MyUserDBEngine.Transaction(AddTryCountWithLock, &newMyUser)
	//if err != nil {
	//	panic("err: " + err.Error())
	//}
	//logrus.Infof("resp: %#v\n", resp)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Static("../static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Pong"})
	})
	r.GET("/", router.IndexHandler())
	r.POST("/insertorupdate", router.RegisterHandler)
	r.POST("/updatewithoptimisticlock", router.UpdateWithOptimisticLockHandler())
	r.POST("/updatewithpessimisticlock", router.UpdateWithPessimisticLockHandler())
	err := r.Run(":3000")
	if err != nil {
		return
	}
}

/*
select * from MY_USER where USER_ID = '9999';

update MY_USER set NAME='another session', NUM_OF_TRIED=NUM_OF_TRIED+1 where USER_ID = '9999';
commit;

select * from MY_USER where USER_ID = '9999';
*/
