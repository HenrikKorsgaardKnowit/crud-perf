package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/henrikkorsgaard/crud-perf/db"
	
)

func main() {)

	db := db.New("studiepraktik.db")
	port := "3000"
	fmt.Printf("CRUD Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, server.New(db)))
}
