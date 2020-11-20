package service

import (
	"html/template"
	"net/http"

	"github.com/unrolled/render"
)

func loginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		template := template.Must(template.New("login.html").ParseFiles("./templates/login.html"))
		_ = template.Execute(w, struct {
			Content string
		}{Content: "欢迎登录 CloudGo!"})
	}
}
