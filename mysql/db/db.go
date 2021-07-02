package db

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func New(cfg MysqlConfig) *sql.DB {
	mysql, err := sql.Open("mysql", cfg.Dsn())
	if err != nil {
		log.Panicln(err)
	}
	mysql.SetConnMaxLifetime(time.Minute * 3)
	mysql.SetMaxOpenConns(10)
	mysql.SetMaxIdleConns(10)
	return mysql
}

func NewConnection(ctx context.Context, cfg MysqlConfig) *sql.Conn {
	mysql := New(cfg)
	conn, err := mysql.Conn(ctx)
	if err != nil {
		log.Panicln(err)
	}
	return conn
}

type CloserConnection interface {
	Close() error
}

func Close(closer CloserConnection) {
	if closer != nil {
		err := closer.Close()
		if err != nil {
			log.Panicln(err)
		}
	}
}