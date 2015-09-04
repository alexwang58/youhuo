package components

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Templater struct {
}

var funcs = template.FuncMap{}

var ViewDir = "view"
var Layout = "main"

func NewTpler() *Templater {
	return &Templater{}
}

func RenderTemplate(w http.ResponseWriter, status int, name string, data interface{}) error {
	pathInfo := strings.Split(name, "/")
	if len(pathInfo) != 2 {
		return errors.New("render view path should like \"prouct/item\", but (\"" + name + "\") given")
	}

	files := []string{
		ViewDir + "/layout/" + Layout + ".html",
		ViewDir + "/" + name + ".html",
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	t, err := template.New("*").Funcs(funcs).ParseFiles(files...)
	if err != nil {
		fmt.Errorf("[RenderError] " + err.Error())
		return err
	}

	return template.Must(t, err).ExecuteTemplate(w, "main", data)
}
