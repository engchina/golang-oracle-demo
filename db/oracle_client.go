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

var (
	DBEngine     *xorm.Engine
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

func InitDBEngine() {
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

	// create table
	//err := MyUserDBEngine.Sync(new(models.MyUser))
	//if err != nil {
	//	panic(err)
	//}

	// set context with timeout
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	//defer cancel()
	//DBEngine.SetDefaultContext(ctx)
}

func InitOracleClient() {
	OracleClient.Engine = DBEngine
}

func init() {
	InitConfig()
	InitDBEngine()
	InitOracleClient()
}
