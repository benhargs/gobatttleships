package game

import (
	turn "battleships/turn"
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

	shipCount := countOfShipsOnBoard(grid)
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

	newPlayer := turn.ChangePlayer(player)

	if shotResult == "hit" {
		gameResult := HasPlayerWon(gridAfterShot)
		return newPlayer, shotResult, gameResult, coordErr
	}

	return newPlayer, shotResult, false, coordErr
}

func HasPlayerWon(grid [7][7]string) bool {
	numberOfSunkShips := countOfElementOnGrid(grid, "Sunk")
	if numberOfSunkShips != 9 {
		return false
	}
	return true
}

func countOfShipsOnBoard(grid [7][7]string) int {
	numberOfShips := countOfElementOnGrid(grid, "Ship")
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

func countOfElementOnGrid(grid [7][7]string, element string) int {
	countOfElement := 0
	for _, row := range grid {
		for _, coordinates := range row {
			if coordinates == element {
				countOfElement++
			}
		}
	}
	return countOfElement
}
