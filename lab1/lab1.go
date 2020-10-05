package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path"
	"time"
)

//Time ...
type Time struct {
	Date    string `json:"date"`
	Hour    int    `json:"hour"`
	Minutes int    `json:"min"`
	Second  int    `json:"sec"`
}

func main() {
	http.HandleFunc("/", mainPage)

	http.HandleFunc("/online", online)

	http.HandleFunc("/time", timePage)
	http.ListenAndServe(":8795", nil)

}

func timePage(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<a href=" + "/" + "><button><h1>/</h1></button></a>"))
	w.Write([]byte("<a href=" + "/online" + "><button><h1>/online</h1></button></a>"))

	p := time.Now()

	time := Time{p.Format("1 January Monday"), p.Hour(), p.Minute(), p.Second()}

	JSON, _ := json.Marshal(time)
	w.Write(JSON)

}

func mainPage(w http.ResponseWriter, r *http.Request) {

	fp := path.Join("index1.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func online(w http.ResponseWriter, r *http.Request) {
	p := time.Now()

	time := Time{p.Format("2006-01-02T15:04:05Z07:00"), p.Hour(), p.Minute(), p.Second()}

	fp := path.Join("index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, time); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
