package parser

import (
	"fmt"
	"github.com/Tanzor-Disco/GOndon/board"
	"strconv"
)

var startPosFEN = FEN{
	PiecePlacement: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
	ActiveColor:    "w",
	Castling:       "KQkq",
	EnPassant:      "-",
	HalfmoveClock:  0,
	FullmoveNumber: 1,
}

type PositionCommand struct {
	Position Position
}

func (PositionCommand) isCommand() {}

type FEN struct {
	PiecePlacement string
	ActiveColor    string
	Castling       string
	EnPassant      string
	HalfmoveClock  uint8
	FullmoveNumber uint16
}

type Position struct {
	FEN   FEN
	Moves []board.Move
}

func parseSquare(square string) (board.Square, error) {
	// Convert algebraic square notation into a Square index.
	// Square indices are ordered by rank, with 8 files per rank:
	// A1 = 0, B1 = 1, ..., H1 = 7, A2 = 8, ..., H8 = 63.
	// Therefore, the index is calculated as rank*8 + file

	if len(square) != 2 {
		return board.EmptySquare, fmt.Errorf("parseSquare: invalid string length")
	}

	file := square[0]
	if file < 'a' || file > 'h' {
		return board.EmptySquare, fmt.Errorf("parseSquare: invalid square letter")
	}

	fileIdx := file - 'a'
	rank := square[1]
	if rank < '1' || rank > '8' {
		return board.EmptySquare, fmt.Errorf("parseSquare: invalid square number")
	}
	rankIdx := (rank - '1') * 8
	return board.Square(rankIdx + fileIdx), nil
}

func parsePromotionPieceType(piece byte) (board.PieceType, error) {
	switch piece {
	case 'q':
		return board.Queen, nil
	case 'n':
		return board.Knight, nil
	case 'b':
		return board.Bishop, nil
	case 'r':
		return board.Rook, nil
	default:
		return board.NoPiece, fmt.Errorf("parsePromotionPieceType: invalid promotion piece type")
	}
}

func parseUCIMove(move string) (board.Move, error) {
	if len(move) != 5 && len(move) != 4 {
		return board.Move{}, fmt.Errorf("parseUCIMove: invalid UCI move length")
	}
	from, err := parseSquare(move[:2])
	if err != nil {
		return board.Move{}, fmt.Errorf("parseUCIMove: %w", err)
	}
	to, err := parseSquare(move[2:4])
	if err != nil {
		return board.Move{}, fmt.Errorf("parseUCIMove: %w", err)
	}
	var promotion board.PieceType
	if len(move) == 5 {
		var err error
		promotion, err = parsePromotionPieceType(move[4])
		if err != nil {
			return board.Move{}, fmt.Errorf("parseUCIMove: %w", err)
		}
	}
	return board.Move{
		From:      from,
		To:        to,
		Promotion: &promotion,
	}, nil
}

func parseUCIMoves(args []string) ([]board.Move, error) {
	moves := make([]board.Move, 0, len(args))
	for _, move := range args {
		parsed, err := parseUCIMove(move)
		if err != nil {
			return make([]board.Move, 0), fmt.Errorf("parseUCIMoves: %w", err)
		}
		moves = append(moves, parsed)
	}
	return moves, nil
}

func parseStartPos(args []string) (Position, error) {
	fen := startPosFEN
	if len(args) == 0 {
		return Position{
			FEN: fen,
		}, nil
	}
	subcommand := args[0]
	if subcommand != "moves" {
		return Position{}, fmt.Errorf("parseStartPos: invalid subcommand")
	}
	args = args[1:]

	moves, err := parseUCIMoves(args)
	if err != nil {
		return Position{}, fmt.Errorf("parseStartPos: %w", err)
	}

	return Position{
		FEN:   fen,
		Moves: moves,
	}, nil
}

func parseFEN(fen []string) (FEN, error) {
	var output FEN
	output.PiecePlacement = fen[0]
	output.ActiveColor = fen[1]
	output.Castling = fen[2]
	output.EnPassant = fen[3]
	halfmove, err := strconv.Atoi(fen[4])
	if err != nil {
		return FEN{}, fmt.Errorf("parseFEN: %w", err)
	}
	if halfmove < 0 {
		return FEN{}, fmt.Errorf("parseFEN: invalid halfmove value: %v", halfmove)
	}
	output.HalfmoveClock = uint8(halfmove)

	fullmove, err := strconv.Atoi(fen[5])
	if err != nil {
		return FEN{}, fmt.Errorf("parseFEN: %w", err)
	}
	if fullmove < 0 {
		return FEN{}, fmt.Errorf("parseFEN: invalid fullmove value: %v", fullmove)
	}
	output.FullmoveNumber = uint16(fullmove)

	return output, nil
}

func parseFenSubcommand(args []string) (Position, error) {
	if len(args) < 6 {
		return Position{}, fmt.Errorf("parseFenSubcommand: invalid amount of arguments passed")
	}
	fen, err := parseFEN(args[:6])
	if err != nil {
		return Position{}, fmt.Errorf("parseFenSubcommand: %w", err)
	}
	var moves []board.Move
	if len(args) > 6 {
		args = args[6:]
		subcommand := args[0]
		if subcommand != "moves" {
			return Position{}, fmt.Errorf("parseFenSubcommand: invalid subcommand")
		}
		args = args[1:]
		var err error
		moves, err = parseUCIMoves(args)
		if err != nil {
			return Position{}, fmt.Errorf("parseFenSubcommand: %w", err)
		}
	}
	return Position{
		FEN:   fen,
		Moves: moves,
	}, nil
}

func parsePosition(args []string) (Command, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("parsePosition: not enough arguments supplied")
	}

	subcommand := args[0]
	args = args[1:]
	switch subcommand {
	case "startpos":
		position, err := parseStartPos(args)
		if err != nil {
			return nil, fmt.Errorf("parsePosition: %w", err)
		}
		return PositionCommand{
			Position: position,
		}, nil

	case "fen":
		position, err := parseFenSubcommand(args)
		if err != nil {
			return nil, fmt.Errorf("parsePositionCommand: %w", err)
		}
		return PositionCommand{
			Position: position,
		}, nil

	default:
		return nil, fmt.Errorf("parsePositionCommand: invalid subcommand")
	}
}
