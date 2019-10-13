// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", connStr)

	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("postgres connected :: %#v \n", connStr)
	}

	r := mux.NewRouter()

	server := &http.Server{
		Addr: ":8080",
	}

	r.HandleFunc("/", hello)
	http.Handle("/", r)

	fmt.Printf("starting server :: %#v \n", server.Addr)

	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))

}

func hello(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "success",
		Message: "Hello world!",
	}

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)

}
