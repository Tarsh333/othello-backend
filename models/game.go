package models

type Position struct {
	Row int `json:"row"`
	Col int `json:"col"`
}
type Game struct {
	ID int        `json:"id"`
	P1 *Player    `json:"p1"`
	P2 *Player    `json:"p2"`
	Bs BoardState `json:"bs"`
}
type AddGameReq struct {
	P1 *Player `json:"p1"`
}
type JoinGameReq struct {
	P2 *Player `json:"p2"`
}
type MoveReq struct {
	PlayerTurn Cell       `json:"pt"`
	Move       []Position `json:"move"`
}
