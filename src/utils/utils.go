package utils

import (
	"encoding/json"
	"os"
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
