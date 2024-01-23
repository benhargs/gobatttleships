package main

import (
	"errors"
	"fmt"
)

/*
This game of battleships is very simple to start:
There are 2 players
Each player has a grid which is 7*7
Each player has 9 Battleships, each of which can occupy only one square on their grid
Each player can place their battleships anywhere on this grid
Players take it in turns to pick any grid square reference
If the player hits a battleship, then it is sunk, and the turn passes to the opponent
If the player misses a battleship then it is called a miss, and the turn passes to the opponent
The player to first sink all their opponent's battleships is the winner
*/

//All code in here is example code, you do not have to keep any of it.

func PlayerOneTurn(playerTwoGrid [7][7]string, shotCoordinates []int) (shotStatus bool) {
	return false //shot missed
}

func PlayerTwoTurn(playerOneGrid [7][7]string, shotCoordinates []int) (shotStatus bool) {
	return true //shot hit
}

func CreateGrid() (grid [7][7]string) {
	//this is a fixed array, not a slice
	return [7][7]string{}
}

func CountOfShipsOnBoard(grid [7][7]string) int {
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

func placeShip(grid [7][7]string, row int, col int) ([7][7]string, error) {
	coordErr := areCoordinatesOnPlayingGrid(row, col)

	if coordErr != nil {
		return grid, coordErr
	}

	if grid[row][col] == "Ship" {
		return grid, fmt.Errorf("ship already placed at coordinates row: %d and column: %d", row, col)
	}

	shipCount := CountOfShipsOnBoard(grid)
	if shipCount == 9 {
		return grid, errors.New("too many ships")
	}

	grid[row][col] = "Ship"
	return grid, nil
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

func hasPlayerWon(grid [7][7]string) bool {
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

func changeActivePlayer(player int) int {
	if player == 1 {
		return 2
	}
	return 1
}

func currentPlayerTakeShot(player int, grid [7][7]string, row int, col int) (int, string, bool, error) {
	gridAfterShot, coordErr, shotResult := shootOpponent(grid, row, col)

	if shotResult == "miss" {
		nextPlayer := changeActivePlayer(player)
		return nextPlayer, "miss", false, coordErr
	}

	if shotResult == "hit" {
		nextPlayer := changeActivePlayer(player)
		gameResult := hasPlayerWon(gridAfterShot)
		return nextPlayer, "hit", gameResult, coordErr
	}

	return player, "invalid", false, coordErr
}
