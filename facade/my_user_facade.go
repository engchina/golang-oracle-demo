package facade

import (
	"github.com/engchina/golang-oracle-demo/db"
	"github.com/engchina/golang-oracle-demo/models"
	"github.com/engchina/golang-oracle-demo/service"
)

func GetMyUserList() (interface{}, error) {
	return db.OracleClient.ReadOnlyTransaction(service.GetMyUserList)
}

func InsertOrUpdate(myUser *models.MyUser) (interface{}, error) {
	return db.OracleClient.ReadWriteTransaction(service.InsertOrUpdate, myUser)
}

func UpdateWithOptimisticLock(myUser *models.MyUser) (interface{}, error) {
	return db.OracleClient.ReadWriteTransaction(service.UpdateWithOptimisticLock, myUser)
}

func UpdateWithPessimisticLock(myUser *models.MyUser) (interface{}, error) {
	return db.OracleClient.ReadWriteTransaction(service.UpdateWithPessimisticLock, myUser)
}
