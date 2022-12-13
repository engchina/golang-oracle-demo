package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
	"xorm.io/xorm"
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
	dataSourceName := viper.GetString("dataSourceName")
	DBEngine, errNewEngine = xorm.NewEngine("godror", dataSourceName)
	if errNewEngine != nil {
		panic(fmt.Errorf("error in init new engine %w", errNewEngine))
	}

	err = DBEngine.Ping()
	if err != nil {
		panic(fmt.Errorf("error on ping db: %w", err))
	}

	DBEngine.SetMaxOpenConns(5)
	DBEngine.SetMaxIdleConns(2)
	DBEngine.SetConnMaxLifetime(10 * time.Minute)
}
