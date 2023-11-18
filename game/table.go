package game

type Table struct{
	players []Player
	board [][]Block
	turn int
}

func NewTable(players []Player, board [][]Block, turn int) *Table{
	return &Table{
		players: players,
		board: board,
		turn: turn,
	}
}

func (t *Table) Players() []Player{
	return t.players
}

func (t *Table) Board() [][]Block{
	return t.board
}

func (t *Table) Turn() int{
	return t.turn
}
