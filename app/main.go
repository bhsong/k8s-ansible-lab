package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("D_NAME"))

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postres", connStr)
		if err = nil {
			err = db.Ping()
		}
		if err = nil {
			fmt.Println("DB connected!")
			break
		}
		fmt.Println("Waiting for DB...", err)
		time.Sleep(2 * time.Second)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS visits (id SERIAL PRIMARY KEY, visited_at TIMESTAMP)")
	if err != nil {
		log.Fatal("Table creation failed:", err)
	}

	hostname, _ := os.Hostname()
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		db.Exec("INSERT INTO visits (visited_at) VALUES ($1)", time.Now())

		var count int
		db.QueryRow("SELECT COUNT(*) FROM visits").Scan(&count)

		fmt.Fprintf(w, "[%s] Hello K8s! Total Visitors: %d\n", hostname, count)
	})

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}