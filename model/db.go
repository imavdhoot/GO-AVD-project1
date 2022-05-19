package model

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/imavdhoot/GO-AVD-project1/constant"
	"log"
	"time"
)

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		constant.DB_username, constant.DB_password, constant.DB_hostname, constant.DB_database)
}

func InitDB() {
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+constant.DB_database)
	if err != nil {
		log.Printf("[DB]Error %s when creating DB\n", err)
		return
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("[DB]Error %s when fetching rows", err)
		return
	}
	log.Printf("[DB]rows affected %d\n", no)

	db.Close()
	db, err = sql.Open("mysql", dsn())
	if err != nil {
		log.Printf("[DB]Error %s when opening DB", err)
		return
	}
	defer db.Close()

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 30)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("[DB]Errors %s pinging DB", err)
		return
	}
	log.Printf("########### [DB]Connected to %s successfully ###########", constant.DB_database)
}
