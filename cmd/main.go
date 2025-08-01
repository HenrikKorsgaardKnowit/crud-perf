package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/henrikkorsgaard/crud-perf/database"
	"github.com/henrikkorsgaard/crud-perf/server"
)

func main() {
	db := database.New("studiepraktik.db")
	port := "3000"
	fmt.Printf("CRUD Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, server.New(db)))
}
