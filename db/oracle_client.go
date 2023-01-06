package db

import (
	_ "github.com/godror/godror"
	"github.com/spf13/viper"
	"log"
	"time"
	"xorm.io/xorm"
	xormlog "xorm.io/xorm/log"
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
		log.Fatalln("error in read config file", err)
	}
}

func InitXormEngine() {
	driverName := viper.GetString("oracle.driverName")
	dataSourceName := viper.GetString("oracle.dataSourceName")
	XormEngine, errNewEngine = xorm.NewEngine(driverName, dataSourceName)
	if errNewEngine != nil {
		log.Fatalln("error in init database connection", errNewEngine)
	}

	//errPing := XormEngine.Ping()
	//if errPing != nil {
	//	log.Fatalln("error on ping db", errPing)
	//}

	XormEngine.ShowSQL(false)
	//XormEngine.Logger().SetLevel(xormlog.LOG_DEBUG)
	XormEngine.Logger().SetLevel(xormlog.LOG_INFO)
	XormEngine.SetTableMapper(names.GonicMapper{})
	XormEngine.SetColumnMapper(names.GonicMapper{})

	//XormEngine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	XormEngine.TZLocation, _ = time.LoadLocation("Asia/Tokyo")
	XormEngine.SetMaxOpenConns(5)
	XormEngine.SetMaxIdleConns(2)
	XormEngine.SetConnMaxLifetime(10 * time.Minute)

	//create table
	//errCreateTable := XormEngine.Sync(new(models.MyUser))
	//if errCreateTable != nil {
	//	log.Fatalln("error in create table", errCreateTable)
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
	defer session.Close()
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := session.Begin(); err != nil {
		return nil, err
	}
	result, err := f(session, in)
	if err != nil {
		session.Rollback()
		return nil, err
	}

	if err := session.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}

func (engine *OracleClientEngine) ReadOnlyTransaction(f func(*xorm.Session) (interface{}, error)) (interface{}, error) {
	session := engine.NewSession()
	defer session.Close()
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := session.Begin(); err != nil {
		return nil, err
	}
	result, err := f(session)
	if err != nil {
		session.Rollback()
		return nil, err
	}

	if err := session.Rollback(); err != nil {
		return nil, err
	}

	return result, nil
}
