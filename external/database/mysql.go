package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dickidarmawansaputra/go-simple-online-shop/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectMysql(cfg config.DBConfig) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.ConnectionPool.MaxIdleConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectionPool.MaxLifetimeConnection) * time.Second)
	db.SetMaxOpenConns(int(cfg.ConnectionPool.MaxOpenConnection))
	db.SetMaxIdleConns(int(cfg.ConnectionPool.MaxIdleConnection))

	return
}
