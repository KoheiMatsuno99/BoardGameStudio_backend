package server

import (
	"context"

	geisterpb "github.com/KoheiMatsuno99/BoardGameStudio_gRPC/pkg/geister/proto"
)

func (gss *GeisterServiceServer) DeleteGameStateWhenGameIsOver(ctx context.Context, req *geisterpb.DeleteGameStateWhenGameIsOverRequest) (*geisterpb.DeleteGameStateWhenGameIsOverResponse, error) {
	tableUuid := req.GetTableUuid()
	delete(gss.gameStateMap, tableUuid)
	return &geisterpb.DeleteGameStateWhenGameIsOverResponse{}, nil
}
