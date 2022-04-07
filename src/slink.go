package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"slink/src/conf"
	"slink/src/rds"
	"slink/src/urls"
	"strings"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", shortenHandler).Methods(http.MethodPost)
	router.HandleFunc("/{code}", redirectHandler).Methods(http.MethodGet)
	router.HandleFunc("/pv", pvHandler).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":"+conf.Port, router))
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

		// 记录PV
		pvKey := conf.PvKeyPrefix + code
		rds.Client.Incr(pvKey)

		http.Redirect(w, req, url, http.StatusFound)
	}
}

func pvHandler(w http.ResponseWriter, req *http.Request) {
	shortLink := req.FormValue("shortLink")
	parts := strings.Split(shortLink, "/")
	code := parts[len(parts)-1]
	cmd := rds.Client.Get(conf.PvKeyPrefix + code)
	pv, err := cmd.Result()
	if err != nil {
		_, err = fmt.Fprintf(w, "get pv error: %s", err.Error())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err = fmt.Fprintf(w, "%s", pv)
		if err != nil {
			log.Fatal(err)
		}
	}

}
