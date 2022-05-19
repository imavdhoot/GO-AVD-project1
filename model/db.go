package model

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/imavdhoot/GO-AVD-project1/constant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// DB :: Create an exported global variable to hold the database connection pool.
var DB *sql.DB
var goDB *gorm.DB

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		constant.DB_username, constant.DB_password, constant.DB_hostname, constant.DB_database)
}

func InitDB() {
	DB, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	defer DB.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := DB.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+constant.DB_database)
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

	DB.Close()
	DB, err = sql.Open("mysql", dsn())
	if err != nil {
		log.Printf("[DB]Error %s when opening DB", err)
		return
	}
	//defer DB.Close()

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(time.Minute * 30)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = DB.PingContext(ctx)
	if err != nil {
		log.Printf("[DB]Errors pinging DB:: %s", err)
		return
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: DB,
	}), &gorm.Config{})

	if err != nil {
		log.Printf("[DB]Error in Gorm DB pinging DB:: %s", err)
		return
	}

	goDB = gormDB
	log.Printf("########### [DB]Connected to %s successfully ###########", constant.DB_database)
}
