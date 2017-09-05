package engine

import (
	"fmt"
	"strconv"
)

// Board represents a set of bitboards for each piece and colour
type Board struct {
	WhitePawns   int64
	WhiteKnights int64
	WhiteBishops int64
	WhiteRooks   int64
	WhiteQueens  int64
	WhiteKing    int64

	BlackPawns   int64
	BlackKnights int64
	BlackBishops int64
	BlackRooks   int64
	BlackQueens  int64
	BlackKing    int64
}

// InitBoard creates the initial board state
func InitBoard() {
	b := &Board{}
	arr := [8][8]rune{
		{'R', 'B', 'N', 'Q', 'K', 'N', 'B', 'R'},
		{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'},
		{'r', 'b', 'n', 'q', 'k', 'n', 'b', 'r'},
	}
	b.arrayToBoard(arr)
}

func (b *Board) arrayToBoard(arr [8][8]rune) {
	for i := 0; i < 64; i++ {
		bin := "0000000000000000000000000000000000000000000000000000000000000000"
		bin = bin[i+1:] + "1" + bin[0:i]

		switch arr[i/8][i%8] {
		case 'P':
			b.WhitePawns += convertStringToBitboard(bin)
		case 'N':
			b.WhiteKnights += convertStringToBitboard(bin)
		case 'B':
			b.WhiteBishops += convertStringToBitboard(bin)
		case 'R':
			b.WhiteRooks += convertStringToBitboard(bin)
		case 'Q':
			b.WhiteQueens += convertStringToBitboard(bin)
		case 'K':
			b.WhiteKing += convertStringToBitboard(bin)
		case 'p':
			b.BlackPawns += convertStringToBitboard(bin)
		case 'n':
			b.BlackKnights += convertStringToBitboard(bin)
		case 'b':
			b.BlackBishops += convertStringToBitboard(bin)
		case 'r':
			b.BlackRooks += convertStringToBitboard(bin)
		case 'q':
			b.BlackQueens += convertStringToBitboard(bin)
		case 'k':
			b.BlackKing += convertStringToBitboard(bin)
		}
	}

	b.drawBoard()
}

func convertStringToBitboard(str string) int64 {
	n, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		fmt.Printf("could not parse int: %v\n", err)
	}
	return n
}

func (b *Board) drawBoard() {
	out := [8][8]rune{}

	for i := 0; i < 64; i++ {
		out[i/8][i%8] = ' '
	}

	for i := 0; i < 64; i++ {
		if (b.WhitePawns >> uint(i)) == 1 {
			out[i/8][i%8] = 'P'
		}
		if (b.WhiteRooks >> uint(i)) == 1 {
			out[i/8][i%8] = 'R'
		}
		if (b.WhiteBishops >> uint(i)) == 1 {
			out[i/8][i%8] = 'B'
		}
		if (b.WhiteKnights >> uint(i)) == 1 {
			out[i/8][i%8] = 'N'
		}
		if (b.WhiteQueens >> uint(i)) == 1 {
			out[i/8][i%8] = 'Q'
		}
		if (b.WhiteKing >> uint(i)) == 1 {
			out[i/8][i%8] = 'K'
		}
		if (b.BlackPawns >> uint(i)) == 1 {
			out[i/8][i%8] = 'p'
		}
		if (b.BlackRooks >> uint(i)) == 1 {
			out[i/8][i%8] = 'r'
		}
		if (b.BlackBishops >> uint(i)) == 1 {
			out[i/8][i%8] = 'b'
		}
		if (b.BlackKnights >> uint(i)) == 1 {
			out[i/8][i%8] = 'n'
		}
		if (b.BlackQueens >> uint(i)) == 1 {
			out[i/8][i%8] = 'q'
		}
		if (b.BlackKing >> uint(i)) == 1 {
			out[i/8][i%8] = 'k'
		}
	}

	for i := 0; i < 8; i++ {
		fmt.Println(out)
	}
}
