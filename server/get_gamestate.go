package server

import (
	"context"

	geisterpb "github.com/KoheiMatsuno99/BoardGameStudio_gRPC/pkg/geister/server"
)

func (gss *GeisterServiceServer) GetGameState(ctx context.Context, req *geisterpb.GetGameStateRequest) (*geisterpb.GetGameStateResponse, error) {
	gameState := gss.gameStateMap[req.GetTableUuid()]
	return &geisterpb.GetGameStateResponse{
		GameState: &geisterpb.Table{
			TableUuid: gameState.TableUuid(),
			Players:   gss.serializePlayers(gameState.Players()),
			Board:     gss.serializeBlockRows(gameState.Board()),
			Turn:      uint32(gameState.Turn()),
		},
	}, nil
}
