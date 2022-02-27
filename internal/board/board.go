package board

import (
	"fmt"
	"strings"
)

const (
	XGrids     = 4
	YGrids     = 5
	TotalGrids = XGrids * YGrids
)

type PieceType int8

const (
	PieceTypeCaoCao PieceType = iota
	PieceType1X2
	PieceType2X1
	PieceType1X1
	PieceTypeEmpty // make one for empty space for convinience
)

type Coord struct {
	X, Y int8
}

type Piece struct {
	Name string
	Type PieceType
	DimX int8
	DimY int8
}

func (p *Piece) Occupy() []Coord {
	c := make([]Coord, 0, int(p.DimX*p.DimY))
	for j := 0; j < (int)(p.DimY); j++ {
		for i := 0; i < int(p.DimX); i++ {
			c = append(c, Coord{X: int8(i), Y: int8(j)})
		}
	}
	return c
}

var PieceCaoCao = &Piece{Name: "CaoCao", Type: PieceTypeCaoCao, DimX: 2, DimY: 2}
var PieceGuanYu = &Piece{Name: "GuanYu", Type: PieceType2X1, DimX: 2, DimY: 1}
var PieceZhangFei = &Piece{Name: "ZhangFei", Type: PieceType1X2, DimX: 1, DimY: 2}
var PieceHuangZhong = &Piece{Name: "HuangZhong", Type: PieceType1X2, DimX: 1, DimY: 2}
var PieceZhaoYun = &Piece{Name: "ZhaoYun", Type: PieceType1X2, DimX: 1, DimY: 2}
var PieceMaChao = &Piece{Name: "MaChao", Type: PieceType1X2, DimX: 1, DimY: 2}
var PieceSolider0 = &Piece{Name: "Soldier0", Type: PieceType1X1, DimX: 1, DimY: 1}
var PieceSolider1 = &Piece{Name: "Soldier1", Type: PieceType1X1, DimX: 1, DimY: 1}
var PieceSolider2 = &Piece{Name: "Soldier2", Type: PieceType1X1, DimX: 1, DimY: 1}
var PieceSolider3 = &Piece{Name: "Soldier3", Type: PieceType1X1, DimX: 1, DimY: 1}
var PieceEmpty0 = &Piece{Name: "Empty0", Type: PieceTypeEmpty, DimX: 1, DimY: 1}
var PieceEmpty1 = &Piece{Name: "Empty1", Type: PieceTypeEmpty, DimX: 1, DimY: 1}

var AllPieces []*Piece = []*Piece{
	PieceCaoCao,
	PieceGuanYu,
	PieceZhangFei,
	PieceHuangZhong,
	PieceZhaoYun,
	PieceMaChao,
	PieceSolider0,
	PieceSolider1,
	PieceSolider2,
	PieceSolider3,
	PieceEmpty0,
	PieceEmpty1,
}

type Move struct {
	P          *Piece
	CurX, CurY int8
	X, Y       int8
}

func (m *Move) Dir() string {
	if m.X > 0 {
		return "right"
	} else if m.X < 0 {
		return "left"
	} else if m.Y > 0 {
		return "down"
	} else {
		return "up"
	}
}

func (m *Move) String() string {
	return fmt.Sprintf("%s move %s", m.P.Name, m.Dir())
}

type Board struct {
	grids                    []*Piece
	pieceMap                 map[*Piece]int8
	emptyPiece0, emptyPiece1 int8 // grid index
	moves                    []*Move
}

// NewBoard is board constructor
func NewBoard() *Board {
	return &Board{
		grids:    make([]*Piece, TotalGrids),
		pieceMap: map[*Piece]int8{},
	}
}

func (b *Board) Grids() []*Piece {
	return b.grids
}

func (b *Board) gridIndex(x, y int8) int8 {
	return y*XGrids + x
}

func (b *Board) indexCoord(idx int8) Coord {
	if idx < 0 || idx > TotalGrids {
		panic(fmt.Sprintf("invalid index %d", idx))
	}
	y := idx / XGrids
	x := idx - y*XGrids
	return Coord{X: x, Y: y}
}

// set the top/left corner to the grid x,y
func (b *Board) SetPiece(piece *Piece, c Coord) *Board {
	x, y := c.X, c.Y
	b.pieceMap[piece] = b.gridIndex(x, y)
	for _, occ := range piece.Occupy() {
		idx := b.gridIndex(x+occ.X, y+occ.Y)
		b.grids[idx] = piece
		// fmt.Printf("SetPiece: %v, idX: %d, X:%d, Y:%d, occ.X:%d, occ.Y:%d \n", *piece, idx, x, y, occ.x, occ.y)
	}
	return b
}

// GetPiece returns the piece at grid(x,y)
func (b *Board) GetPiece(x, y int8) *Piece {
	if x < 0 || x >= XGrids || y < 0 || y >= YGrids {
		return nil
	}
	idx := b.gridIndex(x, y)
	return b.grids[idx]
}

// GetPiecePosition returns the current position of a piece
func (b *Board) GetPiecePosition(piece *Piece) Coord {
	idx := b.pieceMap[piece]
	return b.indexCoord(idx)
}

// CanMovePiece returns whether a piece can move to (x,y) offset ie, (0,1)|(0,-1)|(1,0)|(-1,0)
func (b *Board) CanMovePiece(piece *Piece, x, y int8) bool {
	pos := b.GetPiecePosition(piece)
	nx := pos.X + x
	ny := pos.Y + y

	// out of board
	if nx < 0 || ny < 0 || nx+piece.DimX > XGrids || ny+piece.DimY > YGrids {
		return false
	}

	if x > 0 {
		if b.GetPiece(pos.X+piece.DimX, pos.Y).Type != PieceTypeEmpty {
			return false
		}
		if piece.DimY == 2 && b.GetPiece(pos.X+piece.DimX, pos.Y+1).Type != PieceTypeEmpty {
			return false
		}
	} else if x < 0 {
		if b.GetPiece(pos.X+x, pos.Y).Type != PieceTypeEmpty {
			return false
		}
		if piece.DimY == 2 && b.GetPiece(pos.X+x, pos.Y+1).Type != PieceTypeEmpty {
			return false
		}
	} else if y > 0 {
		if b.GetPiece(pos.X, pos.Y+piece.DimY).Type != PieceTypeEmpty {
			return false
		}
		if piece.DimX == 2 && b.GetPiece(pos.X+1, pos.Y+piece.DimY).Type != PieceTypeEmpty {
			return false
		}
	} else if y < 0 {
		if b.GetPiece(pos.X, pos.Y+y).Type != PieceTypeEmpty {
			return false
		}
		if piece.DimX == 2 && b.GetPiece(pos.X+1, pos.Y+y).Type != PieceTypeEmpty {
			return false
		}
	}
	return true
}

// MovePiece moves the piece in direction (up/down/left/right) specified in (x,y)
func (b *Board) MovePiece(piece *Piece, x, y int8) *Board {
	pos := b.GetPiecePosition(piece)
	// keep record of the move to get to the next board
	b.moves = append(b.moves, &Move{P: piece, CurX: pos.X, CurY: pos.Y, X: x, Y: y})

	// set/move empty pieces
	if x > 0 {
		b.SetPiece(b.GetPiece(pos.X+piece.DimX, pos.Y), Coord{X: pos.X, Y: pos.Y})
		if piece.DimY == 2 {
			b.SetPiece(b.GetPiece(pos.X+piece.DimX, pos.Y+1), Coord{X: pos.X, Y: pos.Y + 1})
		}
	} else if x < 0 {
		b.SetPiece(b.GetPiece(pos.X+x, pos.Y), Coord{X: pos.X + piece.DimX - 1, Y: pos.Y})
		if piece.DimY == 2 {
			b.SetPiece(b.GetPiece(pos.X+x, pos.Y+1), Coord{X: pos.X + piece.DimX - 1, Y: pos.Y + 1})
		}
	} else if y > 0 {
		b.SetPiece(b.GetPiece(pos.X, pos.Y+piece.DimY), Coord{X: pos.X, Y: pos.Y})
		if piece.DimX == 2 {
			b.SetPiece(b.GetPiece(pos.X+1, pos.Y+piece.DimY), Coord{X: pos.X + 1, Y: pos.Y})
		}
	} else if y < 0 {
		b.SetPiece(b.GetPiece(pos.X, pos.Y+y), Coord{X: pos.X, Y: pos.Y + piece.DimY - 1})
		if piece.DimX == 2 {
			b.SetPiece(b.GetPiece(pos.X+1, pos.Y+y), Coord{X: pos.X + 1, Y: pos.Y + piece.DimY - 1})
		}
	}

	b.SetPiece(piece, Coord{X: pos.X + x, Y: pos.Y + y})
	return b
}

// Clone creates a new board with the same state.
// need this since BFS needs to remember all different states
func (b *Board) Clone() *Board {
	b2 := NewBoard()
	for _, p := range AllPieces {
		b2.SetPiece(p, b.GetPiecePosition(p))
	}
	return b2
}

// Solved returns true if board is in solved state, aka, CaoCao is at the gate
func (b *Board) Solved() bool {
	pos := b.GetPiecePosition(PieceCaoCao)
	if pos.X == 1 && pos.Y == 3 {
		return true
	}
	return false
}

// String outputs the grid status
func (b *Board) String() string {
	var s []string
	for i := 0; i < TotalGrids; i++ {
		g := b.grids[i]
		s = append(s, fmt.Sprintf("[%d]: %v", i, g))
	}
	return strings.Join(s, ",")
}

type Solution struct {
	Moves []*Move
}

// DumpMoves outputs the move history into a string for convenience
func (s *Solution) DumpMoves() string {
	ss := []string{}
	for i, m := range s.Moves {
		ss = append(ss, fmt.Sprintf("[%d] %s from pos:(%d,%d) move %s\n", i, m.P.Name, m.CurX, m.CurY, m.Dir()))
	}
	return strings.Join(ss, "\n")
}

type BoardSolver interface {
	Solve(*Board) (bool, *Solution)
}
