package structure

type One struct {
	Title    string
	Message  string
	Historic []Partie
}
type Partie struct {
	Date    string
	Joueur1 string
	Joueur2 string
	Status  string
}
