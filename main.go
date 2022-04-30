package main

import (
	"fmt"
	"html/template"
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
	http.HandleFunc("/", index)
	//err := http.ListenAndServe(":80", nil)
	err := http.ListenAndServeTLS(":443", "/usr/local/etc/letsencrypt/live/www.jakobpeters.com/fullchain.pem", "/usr/local/etc/letsencrypt/live/www.jakobpeters.com/privkey.pem", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	err = Tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}
