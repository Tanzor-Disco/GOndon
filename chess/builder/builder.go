package builder

import (
	"fmt"
	chessboard "github.com/Tanzor-Disco/GOndon/chess/board"
	"github.com/Tanzor-Disco/GOndon/uci/parser"
	"slices"
	"strings"
	"unicode"
)

type Position struct {
	Board          chessboard.Board
	ActiveColor    chessboard.SideColor
	Castling       parser.CastlingRights
	EnPassant      *chessboard.Square
	HalfmoveClock  uint8
	FullmoveNumber uint16
}

func buildElement(el rune, board *chessboard.Board, pieceIdx *chessboard.Square) error {
	if el >= '1' && el <= '8' {
		num := int(el - '0')
		*pieceIdx += chessboard.Square(num)
		return nil
	}

	var color chessboard.SideColor
	if unicode.IsUpper(el) {
		color = chessboard.White
	} else {
		color = chessboard.Black
	}

	lower := unicode.ToLower(el)
	pieceType, ok := chessboard.PieceTypeMap[lower]
	if !ok {
		return fmt.Errorf("buildElement: invalid piece type: %q", el)
	}

	if err := board.AddPiece(*pieceIdx, pieceType, color); err != nil {
		return fmt.Errorf("buildElement: %w", err)
	}
	*pieceIdx++
	return nil
}

func buildPieces(placement string, moves []chessboard.Move, board *chessboard.Board) error {
	// FEN start from rank 8; from A to H
	var pieceIdx chessboard.Square
	parts := strings.Split(placement, "/")
	for _, part := range slices.Backward(parts) {
		for _, el := range part {
			if err := buildElement(el, board, &pieceIdx); err != nil {
				return fmt.Errorf("buildPieces: %w", err)
			}
		}
	}

	for _, move := range moves {
		if err := board.MakeMove(move); err != nil {
			return fmt.Errorf("buildPieces: %w", err)
		}
	}
	return nil
}

func BuildPosition(position parser.Position) (Position, error) {
	var board chessboard.Board
	if err := buildPieces(position.FEN.PiecePlacement, position.Moves, &board); err != nil {
		return Position{}, fmt.Errorf("BuildPosition: %w", err)
	}

	return Position{
		Board:          board,
		ActiveColor:    position.FEN.ActiveColor,
		Castling:       position.FEN.Castling,
		EnPassant:      position.FEN.EnPassant,
		HalfmoveClock:  position.FEN.HalfmoveClock,
		FullmoveNumber: position.FEN.FullmoveNumber,
	}, nil
}
