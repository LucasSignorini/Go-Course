package main

import (
	"encoding/json"
	"log"
	"myapp8/rps"
	"net/http"
	"strconv"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	//Pasamos la struct a json con una libreria de GO:
	out, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage) //Handler maneja los request para llevarte a la pag especificada
	//El / significa ir a la pagina home de cualquier web
	log.Println("Empezando el web server:")
	http.ListenAndServe(":8080", nil)

}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
