package main

import (
	"fmt"

	"godoros.bitbucket.io/huarongdao/internal/board"
)

func main() {
	b := board.NewBoard().
		SetPiece(board.PieceMaChao, board.Coord{0, 0}).
		SetPiece(board.PieceCaoCao, board.Coord{1, 0}).
		SetPiece(board.PieceHuangZhong, board.Coord{3, 0}).
		SetPiece(board.PieceZhangFei, board.Coord{0, 2}).
		SetPiece(board.PieceGuanYu, board.Coord{1, 4}).
		SetPiece(board.PieceZhaoYun, board.Coord{3, 2}).
		SetPiece(board.PieceSolider0, board.Coord{1, 2}).
		SetPiece(board.PieceSolider1, board.Coord{2, 2}).
		SetPiece(board.PieceSolider2, board.Coord{1, 3}).
		SetPiece(board.PieceSolider3, board.Coord{2, 3}).
		SetPiece(board.PieceEmpty0, board.Coord{0, 4}).
		SetPiece(board.PieceEmpty1, board.Coord{3, 4})

	// b := board.NewBoard().
	// 	SetPiece(board.PieceHuangZhong, board.Coord{0, 0}).
	// 	SetPiece(board.PieceMaChao, board.Coord{1, 0}).
	// 	SetPiece(board.PieceZhaoYun, board.Coord{2, 0}).
	// 	SetPiece(board.PieceSolider0, board.Coord{3, 0}).
	// 	SetPiece(board.PieceSolider1, board.Coord{3, 1}).
	// 	SetPiece(board.PieceCaoCao, board.Coord{0, 2}).
	// 	SetPiece(board.PieceZhangFei, board.Coord{3, 3}).
	// 	SetPiece(board.PieceGuanYu, board.Coord{0, 4}).
	// 	SetPiece(board.PieceSolider2, board.Coord{2, 3}).
	// 	SetPiece(board.PieceSolider3, board.Coord{2, 4}).
	// 	SetPiece(board.PieceEmpty0, board.Coord{2, 2}).
	// 	SetPiece(board.PieceEmpty1, board.Coord{3, 2})
	fmt.Println("board: ", b.String())

	if solved, solution := board.NewBFSSolver().Solve(b); solved {
		fmt.Printf("got a solution:\n%s\n", solution.DumpMoves())
	} else {
		fmt.Printf("failed to find a solution\n")
	}
}
