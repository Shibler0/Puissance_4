package utils

import (
	"encoding/json"
	"net/http"
	"os"
	"power4/grid"
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
