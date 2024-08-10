# Design a tic tac toe game

## Requirements

1. The game should support two players.
2. It should be able to determine the winner or if the game is a draw.
3. The board is a 3x3 grid.
4. Players take turns to place their marks (X or O) on the board.

### Structures and Their Relationships

#### 1. `Game`

- **Attributes:**
  - `board: Board`
  - `players: [2]Player`
  - `currentTurn: int` (index to track the current player, 0 or 1)
- **Methods:**
  - `NewGame(player1, player2 Player) *Game`
  - `MakeMove(row, col int) error`
  - `CheckWinner() (Player, bool)`
  - `IsDraw() bool`
  - `SwitchTurn()`

#### 2. `Board`

- **Attributes:**
  - `grid: [3][3]rune`
- **Methods:**
  - `NewBoard() *Board`
  - `Display()`
  - `PlaceMark(row, col int, mark rune) error`
  - `IsFull() bool`

#### 3. `Player`

- **Attributes:**
  - `name: string`
  - `mark: rune` (either 'X' or 'O')
- **Methods:**
  - `NewPlayer(name string, mark rune) Player`

### Code Implementation

Below is a Go implementation of the Tic-Tac-Toe game based on the above design:

### Explanation

1. **Player:** Represents a player with a name and mark ('X' or 'O').
2. **Board:** Manages the 3x3 grid. It can display the board, place marks, and check if the board is full.
3. **Game:** Manages the overall game flow. It initializes the board and players, handles making moves, checks for winners or draws, and switches turns.

This design and implementation cover the requirements for a basic Tic-Tac-Toe game in Go. The `main` function demonstrates a simple game loop where players input their moves, and the game checks for a winner or draw after each move.