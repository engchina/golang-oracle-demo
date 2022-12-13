package utils

import (
	"fmt"
	"time"
	"xorm.io/xorm"
)

// for connect general Database
const dataSourceName = "oracle://pdbadmin:oracle@192.168.31.23:1521/pdb1"

// for connect ADB
//const dataSourceName = `user="admin" password="ToDo" connectString="tcps://adb.ap-singapore-1.oraclecloud.com:1522/g56e4c08bfdf01c_myatp_low.adb.oraclecloud.com?wallet_location=/u01/wallets/Wallet_myatp/"`

var (
	DBEngine     *xorm.Engine
	errNewEngine error
)

func init() {
	DBEngine, errNewEngine = xorm.NewEngine("godror", dataSourceName)
	if errNewEngine != nil {
		panic(fmt.Errorf("error in init new engine %w", errNewEngine))
	}

	err := DBEngine.Ping()
	if err != nil {
		panic(fmt.Errorf("error on ping db: %w", err))
	}

	DBEngine.SetMaxOpenConns(5)
	DBEngine.SetMaxIdleConns(2)
	DBEngine.SetConnMaxLifetime(10 * time.Minute)
}
