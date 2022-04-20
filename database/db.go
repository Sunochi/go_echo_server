package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"go_example_server/config"
)

var db *gorm.DB
var err error

func Init() {
	c := config.GetConfig()
	dsnTest := dsn(&c.DB.Test)
	dsnTestRead := dsn(&c.DB.TestRead)

	db, err = gorm.Open(mysql.Open(dsnTest), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	err = db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(dsnTestRead)},
	}))

	if err != nil {
		panic(err.Error())
	}
}

func dsn(d *config.DBSetting) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", d.User, d.Password, d.Host, d.Port, d.DBName)
}

// GetDB returns database connection
func GetDB() *gorm.DB {
	return db
}
