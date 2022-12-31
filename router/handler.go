package router

import (
	"github.com/engchina/golang-oracle-demo/models"
	"github.com/engchina/golang-oracle-demo/service"
	"github.com/engchina/golang-oracle-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		myUserList, err := service.GetMyUserList(utils.MyUserDBEngine.NewSession())
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "register.html", gin.H{"affected": 0, "msgColor": "bg-info text-dark", "myUserList": myUserList})
	}
}

func RegisterHandler(c *gin.Context) {
	var myuser models.MyUser
	err := c.ShouldBind(&myuser)
	if err != nil {
		return
	}
	affected, err := utils.MyUserDBEngine.Transaction(service.InsertOrUpdate, &myuser)
	if err != nil {
		panic(err)
	}
	var msgColor string
	if affected == int64(1) {
		msgColor = "bg-success text-white"
	} else {
		msgColor = "bg-danger text-white"
	}

	myUserList, err := service.GetMyUserList(utils.MyUserDBEngine.NewSession())
	if err != nil {
		panic(err)
	}
	c.HTML(http.StatusOK, "register.html", gin.H{"myuser": myuser, "affected": affected, "msgColor": msgColor, "myUserList": myUserList})
}

func UpdateWithOptimisticLockHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var myuser models.MyUser
		err := c.ShouldBind(&myuser)
		if err != nil {
			return
		}
		affected, err := utils.MyUserDBEngine.Transaction(service.UpdateWithOptimisticLock, &myuser)
		if err != nil {
			panic(err)
		}
		var msgColor string
		if affected == int64(1) {
			msgColor = "bg-success text-white"
		} else {
			msgColor = "bg-danger text-white"
		}

		myUserList, err := service.GetMyUserList(utils.MyUserDBEngine.NewSession())
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "register.html", gin.H{"myuser": myuser, "affected": affected, "msgColor": msgColor, "myUserList": myUserList})
	}
}

func UpdateWithPessimisticLockHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var myuser models.MyUser
		err := c.ShouldBind(&myuser)
		if err != nil {
			return
		}
		affected, err := utils.MyUserDBEngine.Transaction(service.UpdateWithPessimisticLock, &myuser)
		if err != nil {
			panic(err)
		}
		var msgColor string
		if affected == int64(1) {
			msgColor = "bg-success text-white"
		} else {
			msgColor = "bg-danger text-white"
		}

		myUserList, err := service.GetMyUserList(utils.MyUserDBEngine.NewSession())
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "register.html", gin.H{"myuser": myuser, "affected": affected, "msgColor": msgColor, "myUserList": myUserList})
	}
}
