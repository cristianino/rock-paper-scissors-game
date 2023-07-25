package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rpsweb/models"
	"rpsweb/usecases"
	"strconv"
)

const (
	templateDir = "templates/"
)

var player models.Player

func Index(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "layout", "index", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restarValues()
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
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := usecases.PlayRound(playerChoice)
	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
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

func restarValues() {
	player.Name = ""
	usecases.ComputerScore = 0
	usecases.PlayerScore = 0
}
