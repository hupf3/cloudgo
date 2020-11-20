package service

import (
	"html/template"
	"net/http"

	"github.com/unrolled/render"
)

func infoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "info", struct {
			School string `json:"school"`
			ID     string `json:"id"`
			Name   string `json:"name"`
			GPA    string `json:"gpa"`
		}{School: "中山大学", ID: "18342025", Name: "胡鹏飞", GPA: "4.5"})
	}
}

func checkform(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := template.HTMLEscapeString(r.Form.Get("username"))
	password := template.HTMLEscapeString(r.Form.Get("password"))
	t := template.Must(template.New("info.html").ParseFiles("./templates/info.html"))
	err := t.Execute(w, struct {
		Username string
		Password string
	}{Username: username, Password: password})
	if err != nil {
		panic(err)
	}
}
