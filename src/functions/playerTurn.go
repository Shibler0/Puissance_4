package functions

func PlayerTurn(currentPlayer *string, nextPlayer *string) *string {

	tmp := currentPlayer
	*currentPlayer, *nextPlayer = *nextPlayer, *currentPlayer
	return tmp
}
