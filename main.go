package main

import (
	"log"
	"net/http"

	"./handlers"
	"./utils"
)

func main() {
	http.HandleFunc("/", utils.MakeHandler(handlers.IndexHandler))
	http.HandleFunc("/view/", utils.MakeHandler(handlers.ViewHandler))
	http.HandleFunc("/edit/", utils.MakeHandler(handlers.EditHandler))
	http.HandleFunc("/save/", utils.MakeHandler(handlers.SaveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
