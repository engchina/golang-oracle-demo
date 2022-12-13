package main

import (
	"fmt"
	_ "github.com/godror/godror"
	"time"
	"xorm.io/xorm"
)

const (
	timeTemplate = "2006-01-02 15:04:05.000000"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("godror", "oracle://pdbadmin:oracle@192.168.31.23:1521/pdb1")
	//engine, err = xorm.NewEngine("godror", `user="admin" password="ToDo" connectString="tcps://adb.ap-singapore-1.oraclecloud.com:1522/g56e4c08bfdf01c_myatp_low.adb.oraclecloud.com?wallet_location=/u01/wallets/Wallet_myatp/"`)

	if err != nil {
		panic(fmt.Errorf("error in xorm.NewEngine: %w", err))
	}
	defer func() {
		err = engine.Close()
		if err != nil {
			fmt.Println("error on can't close connection: ", err)
		}
	}()

	err = engine.Ping()
	if err != nil {
		panic(fmt.Errorf("error on ping db: %w", err))
	}

	res, err := engine.Transaction(func(session *xorm.Session) (interface{}, error) {
		/**
		godror WARNING: discrepancy between DBTIMEZONE ("+00:00"=0) and SYSTIMESTAMP ("+08:00"=800) - set connection timezone, see https://github.com/godror/godror/blob/master/doc/timezone.md
		*/
		//err = engine.Sync(new(TblUser))
		//if err != nil {
		//	panic(fmt.Errorf("error creating table: %w", err))
		//}

		var tblUser TblUser
		t := time.Now().Format(timeTemplate)
		tblUser = TblUser{
			Id:         4,
			Name:       "Tom",
			Salt:       "123456",
			Age:        18,
			Passwd:     "123456",
			NumOfTried: 0,
			Created:    t,
			Updated:    t,
		}
		affected, err := session.Insert(&tblUser)
		if err != nil {
			panic(fmt.Errorf("error on insert data: %w", err))
		}
		fmt.Printf("affected is %d\n", affected)

		// on another sqlcli, select * from where user_id = ?, repeat 1000 1
		time.Sleep(20 * time.Second)

		results, err := session.Where("USER_ID = 4").QueryInterface("select * from TBL_USER")
		if err != nil {
			panic(fmt.Errorf("error on query data: %w", err))
		}
		fmt.Printf("%#v\n", results)

		return nil, nil
	})

	fmt.Println(res)

	fmt.Println("okay complete")
}
