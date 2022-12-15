package main

import (
	"fmt"
	"github.com/engchina/golang-oracle-demo/model"
	"github.com/engchina/golang-oracle-demo/utils"
	_ "github.com/godror/godror"
	"time"
	"xorm.io/xorm"
)

const (
	timeTemplate = "2006-01-02 15:04:05.000000"
)

func handleInTransaction(session *xorm.Session) (interface{}, error) {
	// create table, run once
	//err := session.Sync(new(model.TblUser))
	//if err != nil {
	//	panic(fmt.Errorf("error creating table: %w", err))
	//}

	resp := make(map[string]interface{})

	var user model.TblUser
	t := time.Now().Format(timeTemplate)
	user = model.TblUser{
		Id:      16,
		Name:    "Tom",
		Version: 1,
		Created: t,
		Updated: t,
	}

	affected, err := user.InsertTblUserInSession(session)
	if err != nil {
		return nil, err
	}
	fmt.Printf("affected is: %#v\n", affected)

	var userModel *model.TblUser
	userModel, _, err = model.GetTblUserInSession(session, user.Id)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user is: %#v\n", userModel)
	time.Sleep(1 * time.Second)

	userModel.Name = "John"
	affected, err = userModel.UpdateTblUserInSession(session)
	if err != nil {
		return nil, err
	}
	fmt.Printf("affected is: %#v\n", affected)
	time.Sleep(1 * time.Second)

	userModel, _, err = model.GetTblUserInSession(session, user.Id)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user is: %#v\n", userModel)

	resp["data"] = "success"
	return resp, nil
}

func Insert(session *xorm.Session, userId int64) (interface{}, error) {
	resp := make(map[string]interface{})

	var user model.TblUser
	t := time.Now().Format(timeTemplate)
	user = model.TblUser{
		Id:      userId,
		Name:    "Tom",
		Version: 1,
		Created: t,
		Updated: t,
	}

	affected, err := user.InsertTblUserInSession(session)
	if err != nil {
		return nil, err
	}
	fmt.Printf("affected is: %#v\n", affected)

	resp["data"] = "success"
	return resp, nil
}

func AddTryCount(session *xorm.Session, userId int64) (interface{}, error) {
	resp := make(map[string]interface{})

	t := time.Now().Format(timeTemplate)

	var userModel *model.TblUser
	userModel, _, err := model.GetTblUserInSession(session, userId)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user is: %#v\n", userModel)
	time.Sleep(2 * time.Second)

	userModel.Name = "opc"
	userModel.Created = t
	userModel.Updated = t
	numOfTried := userModel.Version
	userModel.Version = numOfTried + 1
	affected, err := userModel.UpdateTblUserInSessionWithTried(session, numOfTried)
	if err != nil {
		return nil, err
	}
	fmt.Printf("affected is: %#v\n", affected)
	time.Sleep(2 * time.Second)

	userModel, _, err = model.GetTblUserInSession(session, userId)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user is: %#v\n", userModel)

	resp["data"] = "success"
	return resp, nil
}

func main() {
	var MyDBEngine model.MyEngine
	MyDBEngine.Engine = utils.DBEngine
	//resp, err := utils.DBEngine.Transaction(handleInTransaction)
	var newUserId int64
	newUserId = 26
	resp, err := MyDBEngine.Transaction(Insert, newUserId)
	if err != nil {
		panic("err: " + err.Error())
	}
	resp, err = MyDBEngine.Transaction(AddTryCount, newUserId)
	if err != nil {
		panic("err: " + err.Error())
	}
	fmt.Printf("resp: %#v\n", resp)
}
