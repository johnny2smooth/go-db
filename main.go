package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/johnny2smooth/go/go-pg/database"
)

const port = 4321

func main() {

	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = database.PrintTitles(db)

	if err != nil {
		log.Fatal(err)
	}

	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	})
	fmt.Printf("Listening on Port %v\n", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%v", port), server)
}
