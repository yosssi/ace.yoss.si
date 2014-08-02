package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/yosssi/ace"
)

func topIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl, err := ace.Load("base", "top/index", &ace.Options{BaseDir: "views", Asset: Asset})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", topIndex)
	router.ServeFiles("/public/*filepath", http.Dir("public"))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
