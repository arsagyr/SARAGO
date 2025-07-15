package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type actor struct {
	id     int
	title  string
	artist string
	price  float32
}

func main() {
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	defer db.Close()
	db.Query("USE `recordings`;")
	rows, err := db.Query("select * from `album`")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []actor{}

	for rows.Next() {
		p := actor{}
		err := rows.Scan(&p.id, &p.title, &p.artist, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Println(p.id, p.title, p.artist, p.price)
	}

}
