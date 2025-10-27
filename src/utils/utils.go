package utils

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"power4/grid"
	"power4/structure"
	"time"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func RenderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl := template.Must(template.ParseFiles("template/" + filename))
	tmpl.Execute(w, data)
}

// redirige vers la bonne route
func SetRoute(w http.ResponseWriter, r *http.Request, route string) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	http.Redirect(w, r, "/"+route, http.StatusSeeOther)
}

// vide la grille
func EmptyGrid() {
	for i := range grid.PointerGrid {
		for j := range grid.PointerGrid[i] {
			grid.PointerGrid[i][j] = 0
		}
	}
}

// charge un json et renvoie ses données sous forme de structure
func LoadJSON() structure.One {
	datahome := structure.One{
		Title:    "Puissance 4",
		Message:  "Bienvenue sur le jeu du Puissance 4 !",
		Historic: []structure.Historic{},
	}

	file, err := os.ReadFile("gamehistoric.json")

	if err == nil {
		if err := json.Unmarshal(file, &datahome.Historic); err != nil {
			fmt.Println("Erreur lors du décodage JSON :", err)
		}
	} else if !os.IsNotExist(err) {
		fmt.Println("Erreur de lecture fichier :", err)
	}
	return datahome
}

// renvoie la date
func GetTime() string {
	return time.Now().Format("02/01/2006")
}

// enregistre un json dans un fichier et le cree si inexistant
func SaveJSON(nomFichier string, data interface{}) error { // Convertir les données en JSON
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(nomFichier, bytes, 0644) // Écrire dans le fichier
}

// ajoute une partie à un fichier json
func AddGameToHistoric(player1 string, player2 string, winner int, date string) { //Sauvegarde en JSON les parties
	var historic []structure.Historic
	var strwinner string

	file, err := os.ReadFile("gamehistoric.json")
	if err == nil {
		if err := json.Unmarshal(file, &historic); err != nil {
			fmt.Println("Erreur lors du décodage JSON :", err)
		}
	} else if !os.IsNotExist(err) {
		fmt.Println("Erreur de lecture fichier :", err)
	}

	if winner == 1 {
		strwinner = player1
	} else if winner == 2 {
		strwinner = player2
	}

	newGame := structure.Historic{
		Player1: player1,
		Player2: player2,
		Winner:  strwinner,
		Date:    date,
	}

	fmt.Println(newGame.Player1, newGame.Player2, newGame.Winner, newGame.Date)
	historic = append(historic, newGame)

	SaveJSON("gamehistoric.json", historic)
}

func Save(w http.ResponseWriter, r *http.Request) { //sauvegarde la partie en cours
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	game := structure.GameData{
		Player1: *structure.PointerPlayer1,
		Player2: *structure.PointerPlayer2,
		Grid:    *grid.PointerGrid,
		Turn:    1,
		IsOver:  false,
	}

	SaveJSON("gamesave.json", game)

	RenderTemplate(w, "save.html", LoadJSON())
}
