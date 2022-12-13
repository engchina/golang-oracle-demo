package service

import (
	"xorm.io/xorm"
)

const TblUserTableName = "TBL_USER"

func (user *TblUser) InsertTblUserInSession(session *xorm.Session) (int64, error) {
	count, err := session.Table(TblUserTableName).Insert(user)
	return count, err
}
