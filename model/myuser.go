package model

import (
	_ "github.com/godror/godror"
	"time"
	"xorm.io/xorm"
)

const (
	MyUserTableName = "MY_USER"
)

type MyUser struct {
	UserId     string    `json:"userId"  xorm:"varchar(200) pk 'USER_ID'"`
	Name       string    `json:"name"    xorm:"varchar(200) notnull 'NAME'"`
	NumOfTried int64     `json:"version" xorm:"version 'NUM_OF_TRIED'"`
	Created    time.Time `json:"created" xorm:"created 'CREATED'"`
	Updated    time.Time `json:"updated" xorm:"updated 'UPDATED'"`
}

func (MyUser) TableName() string {
	return MyUserTableName
}

type MyUserEngine struct {
	*xorm.Engine
}

func (myUser *MyUser) InsertMyUserInTxn(session *xorm.Session) (int64, error) {
	count, err := session.Table(MyUserTableName).Insert(myUser)
	return count, err
}

func GetMyUserInTxn(session *xorm.Session, userId string) (*MyUser, bool, error) {
	myUser := new(MyUser)
	has, err := session.Table(MyUserTableName).ID(userId).Get(myUser)
	return myUser, has, err
}

func (myUser *MyUser) UpdateMyUserInTxn(session *xorm.Session) (int64, error) {
	count, err := session.Table(MyUserTableName).ID(myUser.UserId).Update(myUser)
	return count, err
}

func (engine *MyUserEngine) Transaction(f func(*xorm.Session, *MyUser) (interface{}, error), myUser *MyUser) (interface{}, error) {
	session := engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return nil, err
	}

	result, err := f(session, myUser)
	if err != nil {
		return result, err
	}

	if err := session.Commit(); err != nil {
		return result, err
	}

	return result, nil
}
