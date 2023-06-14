package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "1991932"
	hostname = "127.0.0.1:3306"
	dbname   = "book"
)

func dns(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func ConnectDB() {
	db, err := sql.Open("mysql", dns(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
		return
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		fmt.Printf("Error %s when creating DB\n", err)
		return
	}
	fmt.Printf("Create DB successfully\n")

	no, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error %s when fetching rows", err)
		return
	}
	fmt.Printf("rows affected %d\n", no)

	db.Close()
	db, err = sql.Open("mysql", dns(dbname))
	if err != nil {
		fmt.Printf("Error %s when opening DB", err)
		return
	}
	defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	err = db.PingContext(ctx)
	if err != nil {
		fmt.Printf("Errors %s pinging DB", err)
		return
	}
	fmt.Printf("Connected to DB %s successfully\n", dbname)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books (id int NOT NULL AUTO_INCREMENT PRIMARY KEY, name varchar(255) NOT NULL, desc text)")
	if err != nil {
		fmt.Printf("Error %s when create table", err)
		return
	}
	fmt.Printf("Create table successfully")

}
