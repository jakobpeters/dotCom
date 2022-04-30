package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var Tpl *template.Template
var err error

func init() {
	Tpl = template.Must(template.ParseGlob("*.html"))
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))
	http.HandleFunc("/", index)
	err := http.ListenAndServeTLS(":443", "/usr/local/etc/letsencrypt/live/www.jakobpeters.com/fullchain.pem", "/usr/local/etc/letsencrypt/live/www.jakobpeters.com/privkey.pem", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		// see comments below and consider the codes 308, 302, or 301
		http.StatusTemporaryRedirect)
}
func index(w http.ResponseWriter, r *http.Request) {
	err = Tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}
