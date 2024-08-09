package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gabriel-hawerroth/capitech-back/configs"
	_ "github.com/lib/pq"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", configs.DBUser, configs.DBPassword, configs.DBName, configs.DBHost, configs.DBPort))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	handler := cors(mux)

	mux.HandleFunc("/", helloHandler)

	fmt.Println("Starting web server on port", configs.WebServerPort)
	err = http.ListenAndServe(configs.WebServerPort, handler)
	if err != nil {
		panic(err)
	}
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
