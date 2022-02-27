package board

import (
	"fmt"
	"math/rand"

	"github.com/golang-collections/collections/queue"
)

// NewBFSSolver creates a solver which Breadth First Search to solve with minimal steps
func NewBFSSolver() BoardSolver {
	return &bfsSolver{}
}

type bfsSolver struct {
}

type bfsNode struct {
	board *Board
	moves []*Move
}

// hash friendly to dup board state
type bfsBoardState [TotalGrids]PieceType

// State converts the board state into a hashable byte array
func (sv *bfsSolver) State(board *Board) bfsBoardState {
	var s bfsBoardState
	for i := 0; i < TotalGrids; i++ {
		s[i] = board.grids[i].Type
	}
	return bfsBoardState(s)
}

func (sv *bfsSolver) randomizeDirs(dirs []Coord) {
	size := len(dirs)
	for i := range dirs {
		j := rand.Intn(size)
		if i == j {
			continue
		}
		// swap i with j
		t := dirs[j]
		dirs[j] = dirs[i]
		dirs[i] = t
	}
}

func (sv *bfsSolver) Solve(board *Board) (bool, *Solution) {
	q := queue.New()
	q.Enqueue(&bfsNode{board: board})
	visited := map[bfsBoardState]bool{sv.State(board): true}

	fmt.Printf("======\n")
	boards := 1
	var solved *bfsNode

	for solved == nil && q.Len() > 0 {
		node := q.Dequeue().(*bfsNode)
		b := node.board

		dirs := []Coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		// sv.randomizeDirs(dirs)
		emptyPiecePositions := []Coord{b.GetPiecePosition(PieceEmpty0), b.GetPiecePosition(PieceEmpty1)}

	DONE:
		for _, ep := range emptyPiecePositions {
			for _, m := range dirs {
				px := ep.X - m.X
				py := ep.Y - m.Y
				canMove := func(b *Board, p *Piece, x, y int8) bool {
					v := b.CanMovePiece(p, x, y)
					// fmt.Printf("CanMovePiece: %v, X:%d, Y:%d\n", *p, x, y)
					return v
				}
				if p := b.GetPiece(px, py); p != nil && p.Type != PieceTypeEmpty && canMove(b, p, m.X, m.Y) {
					nb := b.Clone().MovePiece(p, m.X, m.Y)
					s := sv.State(nb)
					if !visited[s] {
						visited[s] = true
						// make copy of moves so BFS won't override it
						moves := make([]*Move, 0, len(node.moves))
						for _, m := range node.moves {
							moves = append(moves, m)
						}
						pos := b.GetPiecePosition(p)
						m := Move{P: p, CurX: pos.X, CurY: pos.Y, X: m.X, Y: m.Y}
						moves = append(moves, &m)
						newNode := bfsNode{board: nb, moves: moves}
						q.Enqueue(&newNode)
						boards++
						if boards%1000 == 0 {
							fmt.Printf("  -- searched boards %d\n", boards)
						}
						// fmt.Println("board: ", nb.String())
						if nb.Solved() {
							fmt.Printf("  -- searched total boards %d\n", boards)
							solved = &newNode
							break DONE
						}
					}
				}
			}
		}
	}

	if solved != nil {
		solution := Solution{
			Moves: solved.moves,
		}
		return true, &solution
	}

	return false, nil
}
