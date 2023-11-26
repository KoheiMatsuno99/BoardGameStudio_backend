package server

import (
	"context"
	game "geister/game"

	geisterpb "github.com/KoheiMatsuno99/BoardGameStudio_gRPC/pkg/geister/proto"
)

func (gss *GeisterServiceServer) Start(ctx context.Context, req *geisterpb.StartRequest) (*geisterpb.StartResponse, error) {
	player1 := game.NewPlayer(req.GetPlayer1Name())
	player2 := game.NewPlayer(req.GetPlayer2Name())
	gameState := game.NewTable([]game.Player{*player1, *player2})

	players := []*geisterpb.Table_Player{
		{
			PlayerUuid:            player1.PlayerUuid(),
			Name:                  player1.Name(),
			Pieces:                gss.convertToProtoPieces(player1.Pieces()),
			PickedRedPiecesCount:  uint32(player1.PickedRedPiecesCount()),
			PickedBluePiecesCount: uint32(player1.PickedBluePiecesCount()),
		},
		{
			PlayerUuid:            player2.PlayerUuid(),
			Name:                  player2.Name(),
			Pieces:                gss.convertToProtoPieces(player2.Pieces()),
			PickedRedPiecesCount:  uint32(player2.PickedRedPiecesCount()),
			PickedBluePiecesCount: uint32(player2.PickedBluePiecesCount()),
		},
	}

	gss.gameStateMap[gameState.TableUuid()] = gameState

	return &geisterpb.StartResponse{
		GameState: &geisterpb.Table{
			TableUuid: gameState.TableUuid(),
			Players:   players,
			Board:     gss.convertToProtoBlockRows(gameState.Board()),
			Turn:      uint32(gameState.Turn()),
		},
	}, nil
}
