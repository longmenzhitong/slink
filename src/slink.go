package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"slink/src/rds"
	"slink/src/urls"
	"strings"
)

func init() {
	err := rds.InitRedisClient()
	if err != nil {
		panic(err)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", shortenHandler).Methods("POST")
	router.HandleFunc("/{code}", redirectHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func shortenHandler(w http.ResponseWriter, req *http.Request) {
	originUrl := req.FormValue("originUrl")
	shortLink, err := urls.Shorten(originUrl)
	if err != nil {
		_, err = fmt.Fprintf(w, "shorten error: %s", err.Error())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err = fmt.Fprintf(w, "%s", shortLink)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func redirectHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	code := vars["code"]
	url, err := urls.Expand(code)
	if err != nil {
		_, err = fmt.Fprintf(w, "expand error: %s", err.Error())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		http.Redirect(w, req, url, http.StatusFound)
	}
}