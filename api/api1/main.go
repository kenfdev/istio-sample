package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Version struct {
	Version string `json:"version"`
}

func main() {
	serverAddr := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	incomingHeaders := []string{
		"x-request-id",
		"x-b3-traceid",
		"x-b3-spanid",
		"x-b3-parentspanid",
		"x-b3-sampled",
		"x-b3-flags",
		"x-ot-span-context",
	}

	http.HandleFunc("/api/db-version", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

		db, err := sql.Open("mysql", connStr)
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		rows, err := db.Query("SELECT VERSION()") //
		if err != nil {
			panic(err.Error())
		}

		var version = &Version{}
		for rows.Next() {
			var v string
			err = rows.Scan(&v)
			if err != nil {
				panic(err.Error())
			}
			version.Version = fmt.Sprintf("MySQL %s", v)
		}

		// pass jaeger tracing headers
		for _, h := range incomingHeaders {
			w.Header().Set(h, r.Header.Get(h))
		}

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(version); err != nil {
			panic(err)
		}
	})

	log.Printf("Listening on %s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
