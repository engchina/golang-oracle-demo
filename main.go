package main

import (
	"fmt"
	"github.com/engchina/golang-oracle-demo/model"
	"github.com/engchina/golang-oracle-demo/utils"
	_ "github.com/godror/godror"
	"time"
	"xorm.io/xorm"
)

func Insert(session *xorm.Session, myUser *model.MyUser) (interface{}, error) {
	resp := make(map[string]interface{})
	affected, err := myUser.InsertMyUserInTxn(session)
	if err != nil {
		return nil, err
	}
	fmt.Printf("affected is: %#v\n", affected)

	resp["data"] = "success"
	return resp, nil
}

func AddTryCount(session *xorm.Session, myUser *model.MyUser) (interface{}, error) {
	var myUserModel *model.MyUser
	myUserModel, _, err := model.GetMyUserInTxn(session, myUser.UserId)
	if err != nil {
		return nil, err
	}
	fmt.Printf("myUserModel is: %#v\n", myUserModel)
	time.Sleep(2 * time.Second)

	myUserModel.Name = myUser.Name
	affected, err := myUserModel.UpdateMyUserInTxn(session)
	if err != nil {
		return nil, err
	}
	fmt.Printf("affected is: %#v\n", affected)
	time.Sleep(2 * time.Second)

	return "success", nil
}

func main() {
	var MyUserDBEngine model.MyUserEngine
	MyUserDBEngine.Engine = utils.DBEngine
	// create table
	err := MyUserDBEngine.Sync(new(model.MyUser))
	if err != nil {
		panic(err)
	}

	// insert
	var newMyUser model.MyUser
	newMyUser.UserId = "20"
	newMyUser.Name = "first"
	resp, err := MyUserDBEngine.Transaction(Insert, &newMyUser)
	if err != nil {
		panic(err)
	}

	// update
	newMyUser.Name = "second"
	resp, err = MyUserDBEngine.Transaction(AddTryCount, &newMyUser)
	if err != nil {
		panic("err: " + err.Error())
	}
	fmt.Printf("resp: %#v\n", resp)

	// update
	newMyUser.Name = "third"
	resp, err = MyUserDBEngine.Transaction(AddTryCount, &newMyUser)
	if err != nil {
		panic("err: " + err.Error())
	}
	fmt.Printf("resp: %#v\n", resp)

	// update
	newMyUser.Name = "fourth"
	resp, err = MyUserDBEngine.Transaction(AddTryCount, &newMyUser)
	if err != nil {
		panic("err: " + err.Error())
	}
	fmt.Printf("resp: %#v\n", resp)
}
