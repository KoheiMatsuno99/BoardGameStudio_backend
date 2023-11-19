package game

import (
	"geister/common"
)

type Player struct {
	playerUuid string
	name string
	pieces map[string]*Piece
	pickedRedPiecesCount int
	pickedBluePiecesCount int
}

func NewPlayer(name string) *Player{
	return &Player{
		playerUuid: common.NewUuid(),
		name: name,
		pieces: map[string]*Piece{},
	}
}

func (p *Player) Name() string{
	return p.name
}

func (p *Player) Pieces() map[string]*Piece{
	return p.pieces
}

func (p *Player) PickedRedPiecesCount() int{
	return p.pickedRedPiecesCount
}

func (p *Player) PickedBluePiecesCount() int{
	return p.pickedBluePiecesCount
}

func (p *Player) AddPickedRedPiecesCount() {
	p.pickedRedPiecesCount++
}

func (p *Player) AddPickedBluePiecesCount() {
	p.pickedBluePiecesCount++
}
