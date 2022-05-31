package services

import (
	"github.com/jinzhu/gorm"
)

var email string
var password string
var dbHost string
var dbPort string
var dbName string

var Db *gorm.DB

func OpenDatabase() {
	//open a db connection
	var err error
	Db, err = gorm.Open("postgres", "postgres://"+"postgres"+":"+"postgres"+"@"+"projectdb.cexdsznfaerp.eu-west-2.rds.amazonaws.com"+":"+"5432"+"/"+"apidb"+"?sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
}
