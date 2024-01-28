package usecase

import (
	"math"
	"math/rand"
)

func (t *Table) IsValidMove(p Piece, dest Block) bool {
	currentPosition := p.Position()
	destAddress := dest.Address()
	// 現在の位置から1マスより離れていたらfalse
	if math.Abs(float64(currentPosition[0]-destAddress[0])) > 1 || math.Abs(float64(currentPosition[1]-destAddress[1])) > 1 {
		return false
	}
	// 移動しないならfalse
	if currentPosition[0] == destAddress[0] && currentPosition[1] == destAddress[1] {
		return false
	}
	// 移動先に自分の駒があったらfalse
	if dest.Piece() != nil && dest.Piece().Owner() == p.Owner() {
		return false
	}
	return true
}

func (t *Table) CpuMove() error {
	p, dest := t.searchPieceAndDest()
	for !t.IsValidMove(p, dest) {
		p, dest = t.searchPieceAndDest()
	}
	err := t.Move(p, dest)
	if err != nil {
		return err
	}
	return nil
}

func (t *Table) searchPieceAndDest() (Piece, Block) {
	// 移動可能な範囲に相手の青駒があったら取る
	for _, p := range t.Players()[1].Pieces() {
		destAddressList := [][]int{
			{p.Position()[0], p.Position()[1] - 1},
			{p.Position()[0], p.Position()[1] + 1},
			{p.Position()[0] - 1, p.Position()[1]},
			{p.Position()[0] + 1, p.Position()[1]},
		}
		for _, destAddress := range destAddressList {
			target := t.Board()[destAddress[0]][destAddress[1]].Piece()
			if target == nil {
				continue
			}
			if target.Owner() == p.Owner() {
				continue
			}
			if target.PieceType() == "blue" {
				return *p, *t.Board()[destAddress[0]][destAddress[1]]
			}
		}
	}
	// 青駒が取れない場合は相手の赤を取らないように移動
	p, dest := t.avoidRedPiece()
	return p, dest
}

func (t *Table) avoidRedPiece() (Piece, Block) {
	for _, p := range t.Players()[1].Pieces() {
		destAddressList := [][]int{
			{p.Position()[0], p.Position()[1] - 1},
			{p.Position()[0], p.Position()[1] + 1},
			{p.Position()[0] - 1, p.Position()[1]},
			{p.Position()[0] + 1, p.Position()[1]},
		}
		for _, destAddress := range destAddressList {
			if t.Board()[destAddress[0]][destAddress[1]].Piece() != nil &&
				t.Board()[destAddress[0]][destAddress[1]].Piece().PieceType() == "red" {
				continue
			}
			return *p, *t.Board()[destAddress[0]][destAddress[1]]
		}
	}
	// どうしても赤を取るしかない場合はランダムに移動
	keys := make([]string, 0, len(t.Players()[1].Pieces()))
	for k := range t.Players()[1].Pieces() {
		keys = append(keys, k)
	}
	r := rand.Intn(len(keys))
	rKey := keys[r]
	p := t.Players()[1].Pieces()[rKey]
	dlist := [][]int{
		{p.Position()[0], p.Position()[1] - 1},
		{p.Position()[0], p.Position()[1] + 1},
		{p.Position()[0] - 1, p.Position()[1]},
		{p.Position()[0] + 1, p.Position()[1]},
	}
	d := dlist[rand.Intn(len(dlist))]
	return *p, *t.Board()[d[0]][d[1]]
}

func (t *Table) InitCpuPiecesPosition() {
	for _, p := range t.Players()[1].Pieces() {
		for {
			x := rand.Intn(8)
			y := rand.Intn(2)
			// [0][0]と[7][0]は脱出マスなので初期配置では置かない
			if x == 0 && y == 0 {
				continue
			}
			if x == 7 && y == 0 {
				continue
			}
			if t.Board()[y][x].Piece() == nil {
				t.Board()[y][x].SetPiece(p)
				p.SetPosition([]int{y, x})
				break
			}
		}
	}
}
