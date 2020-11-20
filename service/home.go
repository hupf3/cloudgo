package service

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/unrolled/render"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		randnum := rand.Intn(1000)
		formatter.JSON(w, http.StatusOK, struct {
			Randnum int
		}{Randnum: randnum})
	}
}

func homeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		template := template.Must(template.New("index.html").ParseFiles("./templates/index.html"))
		_ = template.Execute(w, struct {
			ID      string
			Content string
		}{ID: "345", Content: "Hllo from Go!"})
	}
}

func checkform(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := template.HTMLEscapeString(r.Form.Get("username"))
	password := template.HTMLEscapeString(r.Form.Get("password"))
	randnum := strconv.Itoa(rand.Intn(1000))
	t := template.Must(template.New("signin.html").ParseFiles("./templates/signin.html"))
	err := t.Execute(w, struct {
		Username string
		Password string
		Randnum  string
	}{Username: username, Password: password, Randnum: randnum})
	if err != nil {
		panic(err)
	}
}

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not Implemented"))
}
