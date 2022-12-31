package main

import (
	"fmt"
	"github.com/engchina/golang-oracle-demo/router"
	"github.com/engchina/golang-oracle-demo/utils"
	_ "github.com/godror/godror"
	"github.com/sirupsen/logrus"
)

func main() {
	utils.InitConfig()
	utils.InitOracle()
	utils.InitMyUserDBEngine()

	logrus.SetLevel(logrus.InfoLevel)

	r := router.Router()
	err := r.Run(":3000")
	if err != nil {
		panic(fmt.Errorf("error in start gin server %w", err))
	}
}
