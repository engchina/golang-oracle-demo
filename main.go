package main

import (
	"github.com/engchina/golang-oracle-demo/config"
	"github.com/engchina/golang-oracle-demo/model"
	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"xorm.io/xorm"
)

func GetAllMyUser(session *xorm.Session) (interface{}, error) {
	var allData []*model.MyUser
	allData, err := model.GetAllMyUser(session)
	if err != nil {
		return nil, err
	}
	return allData, nil
}

//func Insert(session *xorm.Session, myUser *model.MyUser) (interface{}, error) {
//	affected, err := myUser.InsertMyUserInTxn(session)
//	if err != nil {
//		return -1, err
//	}
//	logrus.Infof("affected is: %#v\n", affected)
//	return affected, nil
//}

func InsertOrUpdate(session *xorm.Session, myUser *model.MyUser) (interface{}, error) {
	var myUserModel *model.MyUser
	myUserModel, has, err := model.GetMyUserInTxn(session, myUser.UserId)
	if err != nil {
		return nil, err
	}

	var affected int64
	if !has {
		affected, err = myUser.InsertMyUserInTxn(session)
	} else {
		myUserModel.Name = myUser.Name
		affected, err = myUserModel.UpdateMyUserInTxn(session)
	}

	if err != nil {
		return -1, err
	}
	logrus.Infof("affected is: %#v\n", affected)
	return affected, nil
}

// Optimistic Lock
func AddTryCount(session *xorm.Session, myUser *model.MyUser) (interface{}, error) {
	var myUserModel *model.MyUser
	myUserModel, _, err := model.GetMyUserInTxn(session, myUser.UserId)
	if err != nil {
		return nil, err
	}
	logrus.Infof("myUserModel is: %#v\n", myUserModel)
	logrus.Infof("You can try update in another session in 5 seconds, you can succeed and this transaction fail")
	for i := 1; i <= 5; i++ {
		logrus.Print(".")
		time.Sleep(5 * time.Second)
	}

	myUserModel.Name = myUser.Name
	affected, err := myUserModel.UpdateMyUserInTxn(session)
	if err != nil {
		return -1, err
	}
	logrus.Infof("affected is: %#v\n", affected)
	return affected, nil
}

// Pessimistic Lock
func AddTryCountWithLock(session *xorm.Session, myUser *model.MyUser) (interface{}, error) {
	logrus.Infof("Lock created")
	var myUserModel *model.MyUser
	myUserModel, _, err := model.GetMyUserForUpdateInTxn(session, myUser.UserId)
	if err != nil {
		return nil, err
	}
	logrus.Infof("myUserModel is: %#v\n", myUserModel)
	logrus.Infof("You can try update in another session in 5 seconds, you'll fail and this transaction succeed")
	for i := 1; i <= 5; i++ {
		logrus.Print(".")
		time.Sleep(5 * time.Second)
	}

	myUserModel.Name = myUser.Name
	affected, err := myUserModel.UpdateMyUserInTxn(session)
	if err != nil {
		return nil, err
	}
	logrus.Infof("affected is: %#v\n", affected)
	logrus.Infof("Lock released")
	return affected, nil
}

func indexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		allMyUser, err := GetAllMyUser(MyUserDBEngine.NewSession())
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "register.html", gin.H{"affected": 0, "msgColor": "bg-info text-dark", "allMyUser": allMyUser})
	}
}

func registerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var myuser model.MyUser
		err := c.ShouldBind(&myuser)
		if err != nil {
			return
		}
		affected, err := MyUserDBEngine.Transaction(InsertOrUpdate, &myuser)
		if err != nil {
			panic(err)
		}
		var msgColor string
		if affected == int64(1) {
			msgColor = "bg-success text-white"
		} else {
			msgColor = "bg-danger text-white"
		}

		allMyUser, err := GetAllMyUser(MyUserDBEngine.NewSession())
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "register.html", gin.H{"myuser": myuser, "affected": affected, "msgColor": msgColor, "allMyUser": allMyUser})
	}
}

func addTryCountHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var myuser model.MyUser
		err := c.ShouldBind(&myuser)
		if err != nil {
			return
		}
		affected, err := MyUserDBEngine.Transaction(AddTryCount, &myuser)
		if err != nil {
			panic(err)
		}
		var msgColor string
		if affected == int64(1) {
			msgColor = "bg-success text-white"
		} else {
			msgColor = "bg-danger text-white"
		}

		allMyUser, err := GetAllMyUser(MyUserDBEngine.NewSession())
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "register.html", gin.H{"myuser": myuser, "affected": affected, "msgColor": msgColor, "allMyUser": allMyUser})
	}
}

func addTryCountWithLockHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var myuser model.MyUser
		err := c.ShouldBind(&myuser)
		if err != nil {
			return
		}
		affected, err := MyUserDBEngine.Transaction(AddTryCountWithLock, &myuser)
		if err != nil {
			panic(err)
		}
		var msgColor string
		if affected == int64(1) {
			msgColor = "bg-success text-white"
		} else {
			msgColor = "bg-danger text-white"
		}

		allMyUser, err := GetAllMyUser(MyUserDBEngine.NewSession())
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "register.html", gin.H{"myuser": myuser, "affected": affected, "msgColor": msgColor, "allMyUser": allMyUser})
	}
}

func initMyUserDBEngine() {
	MyUserDBEngine.Engine = config.DBEngine
	// set context with timeout
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	//defer cancel()
	//MyUserDBEngine.Engine.SetDefaultContext(ctx)
	// create table
	err := MyUserDBEngine.Sync(new(model.MyUser))
	if err != nil {
		panic(err)
	}
}

var MyUserDBEngine model.MyUserEngine

func main() {

	logrus.SetLevel(logrus.InfoLevel)
	initMyUserDBEngine()

	//// insert
	//var newMyUser model.MyUser
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
	r.GET("/", indexHandler())
	r.POST("/register", registerHandler())
	r.POST("/update", addTryCountHandler())
	r.POST("/updatewithlock", addTryCountWithLockHandler())
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
