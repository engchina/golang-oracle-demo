package model

import (
	_ "github.com/godror/godror"
	"xorm.io/xorm"
)

const (
	TblUserTableName = "TBL_USER"
)

type TblUser struct {
	Id         int64  `json:"userId" xorm:"pk 'USER_ID'"`
	Name       string `json:"name" xorm:"varchar(200) notnull 'NAME'"`
	NumOfTried int64  `json:"numOfTried" xorm:"notnull 'NUM_OF_TRIED'"`
	Created    string `json:"createdTime" xorm:"notnull 'CREATED_TIME'"`
	Updated    string `json:"updatedTime" xorm:"notnull 'UPDATED_TIME'"`
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
	has, err := session.Where("USER_ID = :1", userId).Get(user)
	return user, has, err
}

func (user *TblUser) UpdateTblUserInSession(session *xorm.Session) (int64, error) {
	count, err := session.Table(TblUserTableName).AllCols().Where("USER_ID = :1", user.Id).Update(user)
	return count, err
}
