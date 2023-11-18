package game

import "errors"

func (t *Table) Pick(dest Block, target Piece) (error){
	currentTurn := t.Turn()
	var oppenentTurn int
	if currentTurn == 0 {
		oppenentTurn = 1
	} else {
		oppenentTurn = 0
	}
	player := t.Players()[currentTurn]
	oppenent := t.Players()[oppenentTurn]

	targetKey, err := t.searchTargetKey(oppenent, target)
	if err != nil {
		return err
	}
	delete(oppenent.Pieces(), targetKey)
	target.SetPosition(nil)
	dest.SetPiece(nil)
	
	if target.PieceType() == "red" {
		player.AddPickedRedPiecesCount()
	} else {
		player.AddPickedBluePiecesCount()
	}

	return nil
}

func (t *Table) searchTargetKey(oppenent Player, target Piece) (string, error) {
	var targetKey string
	for key, piece := range oppenent.Pieces() {
		isSamePosition, err := isSamePosition(target, *piece)
		if err != nil {
			return targetKey, err
		}
		if piece.Owner() == target.Owner() && piece.PieceType() == target.PieceType() && isSamePosition{
			targetKey = key
			return targetKey, nil
		}
	}
	return targetKey, errors.New("target piece is not found")
}

func isSamePosition(p1 Piece, p2 Piece) (bool, error) {
	if p1.Position() == nil || p2.Position() == nil {
		return false, errors.New("position is nil")
	}
	if p1.Position()[0] == p2.Position()[0] && p1.Position()[1] == p2.Position()[1] {
		return true, nil
	}
	return false, nil
}

func (t *Table) Move(p Piece, dest Block) (error) {
	originalPosition := p.Position()
	destAddress := dest.Address()
	if dest.Piece() != nil && dest.Piece().Owner() != p.Owner(){
		err := t.Pick(dest, *dest.Piece())
		if err != nil {
			return err
		}
	}
	p.SetPosition(destAddress)
	dest.SetPiece(&p)
	t.Board()[originalPosition[0]][originalPosition[1]].SetPiece(nil)
	return nil
}
