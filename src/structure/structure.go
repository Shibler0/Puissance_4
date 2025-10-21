package structure

type One struct {
	Title    string
	Message  string
	Historic []Historic
}

type PageData struct {
	Title          string
	Grid           [6][7]int
	PlayerTurn     string
	Color          string
	Visibility     string
	Winner         string
	TextVisibility string
	Encouragement  string
}

type GameData struct {
	Player1 string    `json:player1`
	Player2 string    `json:player2`
	Grid    [6][7]int `json:grid`
	Turn    int       `json:turn`
	IsOver  bool      `json:isover`
}

type Historic struct {
	Player1 string `json:player1`
	Player2 string `json:player2`
	Winner  string `json:winner`
	Date    string `json:date`
}

var Player1 string
var Player2 string
var PointerPlayer1 = &Player1
var PointerPlayer2 = &Player2
