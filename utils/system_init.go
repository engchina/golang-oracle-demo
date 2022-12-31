package utils

import (
	"fmt"
	_ "github.com/godror/godror"
	"github.com/spf13/viper"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

var (
	DBEngine     *xorm.Engine
	errNewEngine error
)

// init config
func InitConfig() {
	viper.SetConfigName("application")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error in read config file %w", err))
	}
}

// init database connection
func InitOracle() {
	driverName := viper.GetString("oracle.driverName")
	dataSourceName := viper.GetString("oracle.dataSourceName")
	DBEngine, errNewEngine = xorm.NewEngine(driverName, dataSourceName)
	if errNewEngine != nil {
		panic(fmt.Errorf("error in init new engine %w", errNewEngine))
	}

	err := DBEngine.Ping()
	if err != nil {
		panic(fmt.Errorf("error on ping db: %w", err))
	}

	DBEngine.ShowSQL(false)
	//DBEngine.Logger().SetLevel(log.LOG_DEBUG)
	DBEngine.Logger().SetLevel(log.LOG_INFO)
	DBEngine.SetTableMapper(names.GonicMapper{})
	DBEngine.SetColumnMapper(names.GonicMapper{})

	//DBEngine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	DBEngine.TZLocation, _ = time.LoadLocation("Asia/Tokyo")
	DBEngine.SetMaxOpenConns(5)
	DBEngine.SetMaxIdleConns(2)
	DBEngine.SetConnMaxLifetime(10 * time.Minute)
}

//func init() {
//	initConfig()
//	initOracle()
//}
