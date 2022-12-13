package main

import (
	"fmt"
	_ "github.com/godror/godror"
	"time"
	"xorm.io/xorm"
)

type TblUser struct {
	Id         int64  `xorm:"USER_ID"`
	Name       string `xorm:"NAME"`
	Salt       string `xorm:"SALT"`
	Age        int    `xorm:"AGE"`
	Passwd     string `xorm:"varchar(200) PASSWD"`
	NumOfTried int64  `xorm:"NUM_OF_TRIED"`
	Created    string `xorm:"CREATED_TIME"`
	Updated    string `xorm:"UPDATED_TIME"`
}

func (TblUser) TableName() string {
	return "TBL_USER"
}

func (TblUser) PK() []string {
	return []string{"USER_ID"}
}

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
			fmt.Println("Can't close connection: ", err)
		}
	}()

	err = engine.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}

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
		Id:         2,
		Name:       "Tom",
		Salt:       "123456",
		Age:        18,
		Passwd:     "123456",
		NumOfTried: 0,
		Created:    t,
		Updated:    t,
	}
	affected, err := engine.Insert(&tblUser)
	if err != nil {
		panic(fmt.Errorf("error inserting db: %w", err))
	}
	fmt.Printf("affected is %d\n", affected)

	fmt.Println("complete successfully")
}
