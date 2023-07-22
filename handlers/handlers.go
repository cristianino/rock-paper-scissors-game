package handlers

import (
	"html/template"
	"log"
	"net/http"
)

const (
	templateDir = "templates/"
)

type Player struct {
	Name string
}

var player Player

func Index(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "layout", "index", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "layout", "new-game", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusFound)
			return
		}
		player.Name = r.Form.Get("name")
		renderTemplate(w, "layout", "game", player)
		return
	}
	http.Redirect(w, r, "/new", http.StatusPermanentRedirect)
}

func Play(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "layout", "index", nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "layout", "about", nil)
}

func renderTemplate(w http.ResponseWriter, layout string, nameFile string, data any) {
	tpl := template.Must(template.ParseFiles(templateDir+layout+".html", templateDir+nameFile+".html"))

	err := tpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
