package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

var url string = fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/493")
var CardTemplate *template.Template

type Card struct {
	Name      string      `json:"name"`
	Height    float64     `json:"height"`
	Weight    float64     `json:"weight"`
	Abilities []Abilities `json:"abilities"`
}

type Abilities struct {
	Ability Ability `json:"ability"`
}
type Ability struct {
	Name string `json:"name"`
}

type PassingData struct {
	Name        string
	Height      float64
	Weight      float64
	AbilityName string
}

func ShowCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	cardData := getCardData(url)
	var abilityName string
	for _, w := range cardData.Abilities {
		abilityName = w.Ability.Name
	}
	data := &PassingData{
		Name:        cardData.Name,
		Height:      cardData.Height,
		Weight:      cardData.Weight,
		AbilityName: abilityName,
	}
	err := CardTemplate.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func getCardData(url string) *Card {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var card Card
	err = json.Unmarshal(body, &card)
	if err != nil {
		log.Fatal(err)
	}
	return &card

}

func main() {

	var err error

	//parse html file
	CardTemplate, err = template.ParseFiles("main.html")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", ShowCard)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
