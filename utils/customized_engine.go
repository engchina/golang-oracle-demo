package utils

import (
	"xorm.io/xorm"
)

type CustomizedEngine struct {
	*xorm.Engine
}

func (engine *CustomizedEngine) ReadWriteTransaction(f func(*xorm.Session, interface{}) (interface{}, error), myInterface interface{}) (interface{}, error) {
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

	result, err := f(session, myInterface)
	if err != nil {
		return result, err
	}

	if err := session.Commit(); err != nil {
		return result, err
	}

	return result, nil
}

func (engine *CustomizedEngine) ReadOnlyTransaction(f func(*xorm.Session) (interface{}, error)) (interface{}, error) {
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
