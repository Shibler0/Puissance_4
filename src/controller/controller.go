package controller

import (
	"net/http"
	"power4/grid"
	"power4/structure"
	"power4/utils"
	"strconv"
)

// Home gère la page d'accueil
func Home(w http.ResponseWriter, r *http.Request) {
	x := structure.PointerPlayer1
	y := structure.PointerPlayer2
	*x = ""
	*y = ""

	utils.EmptyGrid()
	utils.LoadJSON()

	utils.RenderTemplate(w, "home.html", utils.LoadJSON()) // Affiche le template index.html avec les données
}

// gere l'écran de jeu
func Game(w http.ResponseWriter, r *http.Request) { //la fonction principale du jeu puissance 4
	title := "Pion à poser : "
	playerTurn := *(grid.PlayerTurnPointer)
	visibility := "auto"
	winner := "dimitri"
	textvisibility := "none"
	encouragement := "Bonne chance"

	var x *string = structure.PointerPlayer1
	var y *string = structure.PointerPlayer2

	if *x == "" && *y == "" {
		*x = "Joueur1"
		*y = "Joueur2"
	}

	if *(grid.PlayerTurnPointer) == "" {
		*(grid.PlayerTurnPointer) = *x
		*(grid.ColorPointer) = 1
	}

	if r.Method == http.MethodPost {
		formPlayer1 := r.FormValue("Joueur1")
		formPlayer2 := r.FormValue("Joueur2")

		if formPlayer1 != "" {
			*structure.PointerPlayer1 = formPlayer1
		}
		if formPlayer2 != "" {
			*structure.PointerPlayer2 = formPlayer2
		}

		encouragement = grid.RandomEncouragement()
		col := r.FormValue("col")

		if col != "" {
			colInt, _ := strconv.Atoi(col)
			player, iswon := grid.SetToken(colInt)

			if iswon && player == 1 {
				visibility = "none"
				textvisibility = "auto"
				winner = "Félicitations joueur " + *x
				utils.AddGameToHistoric(*x, *y, player, utils.GetTime())
				*x = ""
				*y = ""
			}

			if iswon && player == 2 {
				visibility = "none"
				textvisibility = "auto"
				winner = "Félicitations joueur " + *y
				utils.AddGameToHistoric(*x, *y, player, utils.GetTime())
				*x = ""
				*y = ""
			}

			if iswon && player == 0 {
				visibility = "none"
				textvisibility = "auto"
				winner = "Égalité"
				utils.AddGameToHistoric(*x, *y, 0, utils.GetTime())
			}
		}
	}

	data := structure.PageData{
		Title:          title,
		Grid:           *grid.PointerGrid,
		PlayerTurn:     playerTurn,
		Color:          grid.SetColor(),
		Visibility:     visibility,
		Winner:         winner,
		TextVisibility: textvisibility,
		Encouragement:  encouragement,
	}
	utils.RenderTemplate(w, "play.html", data)
}

func Returnmenu(w http.ResponseWriter, r *http.Request) {
	utils.EmptyGrid()
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func Reset(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	utils.EmptyGrid()

	http.Redirect(w, r, "/play", http.StatusSeeOther)
}

func Replay(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	utils.EmptyGrid()
	*grid.ColorPointer = 1
	*grid.PlayerTurnPointer = "1"

	http.Redirect(w, r, "/play", http.StatusSeeOther)
}
