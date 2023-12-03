package server

import (
	"context"
	game "geister/game"

	geisterpb "github.com/KoheiMatsuno99/BoardGameStudio_gRPC/pkg/geister/proto"
)

func (gss *GeisterServiceServer) UpdateGameStateByPlayerMove(ctx context.Context, req *geisterpb.UpdateGameStateByPlayerMoveRequest) (*geisterpb.UpdateGameStateByPlayerMoveResponse, error) {
	tableUuid := req.GetTableUuid()
	pieceKey := req.GetPieceKey()
	serializedDest := req.GetDest()
	gameState := gss.gameStateMap[tableUuid]
	//　Deserializeするので、特にPieceの参照に注意
	dest := game.NewBlock(
		[]int{int(serializedDest.GetAddress()[0]), int(serializedDest.GetAddress()[1])},
		game.NewPiece(
			serializedDest.GetPiece().GetOwner(),
			serializedDest.GetPiece().GetPieceType(),
			[]int{int(serializedDest.GetPiece().GetPosition()[0]), int(serializedDest.GetPiece().GetPosition()[1])},
		),
	)
	err := gameState.PlayerMove(pieceKey, *dest)
	if err != nil {
		return nil, err
	}
	return &geisterpb.UpdateGameStateByPlayerMoveResponse{}, nil
}
