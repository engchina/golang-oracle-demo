package handler

import (
	"github.com/engchina/golang-oracle-demo/facade"
	"github.com/engchina/golang-oracle-demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} pond
// @Router /example/ping [get]
func PingHandler(g *gin.Context) {
	g.JSON(http.StatusOK, "pond")
}

func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		myUserList, err := facade.GetMyUserList()
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "register.html", gin.H{"showAffected": false, "affected": 0, "msgColor": "bg-info text-dark", "myUserList": myUserList})
	}
}

func InsertOrUpdateHandler(c *gin.Context) {
	var myUser models.MyUser
	err := c.ShouldBind(&myUser)
	if err != nil {
		return
	}
	affected, err := facade.InsertOrUpdate(&myUser)
	if err != nil {
		panic(err)
	}
	var msgColor string
	if affected == int64(1) {
		msgColor = "bg-success text-white"
	} else {
		msgColor = "bg-danger text-white"
	}

	myUserList, err := facade.GetMyUserList()
	if err != nil {
		panic(err)
	}
	c.HTML(http.StatusOK, "register.html", gin.H{"myUser": myUser, "showAffected": true, "affected": affected, "msgColor": msgColor, "myUserList": myUserList})
}

func UpdateWithOptimisticLockHandler(c *gin.Context) {
	var myUser models.MyUser
	err := c.ShouldBind(&myUser)
	if err != nil {
		return
	}
	affected, err := facade.UpdateWithOptimisticLock(&myUser)
	if err != nil {
		panic(err)
	}
	var msgColor string
	if affected == int64(1) {
		msgColor = "bg-success text-white"
	} else {
		msgColor = "bg-danger text-white"
	}

	myUserList, err := facade.GetMyUserList()
	if err != nil {
		panic(err)
	}
	c.HTML(http.StatusOK, "register.html", gin.H{"myUser": myUser, "showAffected": true, "affected": affected, "msgColor": msgColor, "myUserList": myUserList})

}

func UpdateWithPessimisticLockHandler(c *gin.Context) {
	var myUser models.MyUser
	err := c.ShouldBind(&myUser)
	if err != nil {
		return
	}
	affected, err := facade.UpdateWithPessimisticLock(&myUser)
	if err != nil {
		panic(err)
	}
	var msgColor string
	if affected == int64(1) {
		msgColor = "bg-success text-white"
	} else {
		msgColor = "bg-danger text-white"
	}

	myUserList, err := facade.GetMyUserList()
	if err != nil {
		panic(err)
	}
	c.HTML(http.StatusOK, "register.html", gin.H{"myUser": myUser, "showAffected": true, "affected": affected, "msgColor": msgColor, "myUserList": myUserList})

}
