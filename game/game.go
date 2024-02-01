package game

import (
	"errors"
	"fmt"
)

func CreateGrid() (grid [7][7]string) {
	return [7][7]string{}
}

func PlaceShip(grid [7][7]string, row int, col int) ([7][7]string, error) {
	maxShip := 9
	coordErr := areCoordinatesOnPlayingGrid(row, col)

	if coordErr != nil {
		return grid, coordErr
	}

	if grid[row][col] == "Ship" {
		return grid, fmt.Errorf("ship already placed at coordinates row: %d and column: %d", row, col)
	}

	shipCount := countOfShipsOnGrid(grid)
	if shipCount == maxShip {
		return grid, errors.New("too many ships")
	}

	grid[row][col] = "Ship"
	return grid, nil
}

func CurrentPlayerTakeShot(player int, grid [7][7]string, row int, col int) (int, string, bool, error) {
	gridAfterShot, coordErr, shotResult := shootOpponent(grid, row, col)

	if coordErr != nil {
		return player, shotResult, false, coordErr
	}

	newPlayer := changePlayer(player)

	if shotResult == "Hit" {
		gameResult := HasPlayerWon(gridAfterShot)
		return newPlayer, shotResult, gameResult, coordErr
	}

	return newPlayer, shotResult, false, coordErr
}

func HasPlayerWon(grid [7][7]string) bool {
	numberOfShips := countOfShipsOnGrid(grid)
	if numberOfShips != 0 {
		return false
	}
	return true
}

func shootOpponent(grid [7][7]string, row int, col int) ([7][7]string, error, string) {
	coordErr := areCoordinatesOnPlayingGrid(row, col)

	if coordErr != nil {
		return grid, coordErr, "Miss"
	}

	if grid[row][col] == "Ship" {
		grid[row][col] = ""
		return grid, nil, "Hit"
	}

	return grid, nil, "Miss"
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

func countOfShipsOnGrid(grid [7][7]string) int {
	shipCount := 0
	for _, row := range grid {
		for _, coordinates := range row {
			if coordinates == "Ship" {
				shipCount++
			}
		}
	}
	return shipCount
}

func changePlayer(player int) int {
	if player == 1 {
		return 2
	}
	return 1
}
