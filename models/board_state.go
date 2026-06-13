package models

type Cell int

const (
	Empty Cell = iota
	P1
	P2
)

type BoardState struct {
	Board [][]Cell
}

func GetNewBoard() *BoardState {
	return initBoard()
}
func initBoard() *BoardState {
	board := make([][]Cell, 8)

	for i := range board {
		board[i] = make([]Cell, 8)
	}

	board[3][3] = P2
	board[3][4] = P1
	board[4][3] = P1
	board[4][4] = P2

	return &BoardState{
		Board: board,
	}

}
