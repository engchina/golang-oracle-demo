package facade

import (
	"github.com/engchina/golang-oracle-demo/models"
	"github.com/engchina/golang-oracle-demo/service"
	"github.com/engchina/golang-oracle-demo/utils"
)

func GetMyUserList() (interface{}, error) {
	return utils.MyUserDBEngine.ReadOnlyTransaction(service.GetMyUserList)
}

func InsertOrUpdate(myUser *models.MyUser) (interface{}, error) {
	return utils.MyUserDBEngine.ReadWriteTransaction(service.InsertOrUpdate, myUser)
}

func UpdateWithOptimisticLock(myUser *models.MyUser) (interface{}, error) {
	return utils.MyUserDBEngine.ReadWriteTransaction(service.UpdateWithOptimisticLock, myUser)
}

func UpdateWithPessimisticLock(myUser *models.MyUser) (interface{}, error) {
	return utils.MyUserDBEngine.ReadWriteTransaction(service.UpdateWithPessimisticLock, myUser)
}
