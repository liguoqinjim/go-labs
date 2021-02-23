package main

import "database/sql"

type TStudent struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;column:id;type:INT4;" json:"id"`
	//[ 1] sid                                            INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Sid int32 `gorm:"column:sid;type:INT4;" json:"sid"`
	//[ 2] sname                                          VARCHAR(10)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
	Sname string `gorm:"column:sname;type:VARCHAR;size:10;" json:"sname"`
	//[ 3] saddress                                       VARCHAR(256)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 256     default: []
	Saddress sql.NullString `gorm:"column:saddress;type:VARCHAR;size:256;" json:"saddress"`
}

// TableName sets the insert table name for this struct type
func (t *TStudent) TableName() string {
	return "t_student"
}
