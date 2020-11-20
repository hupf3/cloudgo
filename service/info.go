package service

import (
	"html/template"
	"net/http"
)

func infoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := template.HTMLEscapeString(r.Form.Get("username"))
	password := template.HTMLEscapeString(r.Form.Get("password"))
	t := template.Must(template.New("info.html").ParseFiles("./templates/info.html"))
	err := t.Execute(w, struct {
		Username string
		Password string
		ID       string
		School   string
		GPA      float32
	}{Username: username, Password: password, ID: "18342025", School: "中山大学", GPA: 4.2})
	if err != nil {
		panic(err)
	}
}
