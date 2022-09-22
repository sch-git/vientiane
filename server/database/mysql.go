package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	dsn string
	db *gorm.DB
	sqlDB *sql.DB
)

func init() {
	dsn = "username:password/@tcp(localhost:3306)/vientiane?charset=utf8mb4&parseTime=True&loc=Local"
}

type DB struct{}

func init()  {
	db,_ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB,_ = db.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetMaxIdleConns(1000)
}

func NewDB() ManagerDB {
	return &DB{}
}

type ManagerDB interface {
	GetDB() (*gorm.DB, error)
	Begin() (db *gorm.DB, err error)
}

func (d *DB) Begin() (db *gorm.DB, err error) {
	fun := "GetTX-->"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if nil != err {
		err = fmt.Errorf("%s: get tx err: %v", fun, err)
		return
	}

	tx := db.Begin()
	return tx, err
}

func (d *DB) GetDB() (*gorm.DB, error) {
	return db,nil
}
