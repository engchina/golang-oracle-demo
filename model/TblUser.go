package model

type TblUser struct {
	Id         int64  `json:"userId" xorm:"pk 'USER_ID'"`
	Name       string `json:"name" xorm:"varchar(200) notnull 'NAME''"`
	NumOfTried int64  `json:"numOfTried" xorm:"notnull 'NUM_OF_TRIED'"`
	Created    string `json:"createdTime" xorm:"notnull 'CREATED_TIME'"`
	Updated    string `json:"updatedTime" xorm:"notnull 'UPDATED_TIME'"`
}

func (TblUser) TableName() string {
	return "TBL_USER"
}
