package game

type Piece struct {
	owner string // Todo Player型に直す
	pieceType string
	position []int
}

func NewPiece(owner string, pieceType string, position []int) *Piece{
	return &Piece{
		owner: owner,
		pieceType: pieceType,
		position: position,
	}
}

func (p *Piece) Owner() string {
	return p.owner
}

func (p *Piece) PieceType() string{
	return p.pieceType
}

func (p *Piece) Position() []int{
	return p.position
}
