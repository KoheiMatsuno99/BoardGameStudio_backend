package server

import (
	"context"
	usecase "geister/usecase"

	geisterpb "github.com/KoheiMatsuno99/BoardGameStudio_gRPC/pkg/geister/server"
)

func (gss *GeisterServiceServer) NotifyGamePreparationCompleted(ctx context.Context, req *geisterpb.NotifyGamePreparationCompletedRequest) (*geisterpb.NotifyGamePreparationCompletedResponse, error) {
	tableUuid := req.GetTableUuid()
	serializedPlayers := req.GetPlayers()
	serializedBoard := req.GetBoard()
	prevTable := gss.gameStateMap[tableUuid]
	prevPlayers := prevTable.Players()

	rows := make([][]*usecase.Block, 8)
	for i, serializedBlockRow := range serializedBoard {
		rows[i] = make([]*usecase.Block, 8)
		serializedBlocks := serializedBlockRow.GetBlocks()
		for j, serializedBlock := range serializedBlocks {
			serializedPiece := serializedBlock.GetPiece()
			var piece *usecase.Piece
			if serializedPiece != nil {
				piece = usecase.NewPiece(
					serializedPiece.GetOwner(),
					serializedPiece.GetPieceType(),
					[]int{int(serializedPiece.GetPosition()[0]), int(serializedPiece.GetPosition()[1])},
				)
			}
			rows[i][j] = usecase.NewBlock(
				[]int{int(serializedBlock.GetAddress()[0]), int(serializedBlock.GetAddress()[1])},
				piece,
			)
		}
	}

	players := make([]usecase.Player, 2)
	for i, serializedPlayer := range serializedPlayers {
		serializedPieces := serializedPlayer.GetPieces()
		pieces := make(map[string]*usecase.Piece, 8)
		for key, serializedPiece := range serializedPieces {
			pieces[key] = usecase.NewPiece(
				serializedPiece.GetOwner(),
				serializedPiece.GetPieceType(),
				[]int{int(serializedPiece.GetPosition()[0]), int(serializedPiece.GetPosition()[1])},
			)
		}
		players[i] = *prevPlayers[i].UpdatePlayerWhenGamePreparationCompleted(serializedPlayer.PlayerUuid,
			serializedPlayer.Name,
			pieces,
		)
	}

	prevTable.UpdateTableWhenGamePreparationCompleted(tableUuid, players, rows)
	return &geisterpb.NotifyGamePreparationCompletedResponse{}, nil
}
