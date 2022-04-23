// Copyright 2022 DomineCore.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package brisk

import (
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConnector struct {
	DriverName string // driver_name: mysql/ sqlite/ postgresql/ sqlserver
}

func NewDBConnector() *DBConnector {
	driverName := Config.GetString("Databases.default.driver")
	return &DBConnector{DriverName: driverName}
}

func (db *DBConnector) Connect() {
	// dispatch driver
	var engine DBEngine
	switch db.DriverName {
	case "mysql":
		engine = new(MySQL)
	case "postgresql":
		engine = new(PostgreSQL)
	// case "sqlite":
	// 	engine = new(SQLite)
	// case "sqlserver":
	// 	engine = new(SQLServer)
	default:
		engine = new(MySQL)
	}
	engine.Init()
	// conn
	db_inst, err := engine.Connect()
	if err != nil {
		panic(err)
	}
	DB = db_inst
}

type DBEngine interface {
	Init()                      // init config
	Connect() (*gorm.DB, error) // connect
}

type MySQL struct {
	Username  string
	Password  string
	Host      string
	Port      string
	DBNAME    string
	Charset   string
	Parsetime string
}

func (engine *MySQL) Init() {
	engine.Username = Config.GetString("Databases.default.username")
	engine.Password = Config.GetString("Databases.default.password")
	engine.Host = Config.GetString("Databases.default.host")
	engine.Port = Config.GetString("Databases.default.port")
	engine.DBNAME = Config.GetString("Databases.default.dbname")
	engine.Charset = Config.GetString("Databases.default.charset")
	engine.Parsetime = Config.GetString("Databases.default.parsetime")
}

func (engine *MySQL) Connect() (*gorm.DB, error) {
	params := url.Values{}
	params.Set("charset", engine.Charset)
	params.Set("parseTime", engine.Parsetime)
	encode_pram := params.Encode()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", engine.Username, engine.Password, engine.Host, engine.Port, engine.DBNAME, encode_pram)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

// type SQLServer struct {
// }

// func (engine *SQLServer) Init() {

// }

// func (engine *SQLServer) Connect(gorm.DB, error) {

// }

// type SQLite struct {
// 	Filepath string
// }

// func (engine *SQLite) Init() {
// 	engine.Filepath = Config.GetString("Databases.default.filepath")
// }

// func (engine *SQLite) Connect() (gorm.DB, error) {
// 	db, err := gorm.Open(sqlite.Open(engine.Filepath), &gorm.Config{})
// 	return db, err
// }

type PostgreSQL struct {
	Host     string
	Username string
	Password string
	DBNAME   string
	Port     string
	Sslmode  string
	TimeZone string
}

func (engine *PostgreSQL) Init() {
	engine.Host = Config.GetString("Databases.default.host")
	engine.Username = Config.GetString("Databases.default.username")
	engine.Password = Config.GetString("Databases.default.password")
	engine.DBNAME = Config.GetString("Databases.default.dbname")
	engine.Port = Config.GetString("Databases.default.port")
	engine.Sslmode = Config.GetString("Databases.default.sslmode")
	engine.TimeZone = Config.GetString("Databases.default.timezone")
}

func (engine *PostgreSQL) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		engine.Host,
		engine.Username,
		engine.Password,
		engine.DBNAME,
		engine.Port,
		engine.Sslmode,
		engine.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
