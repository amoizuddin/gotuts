package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//can handle dynamic urls

	router.HandleFunc("/books/{title}/page/{page}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(writer, "you've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", router)

}
