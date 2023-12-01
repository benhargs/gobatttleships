package main

import (
	"fmt"
	"math/rand"
	"testing"
)

//you can run all you tests by typing
//go test -v
//in the terminal window

// this is a utility function for testing
// it will return a random square on the grid
// it does not keep track of any previously returned grids
func getRandomGridSquare() []int {

	row := []int{1, 2, 3, 4, 5, 6, 7}
	column := []int{1, 2, 3, 4, 5, 6, 7}

	return []int{rand.Intn(len(row)) + 1, rand.Intn(len(column)) + 1}

}

//these are the two tests we have for our functions in main
//the purpose of tests is to mimic interaction with our code
//there is no "user input" - the test is the calling code

// here is an example of a failing test - what do we need to do to fix it?
func TestCreateGrid(t *testing.T) {

	//Act
	grid := CreateGrid()

	//Assert
	assertGridIsCorrectSize(t, grid, 7, 7)
}

//one good place to start here is by using our utility function
//to target a random grid square rather than 1,1 co-ordinates every time

func TestPlayerOneTakingShot(t *testing.T) {
	// Arrange
	grid := CreateGrid()

	//Act - Code Under Test - System under Test (SUT) - "Production Code"
	shotResult := PlayerOneTurn(grid, []int{1, 1})

	// Assert - check the result is what we expected
	if shotResult != false {
		t.Error("Shot should be false!")
	}
}

func TestPlayerTwoTakingShot(t *testing.T) {
	grid := CreateGrid()
	shotResult := PlayerTwoTurn(grid, []int{1, 1})
	if shotResult != true {
		t.Error("Shot should be true!")
	}
}

//other tests here that fail

// sometimes we write tests to test our own functions.
func TestGetRandomGridSquare(t *testing.T) {
	gridSquare := getRandomGridSquare()

	//literally only exists here to show you the output
	//should not exist in a real test
	fmt.Println(gridSquare)

	//poor test making use of magic numbers
	//you should probably re-write it
	if gridSquare[0] <= 0 || gridSquare[0] >= 8 {
		t.Error("Grid square row should be >0 and <8, but got: ", gridSquare[0])
	}

	if gridSquare[1] <= 0 || gridSquare[1] >= 8 {
		t.Error("Grid square column should be >0 and <8, but got: ", gridSquare[1])
	}
}

func TestPlaceAShip(t *testing.T) {
	// Arrange
	grid := CreateGrid()

	//Act
	desiredCol := 3
	desiredRow := 5
	updatedGrid := placeShip(grid, desiredCol, desiredRow)

	// ... by here the ship will have been place on the grid!

	//Assert
	// A ship is placed
	actual := updatedGrid[3][5]
	want := "S"
	if actual != want {
		t.Error("Ship was not placed at col 3, row 5")
	}
}

// Test to see if ship is already there

func TestIsThereAShip(t *testing.T) {
	//Arrange
	grid := CreateGrid()

	ShipCol := 1
	ShipRow := 2
	updatedGrid := placeShip(grid, ShipCol, ShipRow)

	EmptyCol := 2
	EmptyRow := 2

	//Act

	IsEmpty := isShipAt(updatedGrid, EmptyCol, EmptyRow)

	//Assert

	if IsEmpty == false {
		t.Error("These Coordinates already have a ship.")
	}

}

//Shot at ship

func TestShootAShip(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	shipCol := 2
	shipRow := 2

	updatedGrid := placeShip(grid, shipCol, shipRow)

	shotCol := 2
	shotRow := 2

	//Act

	isHit := isShipHit(updatedGrid, shotCol, shotRow)

	//Assert
	if isHit == false {
		t.Error("Your shot missed!")
	}
}

//Already shot at ship

func assertGridIsCorrectSize(t *testing.T, grid [7][7]string, expectedRows int, expectedCols int) {
	GridSizeCols := len(grid)
	if GridSizeCols != expectedCols {
		t.Error("Grid has wrong number of columns. Wanted 7 but was", GridSizeCols)
	}

	GridSizeRows := len(grid[0])
	if GridSizeRows != expectedRows {
		t.Error("Grid has wrong number of rows. Wanted 7, but was", GridSizeRows)
	}
}
