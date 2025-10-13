package structure

type One struct {
	Title   string
	Message string
}

type PageData struct {
	Title          string
	Grid           [6][7]int
	PlayerTurn     int
	Color          string
	Visibility     string
	Winner         string
	TextVisibility string
}

type GameData struct {
	Player1 int       `json:player1`
	Player2 int       `json:player2`
	Grid    [6][7]int `json:grid`
	Turn    int       `json:turn`
	IsOver  bool      `json:isover`
}
