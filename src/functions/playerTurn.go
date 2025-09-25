package functions

func PlayerTurn(nextPlayer *string, currentPlayer *string) *string {

	tmp := currentPlayer
	*currentPlayer, *nextPlayer = *nextPlayer, *currentPlayer
	return tmp
}
