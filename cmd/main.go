package cmd

import (
	"database/sql"
	"flag"
	"log"
)

func main() {
	_ = flag.String("addr", ":4000", "HTTP network address")
	_ = flag.Bool("debug", false, "Debug mode")

	dsn := flag.String("dsn", "admin:epicepic8@/snippetbox?parseTime=true", "My Postgres")

	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
