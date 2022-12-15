package utils

import (
	"fmt"
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

func init() {
	// init config
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error in read config file %w", err))
	}

	// init database connection
	driverName := viper.GetString("driverName")
	dataSourceName := viper.GetString("dataSourceName")
	DBEngine, errNewEngine = xorm.NewEngine(driverName, dataSourceName)
	if errNewEngine != nil {
		panic(fmt.Errorf("error in init new engine %w", errNewEngine))
	}

	err = DBEngine.Ping()
	if err != nil {
		panic(fmt.Errorf("error on ping db: %w", err))
	}

	DBEngine.ShowSQL(true)
	DBEngine.Logger().SetLevel(log.LOG_DEBUG)
	customizedGonicMapper := names.NewPrefixMapper(names.GonicMapper{}, "TBL_")
	DBEngine.SetTableMapper(customizedGonicMapper)
	DBEngine.SetColumnMapper(names.GonicMapper{})

	DBEngine.SetMaxOpenConns(5)
	DBEngine.SetMaxIdleConns(2)
	DBEngine.SetConnMaxLifetime(10 * time.Minute)
}
