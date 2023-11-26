package game

type Block struct {
	address []int
	piece   *Piece
}

func NewBlock(address []int, piece *Piece) *Block {
	return &Block{
		address: address,
		piece:   piece,
	}
}

func (b *Block) Address() []int {
	return b.address
}

func (b *Block) Piece() *Piece {
	return b.piece
}

func (b *Block) SetPiece(piece *Piece) {
	b.piece = piece
}
