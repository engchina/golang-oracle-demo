package utils

import (
	"github.com/engchina/golang-oracle-demo/models"
)

var (
	MyUserDBEngine models.MyUserEngine
)

func initMyUserDBEngine() {
	MyUserDBEngine.Engine = DBEngine
}

func init() {
	initConfig()
	initOracle()
	initMyUserDBEngine()

	// set context with timeout
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	//defer cancel()
	//MyUserDBEngine.Engine.SetDefaultContext(ctx)

	// create table
	//err := MyUserDBEngine.Sync(new(models.MyUser))
	//if err != nil {
	//	panic(err)
	//}
}
