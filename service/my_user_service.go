package service

import (
	"github.com/engchina/golang-oracle-demo/models"
	"github.com/sirupsen/logrus"
	"time"
	"xorm.io/xorm"
)

// GetMyUserList Get MyUser List
func GetMyUserList(session *xorm.Session) (interface{}, error) {
	var allData []*models.MyUser
	allData, err := models.GetMyUserList(session)
	if err != nil {
		return nil, err
	}
	return allData, nil
}

// InsertOrUpdate Insert or Update
func InsertOrUpdate(session *xorm.Session, in interface{}) (interface{}, error) {
	myUser := in.(*models.MyUser)
	var myUserModel *models.MyUser
	myUserModel, has, err := models.GetMyUserInTxn(session, myUser.UserId)
	if err != nil {
		return nil, err
	}

	var affected int64
	if !has {
		affected, err = myUser.InsertMyUserInTxn(session)
	} else {
		myUserModel.Name = myUser.Name
		affected, err = myUserModel.UpdateMyUserInTxn(session)
	}

	if err != nil {
		return -1, err
	}
	return affected, nil
}

// UpdateWithOptimisticLock Optimistic Lock
func UpdateWithOptimisticLock(session *xorm.Session, in interface{}) (interface{}, error) {
	myUser := in.(*models.MyUser)
	var myUserModel *models.MyUser
	myUserModel, _, err := models.GetMyUserInTxn(session, myUser.UserId)
	if err != nil {
		return nil, err
	}
	logrus.Infof("myUserModel is: %#v\n", myUserModel)
	time.Sleep(5 * time.Second)

	myUserModel.Name = myUser.Name
	affected, err := myUserModel.UpdateMyUserInTxn(session)
	if err != nil {
		return -1, err
	}
	return affected, nil
}

// UpdateWithPessimisticLock Pessimistic Lock
func UpdateWithPessimisticLock(session *xorm.Session, in interface{}) (interface{}, error) {
	myUser := in.(*models.MyUser)
	var myUserModel *models.MyUser
	myUserModel, _, err := models.GetMyUserForUpdateInTxn(session, myUser.UserId)
	if err != nil {
		return nil, err
	}
	time.Sleep(5 * time.Second)

	myUserModel.Name = myUser.Name
	affected, err := myUserModel.UpdateMyUserInTxn(session)
	if err != nil {
		return nil, err
	}
	return affected, nil
}
