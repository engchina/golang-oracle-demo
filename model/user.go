package model

import (
	_ "github.com/godror/godror"
	"xorm.io/xorm"
)

const (
	TblUserTableName = "TBL_USER"
)

type TblUser struct {
	Id      int64  `json:"userId" xorm:"pk 'USER_ID'"`
	Name    string `json:"name" xorm:"varchar(200) notnull 'NAME'"`
	Version int64  `json:"version" xorm:"notnull 'VERSION'"`
	Created string `json:"createdTime" xorm:"notnull 'CREATED_TIME'"`
	Updated string `json:"updatedTime" xorm:"notnull 'UPDATED_TIME'"`
}

type MyEngine struct {
	*xorm.Engine
}

func (TblUser) TableName() string {
	return TblUserTableName
}

func (user *TblUser) InsertTblUserInSession(session *xorm.Session) (int64, error) {
	count, err := session.Table(TblUserTableName).Insert(user)
	return count, err
}

func GetTblUserInSession(session *xorm.Session, userId int64) (*TblUser, bool, error) {
	user := new(TblUser)
	has, err := session.Table(TblUserTableName).ID(userId).Get(user)
	return user, has, err
}

func (user *TblUser) UpdateTblUserInSession(session *xorm.Session) (int64, error) {
	count, err := session.Table(TblUserTableName).ID(user.Id).Update(user)
	return count, err
}

func (user *TblUser) UpdateTblUserInSessionWithTried(session *xorm.Session, numOfTried int64) (int64, error) {
	count, err := session.Table(TblUserTableName).ID(user.Id).Where("VERSION = :1", numOfTried).Update(user)
	return count, err
}

func (engine *MyEngine) Transaction(f func(*xorm.Session, int64) (interface{}, error), userId int64) (interface{}, error) {
	session := engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return nil, err
	}

	result, err := f(session, userId)
	if err != nil {
		return result, err
	}

	if err := session.Commit(); err != nil {
		return result, err
	}

	return result, nil
}
