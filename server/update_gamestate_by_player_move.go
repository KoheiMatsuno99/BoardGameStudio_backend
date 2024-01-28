package server

import (
	"context"
	usecase "geister/usecase"

	geisterpb "github.com/KoheiMatsuno99/BoardGameStudio_gRPC/pkg/geister/server"
)

func (gss *GeisterServiceServer) UpdateGameStateByPlayerMove(ctx context.Context, req *geisterpb.UpdateGameStateByPlayerMovementRequest) (*geisterpb.UpdateGameStateByPlayerMovementResponse, error) {
	tableUuid := req.GetTableUuid()
	pieceKey := req.GetPieceKey()
	serializedDest := req.GetDest()
	gameState := gss.gameStateMap[tableUuid]

	// Deserialize
	var piece *usecase.Piece
	serializedPiece := serializedDest.GetPiece()
	if serializedPiece != nil {
		piece = usecase.NewPiece(
			serializedDest.GetPiece().GetOwner(),
			serializedDest.GetPiece().GetPieceType(),
			[]int{int(serializedDest.GetPiece().GetPosition()[0]), int(serializedDest.GetPiece().GetPosition()[1])},
		)
	}
	dest := usecase.NewBlock(
		[]int{int(serializedDest.GetAddress()[0]), int(serializedDest.GetAddress()[1])},
		piece,
	)
	// Deserializeで生成したPiece, Blockで上書きする
	gameState.Players()[gameState.Turn()].Pieces()[pieceKey] = piece
	gameState.Board()[dest.Address()[0]][dest.Address()[1]] = dest

	err := gameState.PlayerMove(pieceKey, *dest)
	if err != nil {
		return nil, err
	}
	return &geisterpb.UpdateGameStateByPlayerMovementResponse{}, nil
}
