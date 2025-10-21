package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"power4/grid"
	"power4/structure"
	"time"
)

func SaveJSON(nomFichier string, data interface{}) error {
	// Convertir les données en JSON (avec indentation pour la lisibilité)
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Écrire dans le fichier
	return os.WriteFile(nomFichier, bytes, 0644)
}

// redirige vers la bonne route
func SetRoute(w http.ResponseWriter, r *http.Request, route string) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	http.Redirect(w, r, "/"+route, http.StatusSeeOther)

}

func EmptyGrid() {
	for i := range grid.PointerGrid {
		for j := range grid.PointerGrid[i] {
			grid.PointerGrid[i][j] = 0
		}
	}
}

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

func GetTime() string {
	return time.Now().Format("02/01/2006")
}
