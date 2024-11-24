package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shiweiwt/huarongdao/internal/board"
)

var logVerbose = flag.Bool("logverbose", false, "whether to enable verbose logging")
var boardID = flag.Int("board", 0, "which board to solve, [0|1|2], default 0")

func getBoard(id int) (*board.Board, error) {
	switch id {
	case 0:
		b := board.NewBoard().
			SetPiece(board.PieceMaChao, board.Coord{0, 0}).
			SetPiece(board.PieceCaoCao, board.Coord{1, 0}).
			SetPiece(board.PieceHuangZhong, board.Coord{3, 0}).
			SetPiece(board.PieceZhangFei, board.Coord{0, 2}).
			SetPiece(board.PieceGuanYu, board.Coord{1, 2}).
			SetPiece(board.PieceZhaoYun, board.Coord{3, 2}).
			SetPiece(board.PieceSolider0, board.Coord{1, 3}).
			SetPiece(board.PieceSolider1, board.Coord{2, 3}).
			SetPiece(board.PieceSolider2, board.Coord{0, 4}).
			SetPiece(board.PieceSolider3, board.Coord{3, 4}).
			SetPiece(board.PieceEmpty0, board.Coord{1, 4}).
			SetPiece(board.PieceEmpty1, board.Coord{2, 4})
		return b, nil
	case 1:
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
		return b, nil
	case 2:
		b := board.NewBoard().
			SetPiece(board.PieceHuangZhong, board.Coord{0, 0}).
			SetPiece(board.PieceMaChao, board.Coord{1, 0}).
			SetPiece(board.PieceZhaoYun, board.Coord{2, 0}).
			SetPiece(board.PieceSolider0, board.Coord{3, 0}).
			SetPiece(board.PieceSolider1, board.Coord{3, 1}).
			SetPiece(board.PieceCaoCao, board.Coord{0, 2}).
			SetPiece(board.PieceZhangFei, board.Coord{3, 3}).
			SetPiece(board.PieceGuanYu, board.Coord{0, 4}).
			SetPiece(board.PieceSolider2, board.Coord{2, 3}).
			SetPiece(board.PieceSolider3, board.Coord{2, 4}).
			SetPiece(board.PieceEmpty0, board.Coord{2, 2}).
			SetPiece(board.PieceEmpty1, board.Coord{3, 2})
		return b, nil
	default:
		err := fmt.Errorf("invalid board id: %d\n", id)
		return nil, err
	}
}

func main() {
	flag.Parse()

	b, err := getBoard(*boardID)
	if err != nil {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("starting board:")
	b.PrettyPrint()

	if solved, solution := board.NewBFSSolver(*logVerbose).Solve(b); solved {
		fmt.Println("\nending board:")
		solution.EndingBoard.PrettyPrint()
		fmt.Println("\nsolution:")
		fmt.Printf("%s\n", solution.DumpMoves())

	} else {
		fmt.Printf("failed to find a solution\n")
	}
}
