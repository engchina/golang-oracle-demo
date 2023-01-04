package service

import (
	"github.com/engchina/golang-oracle-demo/models"
	"github.com/sirupsen/logrus"
	"time"
	"xorm.io/xorm"
)

func GetMyUserList(session *xorm.Session) (interface{}, error) {
	var allData []*models.MyUser
	allData, err := models.GetMyUserList(session)
	if err != nil {
		return nil, err
	}
	return allData, nil
}

//func Insert(session *xorm.Session, myUser *models.MyUser) (interface{}, error) {
//	affected, err := myUser.InsertMyUserInTxn(session)
//	if err != nil {
//		return -1, err
//	}
//	logrus.Infof("affected is: %#v\n", affected)
//	return affected, nil
//}

// func InsertOrUpdate(session *xorm.Session, myUser *models.MyUser) (interface{}, error) {
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

// Optimistic Lock
func UpdateWithOptimisticLock(session *xorm.Session, myInterface interface{}) (interface{}, error) {
	myUser := myInterface.(*models.MyUser)
	var myUserModel *models.MyUser
	myUserModel, _, err := models.GetMyUserInTxn(session, myUser.UserId)
	if err != nil {
		return nil, err
	}
	logrus.Infof("myUserModel is: %#v\n", myUserModel)
	logrus.Infof("You can try update in another session in 5 seconds, you can succeed and this transaction fail")
	for i := 1; i <= 5; i++ {
		logrus.Print(".")
		time.Sleep(5 * time.Second)
	}

	myUserModel.Name = myUser.Name
	affected, err := myUserModel.UpdateMyUserInTxn(session)
	if err != nil {
		return -1, err
	}
	logrus.Infof("affected is: %#v\n", affected)
	return affected, nil
}

// Pessimistic Lock
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
	for i := 1; i <= 5; i++ {
		logrus.Print(".")
		time.Sleep(5 * time.Second)
	}

	myUserModel.Name = myUser.Name
	affected, err := myUserModel.UpdateMyUserInTxn(session)
	if err != nil {
		return nil, err
	}
	logrus.Infof("affected is: %#v\n", affected)
	logrus.Infof("Lock released")
	return affected, nil
}
