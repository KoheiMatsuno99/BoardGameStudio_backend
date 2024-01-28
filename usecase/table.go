package usecase

import (
	"geister/common"
)

type Table struct {
	tableUuid string
	players   []Player
	board     [][]*Block
	turn      int
}

func NewTable(players []Player) *Table {
	board := make([][]*Block, 8)
	for x := 0; x < 8; x++ {
		board[x] = make([]*Block, 8)
		for y := 0; y < 8; y++ {
			board[x][y] = &Block{address: []int{x, y}, piece: nil}
		}
	}
	return &Table{
		tableUuid: common.NewUuid(),
		players:   players,
		board:     board,
	}
}

func (t *Table) UpdateTableWhenGamePreparationCompleted(tableUuid string, players []Player, board [][]*Block) *Table{
	t.tableUuid = tableUuid
	t.players = players
	t.board = board
	return t
}

func (t *Table) TableUuid() string {
	return t.tableUuid
}

func (t *Table) Players() []Player {
	return t.players
}

func (t *Table) Board() [][]*Block {
	return t.board
}

func (t *Table) Turn() int {
	return t.turn
}
