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
		Id:         6,
		Name:       "Tom",
		NumOfTried: 1,
		Created:    t,
		Updated:    t,
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
	time.Sleep(5 * time.Second)

	userModel.Name = "John"
	affected, err = userModel.UpdateTblUserInSession(session)
	if err != nil {
		return nil, err
	}
	fmt.Printf("affected is: %#v\n", affected)
	time.Sleep(5 * time.Second)

	resp["data"] = "success"
	return resp, nil
}

func main() {
	resp, err := utils.DBEngine.Transaction(handleInTransaction)
	if err != nil {
		panic("err: " + err.Error())
	}
	fmt.Printf("resp: %#v\n", resp)
}

//var engine *xorm.Engine
//
//func main() {
//	var err error
//	engine, err = xorm.NewEngine("godror", "oracle://pdbadmin:oracle@192.168.31.23:1521/pdb1")
//	//engine, err = xorm.NewEngine("godror", `user="admin" password="ToDo" connectString="tcps://adb.ap-singapore-1.oraclecloud.com:1522/g56e4c08bfdf01c_myatp_low.adb.oraclecloud.com?wallet_location=/u01/wallets/Wallet_myatp/"`)
//
//	if err != nil {
//		panic(fmt.Errorf("error in xorm.NewEngine: %w", err))
//	}
//	defer func() {
//		err = engine.Close()
//		if err != nil {
//			fmt.Println("error on can't close connection: ", err)
//		}
//	}()
//
//	err = engine.Ping()
//	if err != nil {
//		panic(fmt.Errorf("error on ping db: %w", err))
//	}
//
//	res, err := engine.Transaction(func(session *xorm.Session) (interface{}, error) {
//		/**
//		godror WARNING: discrepancy between DBTIMEZONE ("+00:00"=0) and SYSTIMESTAMP ("+08:00"=800) - set connection timezone, see https://github.com/godror/godror/blob/master/doc/timezone.md
//		*/
//		//err = engine.Sync(new(TblUser))
//		//if err != nil {
//		//	panic(fmt.Errorf("error creating table: %w", err))
//		//}
//
//		var tblUser model.TblUser
//		t := time.Now().Format(timeTemplate)
//		tblUser = model.TblUser{
//			Id:         4,
//			Name:       "Tom",
//			NumOfTried: 0,
//			Created:    t,
//			Updated:    t,
//		}
//		affected, err := session.Insert(&tblUser)
//		if err != nil {
//			panic(fmt.Errorf("error on insert data: %w", err))
//		}
//		fmt.Printf("affected is %d\n", affected)
//
//		// on another sqlcli, select * from where user_id = ?, repeat 1000 1
//		time.Sleep(20 * time.Second)
//
//		results, err := session.Where("USER_ID = 4").QueryInterface("select * from TBL_USER")
//		if err != nil {
//			panic(fmt.Errorf("error on query data: %w", err))
//		}
//		fmt.Printf("%#v\n", results)
//
//		return nil, nil
//	})
//
//	fmt.Println(res)
//
//	fmt.Println("okay complete")
//}
