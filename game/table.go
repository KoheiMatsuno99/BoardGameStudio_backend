package game

import (
	"geister/common"
)

type Table struct{
	tableUuid string
	players []Player
	board [][]Block
	turn int
}

func NewTable(players []Player, board [][]Block, turn int) *Table{
	return &Table{
		tableUuid: common.NewUuid(),
		players: players,
		board: board,
		turn: turn,
	}
}

func (t *Table) TableUuid() string{
	return t.tableUuid
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
