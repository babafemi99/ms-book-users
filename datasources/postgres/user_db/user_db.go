package user_db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

var (
	DbUrl = os.Getenv("DbUrl")
)

func Init() *pgx.Conn {
	//var err error
	log.Println("Database connecting ..")
	log.Println("database: ", DbUrl)
	Conn, err := pgx.Connect(context.Background(), DbUrl)
	_, ExecErr := Conn.Exec(context.Background(), "create table if not exists user_db (id INT, first_name VARCHAR(50), last_name VARCHAR(50), email VARCHAR(50), data_created DATE, date_updated DATE)")
	if ExecErr != nil {
		fmt.Println("Error creating table")
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//defer Conn.Close(context.Background())
	if err = Conn.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}
	log.Println("Database connected successfully")
	return Conn
}
