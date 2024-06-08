package pieces

import "fmt"

type Piece struct {
	Kind     string
	Color    string
	HasMoved bool
}

type Square struct {
	HPos  int
	VPos  int
	Piece Piece
}

type Board struct {
	Squares [8][8]Square
}

func (square Square) GetMove(board *Board) []*Square {

}

func PawnMove(s Square, board Board, turn *int) ([]*Square, error) {
	p := s.Piece

	if p.Kind != "Pawn" {
		err := fmt.Errorf("Wrong piece type, expected Pawn, got %s", p.Kind)

		return nil, err
	}

	if p.HasMoved {
		if board.Squares[s.VPos][s.HPos+1].Piece.Kind == "" {
		}
	}
}
