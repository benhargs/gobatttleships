package main

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

//you can run all you tests by typing
//go test -v
//in the terminal window

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
	expectedCols := 7
	expectedRows := 7

	//Assert
	GridSizeCols := len(grid)
	if GridSizeCols != expectedCols {
		t.Errorf("Grid has wrong number of columns. Expected %d but was %d", expectedCols, GridSizeCols)
		//t.Errorf to allow error message with values %v.
	}

	GridSizeRows := len(grid[0])
	if GridSizeRows != expectedRows {
		t.Errorf("Grid has wrong number of rows. Expected %d, but was %d", expectedRows, GridSizeRows)
	}
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

// sometimes we write tests to test our own functions.
func TestGetRandomGridSquare(t *testing.T) {
	gridSquare := getRandomGridSquare()

	//literally only exists here to show you the output
	//should not exist in a real test
	fmt.Println("From TestGetRandomGridSquare.", gridSquare)

	//poor test making use of magic numbers
	//you should probably re-write it
	if gridSquare[0] < 1 || gridSquare[0] >= 8 {
		t.Error("Grid square row should be >-1 and <7, but got: ", gridSquare[0])
	}

	if gridSquare[1] < 1 || gridSquare[1] >= 8 {
		t.Error("Grid square column should be >-1 and <7, but got: ", gridSquare[1])
	}

}

func TestCannotPlaceShipOnAShip(t *testing.T) {
	// Arrange
	grid := CreateGrid()
	gridWithShip, _ := placeShip(grid, 3, 6)

	//Act
	_, shipErr := placeShip(gridWithShip, 3, 6)

	//Assert
	want := errors.New("ship already placed at coordinates row: 3 and column: 6")

	if shipErr.Error() != want.Error() {
		t.Errorf("wanted %v got %v", want, shipErr)
	}
}

func TestCannotPlaceShipOnAShipReportCoordinates(t *testing.T) {
	// Arrange
	grid := CreateGrid()
	gridWithShip, _ := placeShip(grid, 1, 5)

	//Act
	_, shipErr := placeShip(gridWithShip, 1, 5)

	//Assert
	want := errors.New("ship already placed at coordinates row: 1 and column: 5")

	if shipErr.Error() != want.Error() {
		t.Errorf("wanted %v got %v", want, shipErr)
	}
}

func TestCannotPlaceTenthShip(t *testing.T) {
	//Arrange (set things up)
	grid := CreateGrid()
	gridWith1Ship, _ := placeShip(grid, 1, 2)
	gridWith2Ships, _ := placeShip(gridWith1Ship, 2, 3)
	gridWith3Ships, _ := placeShip(gridWith2Ships, 3, 4)
	gridWith4Ships, _ := placeShip(gridWith3Ships, 4, 5)
	gridWith5Ships, _ := placeShip(gridWith4Ships, 5, 6)
	gridWith6Ships, _ := placeShip(gridWith5Ships, 6, 4)
	gridWith7Ships, _ := placeShip(gridWith6Ships, 5, 1)
	gridWith8Ships, _ := placeShip(gridWith7Ships, 1, 3)
	gridWith9Ships, _ := placeShip(gridWith8Ships, 2, 4)

	//Act
	_, got := placeShip(gridWith9Ships, 3, 5)

	//Assert
	want := errors.New("too many ships")

	if got.Error() != want.Error() {
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestPlacingTenthShipDoesNotChangeGrid(t *testing.T) {
	//Arrange (set things up)
	grid := CreateGrid()
	gridWith1Ship, _ := placeShip(grid, 1, 2)
	gridWith2Ships, _ := placeShip(gridWith1Ship, 2, 3)
	gridWith3Ships, _ := placeShip(gridWith2Ships, 3, 4)
	gridWith4Ships, _ := placeShip(gridWith3Ships, 4, 5)
	gridWith5Ships, _ := placeShip(gridWith4Ships, 5, 6)
	gridWith6Ships, _ := placeShip(gridWith5Ships, 6, 4)
	gridWith7Ships, _ := placeShip(gridWith6Ships, 5, 1)
	gridWith8Ships, _ := placeShip(gridWith7Ships, 1, 3)
	gridWith9Ships, _ := placeShip(gridWith8Ships, 2, 4)

	//Act
	got, _ := placeShip(gridWith9Ships, 3, 5)

	//Assert
	want := gridWith9Ships

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}

}

/*
func TestHasPlayerWon(t *testing.T) {
	//Arrange (set things up)
	grid := CreateGrid()
	gridWith1Ship, _ := placeShip(grid, 1, 2)
	gridWith2Ships, _ := placeShip(gridWith1Ship, 2, 3)
	gridWith3Ships, _ := placeShip(gridWith2Ships, 3, 4)
	gridWith4Ships, _ := placeShip(gridWith3Ships, 4, 5)
	gridWith5Ships, _ := placeShip(gridWith4Ships, 5, 6)
	gridWith6Ships, _ := placeShip(gridWith5Ships, 6, 4)
	gridWith7Ships, _ := placeShip(gridWith6Ships, 5, 1)
	gridWith8Ships, _ := placeShip(gridWith7Ships, 1, 3)
	gridWith9Ships, _ := placeShip(gridWith8Ships, 2, 4)

	gridWith1SunkShip, _ := shootAtOpponent(gridWith9Ships, 1, 2)
	gridWith2SunkShips, _ := shootAtOpponent(gridWith1SunkShip, 2, 3)
	gridWith3SunkShips, _ := shootAtOpponent(gridWith2SunkShips, 3, 4)
	gridWith4SunkShips, _ := shootAtOpponent(gridWith3SunkShips, 4, 5)
	gridWith5SunkShips, _ := shootAtOpponent(gridWith4SunkShips, 5, 6)
	gridWith6SunkShips, _ := shootAtOpponent(gridWith5SunkShips, 6, 4)
	gridWith7SunkShips, _ := shootAtOpponent(gridWith6SunkShips, 5, 1)
	gridWith8SunkShips, _ := shootAtOpponent(gridWith7SunkShips, 1, 3)
	gridWith9SunkShips, _ := shootAtOpponent(gridWith8SunkShips, 2, 4)

	//Act
	playerWin := hasPlayerWon(gridWith9SunkShips)

	//Assert
	if playerWin == false {
		t.Errorf("Player has not won.")
	}

}

func TestHasPlayerNotWon(t *testing.T) {
	//Arrange (set things up)
	grid := CreateGrid()
	gridWith1Ship, _ := placeShip(grid, 1, 2)
	gridWith2Ships, _ := placeShip(gridWith1Ship, 2, 3)
	gridWith3Ships, _ := placeShip(gridWith2Ships, 3, 4)
	gridWith4Ships, _ := placeShip(gridWith3Ships, 4, 5)
	gridWith5Ships, _ := placeShip(gridWith4Ships, 5, 6)
	gridWith6Ships, _ := placeShip(gridWith5Ships, 6, 4)
	gridWith7Ships, _ := placeShip(gridWith6Ships, 5, 1)
	gridWith8Ships, _ := placeShip(gridWith7Ships, 1, 3)
	gridWith9Ships, _ := placeShip(gridWith8Ships, 2, 4)

	gridWith1SunkShip, _ := shootAtOpponent(gridWith9Ships, 1, 2)
	gridWith2SunkShips, _ := shootAtOpponent(gridWith1SunkShip, 2, 3)
	gridWith3SunkShips, _ := shootAtOpponent(gridWith2SunkShips, 3, 4)
	gridWith4SunkShips, _ := shootAtOpponent(gridWith3SunkShips, 4, 5)
	gridWith5SunkShips, _ := shootAtOpponent(gridWith4SunkShips, 5, 6)
	gridWith6SunkShips, _ := shootAtOpponent(gridWith5SunkShips, 6, 4)
	gridWith7SunkShips, _ := shootAtOpponent(gridWith6SunkShips, 5, 1)
	gridWith8SunkShips, _ := shootAtOpponent(gridWith7SunkShips, 1, 3)

	//Act
	playerWin := hasPlayerWon(gridWith8SunkShips)

	//Assert
	if playerWin == true {
		t.Errorf("Player has won.")
	}
}
*/

func TestCannotPlaceShipOutsideGrid(t *testing.T) {
	type coordinates struct {
		row       int
		col       int
		errorText string
	}
	shipCoordinates := []coordinates{
		{row: 7, col: 6, errorText: "invalid row value: row = 7, want between 0 & 6 "},
		{row: -1, col: 0, errorText: "invalid row value: row = -1, want between 0 & 6 "},
		{row: 0, col: -1, errorText: "invalid column value: column = -1, want between 0 & 6 "},
		{row: 6, col: 7, errorText: "invalid column value: column = 7, want between 0 & 6 "},
	}

	//Act (run the code you want to do the thing)
	for _, coordinates := range shipCoordinates {
		//arrange
		grid := CreateGrid()

		//act
		_, got := placeShip(grid, coordinates.row, coordinates.col)

		//assert
		want := errors.New(coordinates.errorText)
		if got.Error() != want.Error() {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestCanPlaceShipAtEdgesOfGrid(t *testing.T) {
	type coordinates struct {
		row int
		col int
	}
	shipCoordinates := []coordinates{
		{row: 6, col: 6},
		{row: 0, col: 0},
		{row: 0, col: 6},
		{row: 6, col: 0},
	}

	//Act (run the code you want to do the thing)
	for _, coordinates := range shipCoordinates {
		//arrange
		grid := CreateGrid()

		//act
		_, got := placeShip(grid, coordinates.row, coordinates.col)

		//assert
		if got != nil {
			t.Errorf("got %v, want no error", got)
		}
	}
}

func TestCannotPlaceShipOutsideGridDoesntChangeGrid(t *testing.T) {
	type coordinates struct {
		row int
		col int
	}

	shipCoordinates := []coordinates{
		{row: 7, col: 6},
		{row: -1, col: 0},
		{row: 0, col: -1},
		{row: 4, col: 7},
	}

	for _, coordinates := range shipCoordinates {
		//arrange
		grid := CreateGrid()

		//act
		got, _ := placeShip(grid, coordinates.row, coordinates.col)

		//assert
		want := grid
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestHasShotHitShip(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	shipOnGrid, _ := placeShip(grid, 1, 2)

	//Act
	got := Fireshot{}
	got.shot(Fireshot{
		grid: shipOnGrid,
		shotCoords: Coordinates{
			row: 1,
			col: 2,
		},
	})

	//Assert
	want := Fireshot{
		grid: shipOnGrid,
		shotCoords: Coordinates{
			row: 1,
			col: 2,
		},
		hitShip: true,
	}
	if got.hitShip != want.hitShip {
		t.Error("shot did not hit ship")
	}
}

/*
Three return shootfunction.
func TestCannotShootOutsideGrid(t *testing.T) {
	type coordinates struct {
		row       int
		col       int
		errorText string
	}

	shotCoordinates := []coordinates{
		{row: 7, col: 6, errorText: "invalid row value: row = 7, want between 0 & 6 "},
		{row: -1, col: 0, errorText: "invalid row value: row = -1, want between 0 & 6 "},
		{row: 0, col: -1, errorText: "invalid column value: column = -1, want between 0 & 6 "},
		{row: 6, col: 7, errorText: "invalid column value: column = 7, want between 0 & 6 "},
	}

	for _, coordinates := range shotCoordinates {
		//arrange
		grid := CreateGrid()

		//act
		_, got, _ := shootAtOpponent(grid, coordinates.row, coordinates.col)

		//assert
		want := errors.New(coordinates.errorText)
		if got.Error() != want.Error() {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestHasShotHitShip(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	GridWithTargetShip, _ := placeShip(grid, 4, 5)

	//Act
	_, _, got := shootAtOpponent(GridWithTargetShip, 4, 5)

	//Assert
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestHasShotHitShipAndChangedPlayerGrid(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	GridWithTargetShip, _ := placeShip(grid, 4, 5)

	//Act
	got, _, _ := shootAtOpponent(GridWithTargetShip, 4, 5)

	//Assert
	grid[4][5] = "Sunk"
	want := grid

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestShipCannotBeSunkTwice(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	gridWithShip, _ := placeShip(grid, 1, 4)
	gridWithSunkShip, _, _ := shootAtOpponent(gridWithShip, 1, 4)

	//Act
	_, _, got := shootAtOpponent(gridWithSunkShip, 1, 4)

	//Assert
	want := false
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
*/

//TestShotReportsMiss
