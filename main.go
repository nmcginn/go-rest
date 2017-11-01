package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var host = flag.String("host", "127.0.0.1", "the postgres host to connect to")
var port = flag.String("port", "5432", "the postgres port to connect to")
var username = flag.String("username", "postgres", "the username to connect with")
var password = flag.String("password", "postgres", "the password to connect with")
var database = flag.String("database", "postgres", "the database to connect with")
var schema = flag.String("schema", "public", "the schema to serve")

func init() {
	flag.Parse()
}

func main() {
	router := mux.NewRouter()
	router.StrictSlash(true)

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", available_tables)
	http.Handle("/", router)

	fmt.Println("Server started on port 8080")
	fmt.Fprintf(os.Stderr, "%v", http.ListenAndServe(":8080", nil))
}

func available_tables(w http.ResponseWriter, r *http.Request) {
	database := postgres_db{
		Host:     *host,
		Port:     *port,
		Username: *username,
		Password: *password,
		Database: *database,
		Schema:   *schema,
	}
	results, err := describe_tables(database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%v", results)
	}
}
