package server

import (
	"context"
	game "geister/game"

	geisterpb "github.com/KoheiMatsuno99/BoardGameStudio_gRPC/pkg/geister/proto"
)

func (s *GeisterServer) Start(ctx context.Context, req *geisterpb.StartRequest) (*geisterpb.StartResponse, error) {
	player1 := game.NewPlayer(req.GetPlayer1Name())
	player2 := game.NewPlayer(req.GetPlayer2Name())
	gameState := game.NewTable([]game.Player{*player1, *player2})

	players := []*geisterpb.Table_Player{
		{
			PlayerUuid:            player1.PlayerUuid(),
			Name:                  player1.Name(),
			Pieces:                convertToProtoPieces(player1.Pieces()),
			PickedRedPiecesCount:  uint32(player1.PickedRedPiecesCount()),
			PickedBluePiecesCount: uint32(player1.PickedBluePiecesCount()),
		},
		{
			PlayerUuid:            player2.PlayerUuid(),
			Name:                  player2.Name(),
			Pieces:                convertToProtoPieces(player2.Pieces()),
			PickedRedPiecesCount:  uint32(player2.PickedRedPiecesCount()),
			PickedBluePiecesCount: uint32(player2.PickedBluePiecesCount()),
		},
	}

	return &geisterpb.StartResponse{
		GameState: &geisterpb.Table{
			TableUuid: gameState.TableUuid(),
			Players:   players,
			Board:     convertToProtoBlockRows(gameState.Board()),
			Turn:      uint32(gameState.Turn()),
		},
	}, nil
}

func convertToProtoBlockRows(gameBlocks [][]game.Block) []*geisterpb.Table_BlockRow {
    protoBlockRows := make([]*geisterpb.Table_BlockRow, len(gameBlocks))
    for i, gameBlockRow := range gameBlocks {
        protoBlockRow := &geisterpb.Table_BlockRow{
            Blocks: convertToProtoBlocks(gameBlockRow),
        }
        protoBlockRows[i] = protoBlockRow
    }
    return protoBlockRows
}

func convertToProtoBlocks(gameBlocks []game.Block) []*geisterpb.Table_Block {
	protoBlocks := make([]*geisterpb.Table_Block, len(gameBlocks))
	for i, gameBlock := range gameBlocks {
		protoBlock := &geisterpb.Table_Block{
			Address: []uint32{uint32(gameBlock.Address()[0]), uint32(gameBlock.Address()[1])},
		}
		piece := gameBlock.Piece(); if piece != nil {
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

func convertToProtoPieces(gamePieces map[string]*game.Piece) map[string]*geisterpb.Table_Piece {
    protoPieces := make(map[string]*geisterpb.Table_Piece)
    for key, gamePiece := range gamePieces {
        protoPieces[key] = &geisterpb.Table_Piece{
            Owner:     gamePiece.Owner(),
            PieceType: gamePiece.PieceType(),
        }

		position := gamePiece.Position(); if position != nil {
			protoPieces[key].Position = []uint32{uint32(position[0]), uint32(position[1])}
		}
    }
    return protoPieces
}
