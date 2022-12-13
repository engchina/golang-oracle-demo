package main

import (
	_ "github.com/godror/godror"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("oracle", "oracle://pdbadmin:oracle@192.168.31.23:1521/pdb1")
	if err != nil {
		panic(err)
	}
}
