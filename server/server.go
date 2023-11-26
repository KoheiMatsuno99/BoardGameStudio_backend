package server

import (
	game "geister/game"

	geisterpb "github.com/KoheiMatsuno99/BoardGameStudio_gRPC/pkg/geister/proto"
)

type GeisterServer struct {
	geisterpb.UnimplementedGeisterServiceServer
}

func NewGeisterServer() *GeisterServer {
	return &GeisterServer{}
}

func (gs *GeisterServer) convertToProtoBlockRows(gameBlocks [][]game.Block) []*geisterpb.Table_BlockRow {
	protoBlockRows := make([]*geisterpb.Table_BlockRow, len(gameBlocks))
	for i, gameBlockRow := range gameBlocks {
		protoBlockRow := &geisterpb.Table_BlockRow{
			Blocks: gs.convertToProtoBlocks(gameBlockRow),
		}
		protoBlockRows[i] = protoBlockRow
	}
	return protoBlockRows
}

func (gs *GeisterServer) convertToProtoBlocks(gameBlocks []game.Block) []*geisterpb.Table_Block {
	protoBlocks := make([]*geisterpb.Table_Block, len(gameBlocks))
	for i, gameBlock := range gameBlocks {
		protoBlock := &geisterpb.Table_Block{
			Address: []uint32{uint32(gameBlock.Address()[0]), uint32(gameBlock.Address()[1])},
		}
		piece := gameBlock.Piece()
		if piece != nil {
			protoBlock.Piece = &geisterpb.Table_Piece{
				Owner:     piece.Owner(),
				PieceType: piece.PieceType(),
				Position:  []uint32{uint32(piece.Position()[0]), uint32(piece.Position()[1])},
			}
		}
		protoBlocks[i] = protoBlock
	}
	return protoBlocks
}

func (gs *GeisterServer) convertToProtoPieces(gamePieces map[string]*game.Piece) map[string]*geisterpb.Table_Piece {
	protoPieces := make(map[string]*geisterpb.Table_Piece)
	for key, gamePiece := range gamePieces {
		protoPieces[key] = &geisterpb.Table_Piece{
			Owner:     gamePiece.Owner(),
			PieceType: gamePiece.PieceType(),
		}

		position := gamePiece.Position()
		if position != nil {
			protoPieces[key].Position = []uint32{uint32(position[0]), uint32(position[1])}
		}
	}
	return protoPieces
}
