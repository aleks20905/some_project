package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go app...")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./index.html"))

		tmpl.Execute(w, nil)
	}

	// handler function #2 - returns the template block with the newly added film, as an HTMX response
	// h2 := func(w http.ResponseWriter, r *http.Request) {
	// 	time.Sleep(1 * time.Second)

	// 	tmpl := template.Must(template.ParseFiles("index.html"))
	// 	tmpl.ExecuteTemplate(w, "base", )
	// }

	// define handlers
	http.HandleFunc("/", h1)

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./css")))) //from where to be accest in the browser, accest(repeat), whats the dir for the css file
	//http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
