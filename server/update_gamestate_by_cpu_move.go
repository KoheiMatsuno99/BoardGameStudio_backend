package server

import (
	"context"

	geisterpb "github.com/KoheiMatsuno99/BoardGameStudio_gRPC/pkg/geister/proto"
)

func (gss *GeisterServiceServer) UpdateGameStateByCpuMove(ctx context.Context, req *geisterpb.UpdateGameStateByCpuMoveRequest) (*geisterpb.UpdateGameStateByCpuMoveResponse, error) {
	tableUuid := req.GetTableUuid()
	gameState := gss.gameStateMap[tableUuid]
	err := gameState.CpuMove()
	if err != nil {
		return nil, err
	}
	return &geisterpb.UpdateGameStateByCpuMoveResponse{}, nil
}
