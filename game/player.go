package game

import (
	"fmt"
	"geister/common"
)

type Player struct {
	playerUuid            string
	name                  string
	pieces                map[string]*Piece
	pickedRedPiecesCount  int
	pickedBluePiecesCount int
}

func NewPlayer(name string) *Player {
	playerUuid := common.NewUuid()
	pieces := make(map[string]*Piece, 8)
	for i := 0; i < 4; i++ {
		pieces[fmt.Sprintf("%s_blue_%d", playerUuid, i)] = NewPiece(name, "blue", nil)
		pieces[fmt.Sprintf("%s_red_%d", playerUuid, i)] = NewPiece(name, "red", nil)
	}
	return &Player{
		playerUuid: playerUuid,
		name:       name,
		pieces:     pieces,
	}
}

func (p *Player) PlayerUuid() string {
	return p.playerUuid
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Pieces() map[string]*Piece {
	return p.pieces
}

func (p *Player) PickedRedPiecesCount() int {
	return p.pickedRedPiecesCount
}

func (p *Player) PickedBluePiecesCount() int {
	return p.pickedBluePiecesCount
}

func (p *Player) AddPickedRedPiecesCount() {
	p.pickedRedPiecesCount++
}

func (p *Player) AddPickedBluePiecesCount() {
	p.pickedBluePiecesCount++
}
