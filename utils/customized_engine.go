package utils

import (
	"xorm.io/xorm"
)

type CustomizedEngine struct {
	*xorm.Engine
}

func (engine *CustomizedEngine) ReadWriteTransaction(f func(*xorm.Session, interface{}) (interface{}, error), baseModel interface{}) (interface{}, error) {
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

	result, err := f(session, baseModel)
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

	if err := session.Commit(); err != nil {
		return result, err
	}

	return result, nil
}

//func init() {
//	// set context with timeout
//	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
//	//defer cancel()
//	//MyUserDBEngine.Engine.SetDefaultContext(ctx)
//
//	// create table
//	//err := MyUserDBEngine.Sync(new(models.MyUser))
//	//if err != nil {
//	//	panic(err)
//	//}
//}
