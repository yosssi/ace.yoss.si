package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/yosssi/ace"
)

var startupTime = time.Now()

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

func serveAsset(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	url := r.URL.Path

	b, err := Asset(url[1:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, r, url, startupTime, bytes.NewReader(b))
}

func main() {
	router := httprouter.New()

	router.GET("/", topIndex)
	router.GET("/public/*filepath", serveAsset)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
