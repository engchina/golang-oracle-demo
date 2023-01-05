package db

import (
	"fmt"
	_ "github.com/godror/godror"
	"github.com/spf13/viper"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

type OracleClientEngine struct {
	*xorm.Engine
}

var (
	XormEngine   *xorm.Engine
	errNewEngine error
	OracleClient OracleClientEngine
)

func InitConfig() {
	viper.SetConfigName("application")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error in read config file %w", err))
	}
}

func InitXormEngine() {
	driverName := viper.GetString("oracle.driverName")
	dataSourceName := viper.GetString("oracle.dataSourceName")
	XormEngine, errNewEngine = xorm.NewEngine(driverName, dataSourceName)
	if errNewEngine != nil {
		panic(fmt.Errorf("error in init new engine %w", errNewEngine))
	}

	errPing := XormEngine.Ping()
	if errPing != nil {
		panic(fmt.Errorf("error on ping db: %w", errPing))
	}

	XormEngine.ShowSQL(false)
	//XormEngine.Logger().SetLevel(log.LOG_DEBUG)
	XormEngine.Logger().SetLevel(log.LOG_INFO)
	XormEngine.SetTableMapper(names.GonicMapper{})
	XormEngine.SetColumnMapper(names.GonicMapper{})

	//XormEngine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	XormEngine.TZLocation, _ = time.LoadLocation("Asia/Tokyo")
	XormEngine.SetMaxOpenConns(5)
	XormEngine.SetMaxIdleConns(2)
	XormEngine.SetConnMaxLifetime(10 * time.Minute)

	// create table
	//err := MyUserXormEngine.Sync(new(models.MyUser))
	//if err != nil {
	//	panic(err)
	//}
}

func InitOracleClient() {
	OracleClient.Engine = XormEngine
}

func init() {
	InitConfig()
	InitXormEngine()
	InitOracleClient()
}

func (engine *OracleClientEngine) ReadWriteTransaction(f func(*xorm.Session, interface{}) (interface{}, error), in interface{}) (interface{}, error) {
	session := engine.NewSession()
	defer func(session *xorm.Session) {
		err := session.Close()
		if err != nil {
			return
		}
	}(session)

	if err := session.Begin(); err != nil {
		return nil, err
	}

	result, err := f(session, in)
	if err != nil {
		return result, err
	}

	if err := session.Commit(); err != nil {
		return result, err
	}

	return result, nil
}

func (engine *OracleClientEngine) ReadOnlyTransaction(f func(*xorm.Session) (interface{}, error)) (interface{}, error) {
	session := engine.NewSession()
	defer func(session *xorm.Session) {
		err := session.Close()
		if err != nil {
			return
		}
	}(session)

	if err := session.Begin(); err != nil {
		return nil, err
	}

	result, err := f(session)
	if err != nil {
		return result, err
	}

	if err := session.Rollback(); err != nil {
		return result, err
	}

	return result, nil
}
