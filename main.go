package main

import (
	"fmt"
	"github.com/engchina/golang-oracle-demo/config"
	"github.com/engchina/golang-oracle-demo/model"
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

// Optimistic Lock
func AddTryCount(session *xorm.Session, myUser *model.MyUser) (interface{}, error) {
	var myUserModel *model.MyUser
	myUserModel, _, err := model.GetMyUserInTxn(session, myUser.UserId)
	if err != nil {
		return nil, err
	}
	fmt.Printf("myUserModel is: %#v\n", myUserModel)
	fmt.Println("You can try update in another session in 5 seconds, you can succeed and this transaction fail")
	for i := 1; i <= 5; i++ {
		fmt.Print(".")
		time.Sleep(5 * time.Second)
	}

	myUserModel.Name = myUser.Name
	affected, err := myUserModel.UpdateMyUserInTxn(session)
	if err != nil {
		return nil, err
	}
	fmt.Printf("affected is: %#v\n", affected)
	if affected == 0 {
		fmt.Println("This transaction fail")
		return "fail", nil
	} else {
		fmt.Println("This transaction succeed")
		return "success", nil
	}
}

// Pessimistic Lock
func AddTryCountWithLock(session *xorm.Session, myUser *model.MyUser) (interface{}, error) {
	fmt.Println("Lock created")
	var myUserModel *model.MyUser
	myUserModel, _, err := model.GetMyUserForUpdateInTxn(session, myUser.UserId)
	if err != nil {
		return nil, err
	}
	fmt.Printf("myUserModel is: %#v\n", myUserModel)
	fmt.Println("You can try update in another session in 5 seconds, you'll fail and this transaction succeed")
	for i := 1; i <= 5; i++ {
		fmt.Print(".")
		time.Sleep(5 * time.Second)
	}

	myUserModel.Name = myUser.Name
	affected, err := myUserModel.UpdateMyUserInTxn(session)
	if err != nil {
		return nil, err
	}
	fmt.Printf("affected is: %#v\n", affected)
	fmt.Println("Lock released")
	if affected == 0 {
		fmt.Println("This transaction fail")
		return "fail", nil
	} else {
		fmt.Println("This transaction succeed")
		return "success", nil
	}
}

func main() {
	var MyUserDBEngine model.MyUserEngine
	MyUserDBEngine.Engine = config.DBEngine
	// create table
	err := MyUserDBEngine.Sync(new(model.MyUser))
	if err != nil {
		panic(err)
	}

	// insert
	var newMyUser model.MyUser
	newMyUser.UserId = "120"
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
	resp, err = MyUserDBEngine.Transaction(AddTryCountWithLock, &newMyUser)
	if err != nil {
		panic("err: " + err.Error())
	}
	fmt.Printf("resp: %#v\n", resp)

	// update
	newMyUser.Name = "fifth"
	resp, err = MyUserDBEngine.Transaction(AddTryCountWithLock, &newMyUser)
	if err != nil {
		panic("err: " + err.Error())
	}
	fmt.Printf("resp: %#v\n", resp)

	// update
	newMyUser.Name = "sixth"
	resp, err = MyUserDBEngine.Transaction(AddTryCountWithLock, &newMyUser)
	if err != nil {
		panic("err: " + err.Error())
	}
	fmt.Printf("resp: %#v\n", resp)
}

/*
select * from MY_USER where USER_ID = '9999';

update MY_USER set NAME='another session', NUM_OF_TRIED=NUM_OF_TRIED+1 where USER_ID = '9999';
commit;

select * from MY_USER where USER_ID = '9999';
*/
