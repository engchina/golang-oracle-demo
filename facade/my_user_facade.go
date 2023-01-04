package facade

import (
	"github.com/engchina/golang-oracle-demo/models"
	"github.com/engchina/golang-oracle-demo/service"
	"github.com/engchina/golang-oracle-demo/utils"
)

func GetMyUserList() (interface{}, error) {
	return utils.OracleDBEngine.ReadOnlyTransaction(service.GetMyUserList)
}

func InsertOrUpdate(myUser *models.MyUser) (interface{}, error) {
	return utils.OracleDBEngine.ReadWriteTransaction(service.InsertOrUpdate, myUser)
}

func UpdateWithOptimisticLock(myUser *models.MyUser) (interface{}, error) {
	return utils.OracleDBEngine.ReadWriteTransaction(service.UpdateWithOptimisticLock, myUser)
}

func UpdateWithPessimisticLock(myUser *models.MyUser) (interface{}, error) {
	return utils.OracleDBEngine.ReadWriteTransaction(service.UpdateWithPessimisticLock, myUser)
}
