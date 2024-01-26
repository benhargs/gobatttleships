package game

import (
	"errors"
	"fmt"
)

func CreateGrid() (grid [7][7]string) {
	return [7][7]string{}
}

func PlaceShip(grid [7][7]string, row int, col int) ([7][7]string, error) {
	coordErr := areCoordinatesOnPlayingGrid(row, col)

	if coordErr != nil {
		return grid, coordErr
	}

	if grid[row][col] == "Ship" {
		return grid, fmt.Errorf("ship already placed at coordinates row: %d and column: %d", row, col)
	}

	shipCount := countOfShipsOnBoard(grid)
	if shipCount == 9 {
		return grid, errors.New("too many ships")
	}

	grid[row][col] = "Ship"
	return grid, nil
}

func HasPlayerWon(grid [7][7]string) bool {
	numberOfSunkShips := 0
	for _, row := range grid {
		for _, coordinates := range row {
			if coordinates == "Sunk" {
				numberOfSunkShips++
			}
		}
	}

	if numberOfSunkShips == 9 {
		return true
	}

	return false
}

func CurrentPlayerTakeShot(player int, grid [7][7]string, row int, col int) (int, string, bool, error) {
	gridAfterShot, coordErr, shotResult := shootOpponent(grid, row, col)
	if shotResult == "miss" || shotResult == "hit" {
		gameResult := HasPlayerWon(gridAfterShot)
		if player == 1 {
			return 2, shotResult, gameResult, coordErr
		}

		return 1, shotResult, gameResult, coordErr
	}

	return player, "invalid", false, coordErr
}
func countOfShipsOnBoard(grid [7][7]string) int {
	numberOfShips := 0

	for _, row := range grid {
		for _, coordinates := range row {
			if coordinates == "Ship" {
				numberOfShips++
			}
		}
	}

	return numberOfShips
}

func shootOpponent(grid [7][7]string, row int, col int) ([7][7]string, error, string) {
	coordErr := areCoordinatesOnPlayingGrid(row, col)

	if coordErr != nil {
		return grid, coordErr, "invalid"
	}

	if grid[row][col] == "Ship" {
		grid[row][col] = "Sunk"
		return grid, nil, "hit"
	}

	return grid, nil, "miss"
}

func areCoordinatesOnPlayingGrid(row int, col int) error {
	if row < 0 || row > 6 {
		return fmt.Errorf("invalid row value: row = %d, want between 0 & 6 ", row)
	}
	if col < 0 || col > 6 {
		return fmt.Errorf("invalid column value: column = %d, want between 0 & 6 ", col)
	}
	return nil
}
