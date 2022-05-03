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
	//go http.ListenAndServe(":80", http.HandlerFunc(redirect))
	http.HandleFunc("/", index)
	err := http.ListenAndServeTLS(":443", "/usr/local/etc/letsencrypt/live/www.jakobpeters.com/fullchain.pem", "/usr/local/etc/letsencrypt/live/www.jakobpeters.com/privkey.pem", nil)
	if err != nil {
		fmt.Println(err)
	}

}

//func redirect(w http.ResponseWriter, r *http.Request) {
//	target := "https://" + r.Host + r.URL.Path
//	if len(r.URL.RawQuery) > 0 {
//		target += "?" + r.URL.RawQuery
//	}
//	log.Print(target)
//	http.Redirect(w, r, target,
//		http.StatusTemporaryRedirect)
//}
func index(w http.ResponseWriter, r *http.Request) {
	err = Tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}
