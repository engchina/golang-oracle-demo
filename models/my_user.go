package models

import (
	_ "github.com/godror/godror"
	"time"
	"xorm.io/xorm"
)

const (
	MyUserTableName = "MY_USER"
)

type MyUser struct {
	UserId     string    `json:"userId"  xorm:"varchar(200) pk 'USER_ID'"   form:"userId"`
	Name       string    `json:"name"    xorm:"varchar(200) notnull 'NAME'" form:"name"`
	NumOfTried int64     `json:"version" xorm:"version 'NUM_OF_TRIED'"`
	Created    time.Time `json:"created" xorm:"created 'CREATED'"`
	Updated    time.Time `json:"updated" xorm:"updated 'UPDATED'"`
	Deleted    time.Time `json:"deleted" xorm:"deleted 'DELETED'"`
}

func (*MyUser) TableName() string {
	return MyUserTableName
}

func GetMyUserList(session *xorm.Session) ([]*MyUser, error) {
	allData := make([]*MyUser, 0)
	err := session.Table(MyUserTableName).OrderBy("user_id").Find(&allData)
	return allData, err
}

func GetMyUserInTxn(session *xorm.Session, userId string) (*MyUser, bool, error) {
	myUser := new(MyUser)
	has, err := session.Table(MyUserTableName).ID(userId).Get(myUser)
	return myUser, has, err
}

func GetMyUserForUpdateInTxn(session *xorm.Session, userId string) (*MyUser, bool, error) {
	myUser := new(MyUser)
	has, err := session.ForUpdate().Table(MyUserTableName).ID(userId).Get(myUser)
	return myUser, has, err
}

func (myUser *MyUser) InsertMyUserInTxn(session *xorm.Session) (int64, error) {
	count, err := session.Table(MyUserTableName).Insert(myUser)
	return count, err
}

func (myUser *MyUser) UpdateMyUserInTxn(session *xorm.Session) (int64, error) {
	count, err := session.Table(MyUserTableName).ID(myUser.UserId).Update(myUser)
	return count, err
}
