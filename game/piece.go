package game

type Piece struct {
	owner     string
	pieceType string
	position  []int
}

func NewPiece(ownerUuid string, pieceType string, position []int) *Piece {
	return &Piece{
		owner:     ownerUuid,
		pieceType: pieceType,
		position:  position,
	}
}

func (p *Piece) Owner() string {
	return p.owner
}

func (p *Piece) PieceType() string {
	return p.pieceType
}

func (p *Piece) Position() []int {
	return p.position
}

func (p *Piece) SetPosition(position []int) {
	p.position = position
}
