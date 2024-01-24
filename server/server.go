package server

import (
	usecase "geister/usecase"

	geisterpb "github.com/KoheiMatsuno99/BoardGameStudio_gRPC/pkg/geister/server"
)

type GeisterServiceServer struct {
	geisterpb.UnimplementedGeisterServiceServer
	gameStateMap map[string]*usecase.Table
}

func NewGeisterServiceServer() *GeisterServiceServer {
	return &GeisterServiceServer{
		gameStateMap: make(map[string]*usecase.Table),
	}
}

func (gss *GeisterServiceServer) serializeBlockRows(gameBlocks [][]*usecase.Block) []*geisterpb.BlockRow {
	protoBlockRows := make([]*geisterpb.BlockRow, len(gameBlocks))
	for i, gameBlockRow := range gameBlocks {
		protoBlockRow := &geisterpb.BlockRow{
			Blocks: gss.serializeBlocks(gameBlockRow),
		}
		protoBlockRows[i] = protoBlockRow
	}
	return protoBlockRows
}

func (gss *GeisterServiceServer) serializeBlocks(gameBlocks []*usecase.Block) []*geisterpb.Block {
	protoBlocks := make([]*geisterpb.Block, len(gameBlocks))
	for i, gameBlock := range gameBlocks {
		protoBlock := &geisterpb.Block{
			Address: []uint32{uint32(gameBlock.Address()[0]), uint32(gameBlock.Address()[1])},
		}
		piece := gameBlock.Piece()
		if piece != nil {
			protoBlock.Piece = &geisterpb.Piece{
				Owner:     piece.Owner(),
				PieceType: piece.PieceType(),
				Position:  []uint32{uint32(piece.Position()[0]), uint32(piece.Position()[1])},
			}
		}
		protoBlocks[i] = protoBlock
	}
	return protoBlocks
}

func (gss *GeisterServiceServer) serializePieces(gamePieces map[string]*usecase.Piece) map[string]*geisterpb.Piece {
	protoPieces := make(map[string]*geisterpb.Piece)
	for key, gamePiece := range gamePieces {
		protoPieces[key] = &geisterpb.Piece{
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

func (gss *GeisterServiceServer) serializePlayers(gamePlayers []usecase.Player) []*geisterpb.Player {
	protoPlayers := make([]*geisterpb.Player, len(gamePlayers))
	for i, gamePlayer := range gamePlayers {
		protoPlayers[i] = &geisterpb.Player{
			PlayerUuid:            gamePlayer.PlayerUuid(),
			Name:                  gamePlayer.Name(),
			Pieces:                gss.serializePieces(gamePlayer.Pieces()),
			PickedRedPiecesCount:  uint32(gamePlayer.PickedRedPiecesCount()),
			PickedBluePiecesCount: uint32(gamePlayer.PickedBluePiecesCount()),
		}
	}
	return protoPlayers
}
