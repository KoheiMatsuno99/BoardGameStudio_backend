package game

type Player struct {
	// TODO: guidを追加する
	name string
	pieces map[string]*Piece
	pickedRedPiecesCount int
	pickedBluePiecesCount int
}

func NewPlayer(name string, pieces map[string]*Piece, pickedRedPiecesCount, pickedBluePiecesCount int) *Player{
	return &Player{
		name: name,
		pieces: map[string]*Piece{},
		pickedRedPiecesCount: pickedRedPiecesCount,
		pickedBluePiecesCount: pickedBluePiecesCount,
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
