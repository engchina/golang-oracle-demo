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
func InsertOrUpdate(session *xorm.Session, myInterface interface{}) (interface{}, error) {
	myUser := myInterface.(*models.MyUser)
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
	logrus.Infof("affected is: %#v\n", affected)
	return affected, nil
}

// UpdateWithOptimisticLock Optimistic Lock
func UpdateWithOptimisticLock(session *xorm.Session, myInterface interface{}) (interface{}, error) {
	myUser := myInterface.(*models.MyUser)
	var myUserModel *models.MyUser
	myUserModel, _, err := models.GetMyUserInTxn(session, myUser.UserId)
	if err != nil {
		return nil, err
	}
	logrus.Infof("myUserModel is: %#v\n", myUserModel)
	logrus.Infof("You can try update in another session in 5 seconds, you can succeed and this transaction fail")
	time.Sleep(5 * time.Second)

	myUserModel.Name = myUser.Name
	affected, err := myUserModel.UpdateMyUserInTxn(session)
	if err != nil {
		return -1, err
	}
	logrus.Infof("affected is: %#v\n", affected)
	return affected, nil
}

// UpdateWithPessimisticLock Pessimistic Lock
func UpdateWithPessimisticLock(session *xorm.Session, myInterface interface{}) (interface{}, error) {
	logrus.Infof("Lock created")
	myUser := myInterface.(*models.MyUser)
	var myUserModel *models.MyUser
	myUserModel, _, err := models.GetMyUserForUpdateInTxn(session, myUser.UserId)
	if err != nil {
		return nil, err
	}
	logrus.Infof("myUserModel is: %#v\n", myUserModel)
	logrus.Infof("You can try update in another session in 5 seconds, you'll fail and this transaction succeed")
	time.Sleep(5 * time.Second)

	myUserModel.Name = myUser.Name
	affected, err := myUserModel.UpdateMyUserInTxn(session)
	if err != nil {
		return nil, err
	}
	logrus.Infof("affected is: %#v\n", affected)
	logrus.Infof("Lock released")
	return affected, nil
}
