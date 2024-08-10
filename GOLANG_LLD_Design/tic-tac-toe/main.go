package main

import (
	"errors"
	"fmt"
)

type Game struct {
	board         *Board
	players       [2]Player
	currentPlayer Player
}

type Player struct {
	name   string
	symbol rune
}

func NewGame(player1, player2 Player) *Game {
	return &Game{
		board:   NewBoard(),
		players: [2]Player{player1, player2},
	}
}

// Board represents the Tic-Tac-Toe board
type Board struct {
	grid [3][3]rune
}

// NewBoard initializes a new board
func NewBoard() *Board {
	return &Board{}
}

// PlaceMark places a mark on the board
func (b *Board) PlaceMark(row, col int, mark rune) error {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return errors.New("invalid position")
	}
	if b.grid[row][col] != 0 {
		return errors.New("cell already occupied")
	}
	b.grid[row][col] = mark
	return nil
}

// Display prints the board
func (b *Board) Display() {
	for _, row := range b.grid {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%c ", cell)
			}
		}
		fmt.Println()
	}
}

func (b *Board) IsFull() bool {
	for _, row := range b.grid {
		for _, cell := range row {
			if cell == 0 {
				return false
			}
		}
	}
	return true
}

func (g *Game) MakeMove(row, col int) error {
	err := g.board.PlaceMark(row, col, g.currentPlayer.symbol)
	if err != nil {
		return err
	}
	g.board.Display()
	return nil
}

func (g *Game) CheckWinner() (Player, bool) {
	// check if b.currentPlayer is winner
	b := g.board.grid
	// winning conditions -
	lines := [][][2]int{
		{{0, 0}, {0, 1}, {0, 2}}, {{1, 0}, {1, 1}, {1, 2}}, {{2, 0}, {2, 1}, {2, 2}}, // rows
		{{0, 0}, {1, 0}, {2, 0}}, {{0, 1}, {1, 1}, {2, 1}}, {{0, 2}, {1, 2}, {2, 2}}, // columns
		{{0, 0}, {1, 1}, {2, 2}}, {{0, 2}, {1, 1}, {2, 0}}, // diagonals
	}

	for _, line := range lines {
		if b[line[0][0]][line[0][1]] != 0 && b[line[0][0]][line[0][1]] == b[line[1][0]][line[1][1]] &&
			b[line[0][0]][line[0][1]] == b[line[2][0]][line[2][1]] {
			return g.currentPlayer, true
		}
	}

	return Player{}, false
}

func (g *Game) IsDraw() bool {
	return g.board.IsFull()
}

func (g *Game) SwitchTurn() {
	if g.currentPlayer == g.players[0] {
		g.currentPlayer = g.players[1]
	} else {
		g.currentPlayer = g.players[0]
	}
}

func main() {

	var p1Name, p2Name string
	println("give first player name")
	_, err := fmt.Scanln(&p1Name)
	if err != nil {
		println(err)
		return
	}
	println(`and second player name`)
	_, err = fmt.Scanln(&p2Name)
	if err != nil {
		println(err)
		return
	}

	p1, p2 := Player{name: p1Name, symbol: 'x'}, Player{name: p2Name, symbol: 'o'}

	game := NewGame(p1, p2)
	game.currentPlayer = p1

	game.board.Display()
	var row, col int
	for {

		fmt.Printf("\n\n\n turn for %s give row <space> col\n\n", game.currentPlayer.name)

		if _, err := fmt.Scan(&row, &col); err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		err := game.MakeMove(row, col)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		if winner, won := game.CheckWinner(); won {
			fmt.Printf(" winner is %s!!!!!!\n\n", winner.name)
			break
		}
		if game.IsDraw() {
			println("game drawn!!")
			break
		}
		game.SwitchTurn()
	}
}
