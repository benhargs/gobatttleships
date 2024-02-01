package game

import (
	"errors"
	"testing"
)

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

func TestCannotPlaceShipOnAShip(t *testing.T) {
	// Arrange
	grid := CreateGrid()
	gridWithShip, _ := PlaceShip(grid, 3, 6)

	//Act
	_, shipErr := PlaceShip(gridWithShip, 3, 6)

	//Assert
	want := errors.New("ship already placed at coordinates row: 3 and column: 6")

	if shipErr.Error() != want.Error() {
		t.Errorf("wanted %v got %v", want, shipErr)
	}
}

func TestCannotPlaceShipOnAShipReportCoordinates(t *testing.T) {
	// Arrange
	grid := CreateGrid()
	gridWithShip, _ := PlaceShip(grid, 1, 5)

	//Act
	_, shipErr := PlaceShip(gridWithShip, 1, 5)

	//Assert
	want := errors.New("ship already placed at coordinates row: 1 and column: 5")

	if shipErr.Error() != want.Error() {
		t.Errorf("wanted %v got %v", want, shipErr)
	}
}

func TestCannotPlaceTenthShip(t *testing.T) {
	//Arrange (set things up)
	grid := CreateGrid()
	gridWith1Ship, _ := PlaceShip(grid, 1, 2)
	gridWith2Ships, _ := PlaceShip(gridWith1Ship, 2, 3)
	gridWith3Ships, _ := PlaceShip(gridWith2Ships, 3, 4)
	gridWith4Ships, _ := PlaceShip(gridWith3Ships, 4, 5)
	gridWith5Ships, _ := PlaceShip(gridWith4Ships, 5, 6)
	gridWith6Ships, _ := PlaceShip(gridWith5Ships, 6, 4)
	gridWith7Ships, _ := PlaceShip(gridWith6Ships, 5, 1)
	gridWith8Ships, _ := PlaceShip(gridWith7Ships, 1, 3)
	gridWith9Ships, _ := PlaceShip(gridWith8Ships, 2, 4)

	//Act
	_, got := PlaceShip(gridWith9Ships, 3, 5)

	//Assert
	want := errors.New("too many ships")

	if got.Error() != want.Error() {
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestPlacingTenthShipDoesNotChangeGrid(t *testing.T) {
	//Arrange (set things up)
	grid := CreateGrid()
	gridWith1Ship, _ := PlaceShip(grid, 1, 2)
	gridWith2Ships, _ := PlaceShip(gridWith1Ship, 2, 3)
	gridWith3Ships, _ := PlaceShip(gridWith2Ships, 3, 4)
	gridWith4Ships, _ := PlaceShip(gridWith3Ships, 4, 5)
	gridWith5Ships, _ := PlaceShip(gridWith4Ships, 5, 6)
	gridWith6Ships, _ := PlaceShip(gridWith5Ships, 6, 4)
	gridWith7Ships, _ := PlaceShip(gridWith6Ships, 5, 1)
	gridWith8Ships, _ := PlaceShip(gridWith7Ships, 1, 3)
	gridWith9Ships, _ := PlaceShip(gridWith8Ships, 2, 4)

	//Act
	got, _ := PlaceShip(gridWith9Ships, 3, 5)

	//Assert
	want := gridWith9Ships

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}

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
		_, got := PlaceShip(grid, coordinates.row, coordinates.col)

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
		_, got := PlaceShip(grid, coordinates.row, coordinates.col)

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
		got, _ := PlaceShip(grid, coordinates.row, coordinates.col)

		//assert
		want := grid
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestReportsShipBeingHit(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	shipOnGrid, _ := PlaceShip(grid, 1, 2)

	//Act
	_, _, got := shootOpponent(shipOnGrid, 1, 2)

	//Assert
	want := true
	if got != want {
		t.Errorf("shot did not report a hit ship, got %v want %v", got, want)
	}
}

func TestShipCannotBeShotTwice(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	gridWithShip, _ := PlaceShip(grid, 1, 2)
	gridWithSunkShip, _, _ := shootOpponent(gridWithShip, 1, 2)

	//Act
	_, _, got := shootOpponent(gridWithSunkShip, 1, 2)

	//Arrange
	want := false
	if got != want {
		t.Errorf("shot was not a miss, got %v want %v", got, want)
	}
}

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
		_, got, _ := shootOpponent(grid, coordinates.row, coordinates.col)

		//assert
		want := errors.New(coordinates.errorText)
		if got.Error() != want.Error() {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestReportCanReportMissAtEdgesOfGrid(t *testing.T) {
	type coordinates struct {
		row int
		col int
	}

	shotCoordinates := []coordinates{
		{row: 0, col: 6},
		{row: 0, col: 0},
		{row: 6, col: 0},
		{row: 6, col: 6},
	}

	for _, coordinates := range shotCoordinates {
		//arrange
		grid := CreateGrid()

		//act
		_, _, got := shootOpponent(grid, coordinates.row, coordinates.col)

		//assert
		want := false
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestReportCanShootAtEdgesOfGrid(t *testing.T) {
	type coordinates struct {
		row int
		col int
	}

	shotCoordinates := []coordinates{
		{row: 0, col: 6},
		{row: 0, col: 0},
		{row: 6, col: 0},
		{row: 6, col: 6},
	}

	for _, coordinates := range shotCoordinates {
		//arrange
		grid := CreateGrid()

		//act
		_, _, got := shootOpponent(grid, coordinates.row, coordinates.col)

		//assert
		want := false
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestHasPlayerWon(t *testing.T) {
	//Arrange (set things up)
	grid := CreateGrid()
	gridWith1Ship, _ := PlaceShip(grid, 1, 2)
	gridWith2Ships, _ := PlaceShip(gridWith1Ship, 2, 3)
	gridWith3Ships, _ := PlaceShip(gridWith2Ships, 3, 4)
	gridWith4Ships, _ := PlaceShip(gridWith3Ships, 4, 5)
	gridWith5Ships, _ := PlaceShip(gridWith4Ships, 5, 6)
	gridWith6Ships, _ := PlaceShip(gridWith5Ships, 6, 4)
	gridWith7Ships, _ := PlaceShip(gridWith6Ships, 5, 1)
	gridWith8Ships, _ := PlaceShip(gridWith7Ships, 1, 3)
	gridWith9Ships, _ := PlaceShip(gridWith8Ships, 2, 4)

	gridWith1SunkShip, _, _ := shootOpponent(gridWith9Ships, 1, 2)
	gridWith2SunkShips, _, _ := shootOpponent(gridWith1SunkShip, 2, 3)
	gridWith3SunkShips, _, _ := shootOpponent(gridWith2SunkShips, 3, 4)
	gridWith4SunkShips, _, _ := shootOpponent(gridWith3SunkShips, 4, 5)
	gridWith5SunkShips, _, _ := shootOpponent(gridWith4SunkShips, 5, 6)
	gridWith6SunkShips, _, _ := shootOpponent(gridWith5SunkShips, 6, 4)
	gridWith7SunkShips, _, _ := shootOpponent(gridWith6SunkShips, 5, 1)
	gridWith8SunkShips, _, _ := shootOpponent(gridWith7SunkShips, 1, 3)
	gridWith9SunkShips, _, _ := shootOpponent(gridWith8SunkShips, 2, 4)

	//Act
	got := HasPlayerWon(gridWith9SunkShips)

	//Assert
	want := true
	if got != want {
		t.Errorf("player should of won, wanted %v got %v", want, got)
	}
}

func TestHasPlayerNotWon(t *testing.T) {
	//Arrange (set things up)
	grid := CreateGrid()
	gridWith1Ship, _ := PlaceShip(grid, 1, 2)
	gridWith2Ships, _ := PlaceShip(gridWith1Ship, 2, 3)
	gridWith3Ships, _ := PlaceShip(gridWith2Ships, 3, 4)
	gridWith4Ships, _ := PlaceShip(gridWith3Ships, 4, 5)
	gridWith5Ships, _ := PlaceShip(gridWith4Ships, 5, 6)
	gridWith6Ships, _ := PlaceShip(gridWith5Ships, 6, 4)
	gridWith7Ships, _ := PlaceShip(gridWith6Ships, 5, 1)
	gridWith8Ships, _ := PlaceShip(gridWith7Ships, 1, 3)
	gridWith9Ships, _ := PlaceShip(gridWith8Ships, 2, 4)

	gridWith1SunkShip, _, _ := shootOpponent(gridWith9Ships, 1, 2)
	gridWith2SunkShips, _, _ := shootOpponent(gridWith1SunkShip, 2, 3)
	gridWith3SunkShips, _, _ := shootOpponent(gridWith2SunkShips, 3, 4)
	gridWith4SunkShips, _, _ := shootOpponent(gridWith3SunkShips, 4, 5)
	gridWith5SunkShips, _, _ := shootOpponent(gridWith4SunkShips, 5, 6)
	gridWith6SunkShips, _, _ := shootOpponent(gridWith5SunkShips, 6, 4)
	gridWith7SunkShips, _, _ := shootOpponent(gridWith6SunkShips, 5, 1)
	gridWith8SunkShips, _, _ := shootOpponent(gridWith7SunkShips, 1, 3)

	//Act
	got := HasPlayerWon(gridWith8SunkShips)

	//Assert
	want := false
	if got != want {
		t.Errorf("wanted %v got %v", want, got)
	}
}

func TestPlayerHitsShipOnTurnAndReportsHit(t *testing.T) {
	//Arrange
	player := 1
	grid := CreateGrid()
	gridWith1Ship, _ := PlaceShip(grid, 1, 2)

	//Act
	_, got, _, _ := CurrentPlayerTakeShot(player, gridWith1Ship, 1, 2)

	//Assert
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPlayerMissesShipOnTurnAndReportsMiss(t *testing.T) {
	//Arrange
	row := 1
	col := 4
	player := 1
	grid := CreateGrid()

	//Act
	_, got, _, _ := CurrentPlayerTakeShot(player, grid, row, col)

	//Assert
	want := false
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPlayerMissesGridOnTurnAndReportsInvalidShot(t *testing.T) {
	//Arrange
	player := 1
	grid := CreateGrid()

	//Act
	_, got, _, _ := CurrentPlayerTakeShot(player, grid, -1, 4)

	//Assert
	want := false
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTurnDoesntChangePlayerAfterInvalidShot(t *testing.T) {
	//Arrange
	player := 1
	row := -1
	col := 4
	grid := CreateGrid()

	//Act
	got, _, _, _ := CurrentPlayerTakeShot(player, grid, row, col)

	//Assert
	want := 1
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTurnDoesChangeFromPlayer1ToPlayer2AfterMissedShot(t *testing.T) {
	//Arrange
	player := 1
	grid := CreateGrid()
	gridWith1Ship, _ := PlaceShip(grid, 1, 2)

	//Act
	got, _, _, _ := CurrentPlayerTakeShot(player, gridWith1Ship, 3, 5)

	//Assert
	want := 2
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTurnDoesChangeFromPlayer2ToPlayer1AfterMissedShot(t *testing.T) {
	//Arrange
	player := 2
	grid := CreateGrid()
	gridWith1Ship, _ := PlaceShip(grid, 1, 2)

	//Act
	got, _, _, _ := CurrentPlayerTakeShot(player, gridWith1Ship, 3, 5)

	//Assert
	want := 1
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTurnDoesChangeFromPlayer2ToPlayer1AfterHitShot(t *testing.T) {
	//Arrange
	player := 2
	grid := CreateGrid()
	gridWith1Ship, _ := PlaceShip(grid, 1, 2)

	//Act
	got, _, _, _ := CurrentPlayerTakeShot(player, gridWith1Ship, 3, 5)

	//Assert
	want := 1
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTurnDoesChangeFromPlayer1ToPlayer2AfterHitShot(t *testing.T) {
	//Arrange
	player := 1
	grid := CreateGrid()
	gridWith1Ship, _ := PlaceShip(grid, 1, 2)

	//Act
	got, _, _, _ := CurrentPlayerTakeShot(player, gridWith1Ship, 3, 5)

	//Assert
	want := 2
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPlayerHitsShipOnTurnAndReportsWin(t *testing.T) {
	//Arrange
	player := 1
	grid := CreateGrid()
	gridWith1Ship, _ := PlaceShip(grid, 1, 2)
	gridWith2Ships, _ := PlaceShip(gridWith1Ship, 2, 3)
	gridWith3Ships, _ := PlaceShip(gridWith2Ships, 3, 4)
	gridWith4Ships, _ := PlaceShip(gridWith3Ships, 4, 5)
	gridWith5Ships, _ := PlaceShip(gridWith4Ships, 5, 6)
	gridWith6Ships, _ := PlaceShip(gridWith5Ships, 6, 4)
	gridWith7Ships, _ := PlaceShip(gridWith6Ships, 5, 1)
	gridWith8Ships, _ := PlaceShip(gridWith7Ships, 1, 3)
	gridWith9Ships, _ := PlaceShip(gridWith8Ships, 2, 4)

	gridWith1SunkShip, _, _ := shootOpponent(gridWith9Ships, 1, 2)
	gridWith2SunkShips, _, _ := shootOpponent(gridWith1SunkShip, 2, 3)
	gridWith3SunkShips, _, _ := shootOpponent(gridWith2SunkShips, 3, 4)
	gridWith4SunkShips, _, _ := shootOpponent(gridWith3SunkShips, 4, 5)
	gridWith5SunkShips, _, _ := shootOpponent(gridWith4SunkShips, 5, 6)
	gridWith6SunkShips, _, _ := shootOpponent(gridWith5SunkShips, 6, 4)
	gridWith7SunkShips, _, _ := shootOpponent(gridWith6SunkShips, 5, 1)
	gridWith8SunkShips, _, _ := shootOpponent(gridWith7SunkShips, 1, 3)

	//Act
	_, _, got, _ := CurrentPlayerTakeShot(player, gridWith8SunkShips, 2, 4)

	//Assert
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPlayerHitsShipOnTurnAndDoesNotReportWin(t *testing.T) {
	//Arrange
	player := 1
	grid := CreateGrid()
	gridWithShip, _ := PlaceShip(grid, 2, 4)
	gridWith2Ships, _ := PlaceShip(gridWithShip, 1, 3)

	//Act
	_, _, got, _ := CurrentPlayerTakeShot(player, gridWith2Ships, 2, 4)

	//Assert
	want := false
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPlayerMissesShipWithValidShotAndDoesNotReportWin(t *testing.T) {
	//Arrange
	player := 1
	grid := CreateGrid()
	gridWith1Ship, _ := PlaceShip(grid, 1, 2)
	gridWith2Ships, _ := PlaceShip(gridWith1Ship, 2, 3)
	gridWith3Ships, _ := PlaceShip(gridWith2Ships, 3, 4)
	gridWith4Ships, _ := PlaceShip(gridWith3Ships, 4, 5)
	gridWith5Ships, _ := PlaceShip(gridWith4Ships, 5, 6)
	gridWith6Ships, _ := PlaceShip(gridWith5Ships, 6, 4)
	gridWith7Ships, _ := PlaceShip(gridWith6Ships, 5, 1)
	gridWith8Ships, _ := PlaceShip(gridWith7Ships, 1, 3)
	gridWith9Ships, _ := PlaceShip(gridWith8Ships, 2, 4)

	gridWith1SunkShip, _, _ := shootOpponent(gridWith9Ships, 1, 2)
	gridWith2SunkShips, _, _ := shootOpponent(gridWith1SunkShip, 2, 3)
	gridWith3SunkShips, _, _ := shootOpponent(gridWith2SunkShips, 3, 4)
	gridWith4SunkShips, _, _ := shootOpponent(gridWith3SunkShips, 4, 5)
	gridWith5SunkShips, _, _ := shootOpponent(gridWith4SunkShips, 5, 6)
	gridWith6SunkShips, _, _ := shootOpponent(gridWith5SunkShips, 6, 4)
	gridWith7SunkShips, _, _ := shootOpponent(gridWith6SunkShips, 5, 1)
	gridWith8SunkShips, _, _ := shootOpponent(gridWith7SunkShips, 1, 3)

	//Act
	_, _, got, _ := CurrentPlayerTakeShot(player, gridWith8SunkShips, 2, 6)

	//Assert
	want := false
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPlayerDoesNotWinGameWithInvalidShotOnTurn(t *testing.T) {
	//Arrange
	player := 1
	grid := CreateGrid()
	gridWith1Ship, _ := PlaceShip(grid, 1, 2)
	gridWith2Ships, _ := PlaceShip(gridWith1Ship, 2, 3)
	gridWith3Ships, _ := PlaceShip(gridWith2Ships, 3, 4)
	gridWith4Ships, _ := PlaceShip(gridWith3Ships, 4, 5)
	gridWith5Ships, _ := PlaceShip(gridWith4Ships, 5, 6)
	gridWith6Ships, _ := PlaceShip(gridWith5Ships, 6, 4)
	gridWith7Ships, _ := PlaceShip(gridWith6Ships, 5, 1)
	gridWith8Ships, _ := PlaceShip(gridWith7Ships, 1, 3)
	gridWith9Ships, _ := PlaceShip(gridWith8Ships, 2, 4)

	gridWith1SunkShip, _, _ := shootOpponent(gridWith9Ships, 1, 2)
	gridWith2SunkShips, _, _ := shootOpponent(gridWith1SunkShip, 2, 3)
	gridWith3SunkShips, _, _ := shootOpponent(gridWith2SunkShips, 3, 4)
	gridWith4SunkShips, _, _ := shootOpponent(gridWith3SunkShips, 4, 5)
	gridWith5SunkShips, _, _ := shootOpponent(gridWith4SunkShips, 5, 6)
	gridWith6SunkShips, _, _ := shootOpponent(gridWith5SunkShips, 6, 4)
	gridWith7SunkShips, _, _ := shootOpponent(gridWith6SunkShips, 5, 1)
	gridWith8SunkShips, _, _ := shootOpponent(gridWith7SunkShips, 1, 3)

	//Act
	_, _, got, _ := CurrentPlayerTakeShot(player, gridWith8SunkShips, -1, 4)

	//Assert
	want := false
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReportInvalidShotCoordinatesErrorToUserOnTurn(t *testing.T) {
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
		_, _, _, got := CurrentPlayerTakeShot(1, grid, coordinates.row, coordinates.col)

		//assert
		want := errors.New(coordinates.errorText)
		if got.Error() != want.Error() {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestValidMissedShotReturnsNoErrorOnTurn(t *testing.T) {
	type coordinates struct {
		row int
		col int
	}

	shotCoordinates := []coordinates{
		{row: 0, col: 6},
		{row: 0, col: 0},
		{row: 6, col: 0},
		{row: 6, col: 6},
	}

	for _, coordinates := range shotCoordinates {
		//arrange
		grid := CreateGrid()

		//act
		_, _, _, got := CurrentPlayerTakeShot(1, grid, coordinates.row, coordinates.col)

		//assert
		if got != nil {
			t.Errorf("got %v, want no error", got)
		}
	}
}

func TestValidHitShotReturnsNoErrorOnTurn(t *testing.T) {
	type coordinates struct {
		row int
		col int
	}

	shotCoordinates := []coordinates{
		{row: 0, col: 6},
		{row: 0, col: 0},
		{row: 6, col: 0},
		{row: 6, col: 6},
	}

	for _, coordinates := range shotCoordinates {
		//arrange
		grid := CreateGrid()
		gridWith1Ship, _ := PlaceShip(grid, 0, 6)
		gridWith2Ships, _ := PlaceShip(gridWith1Ship, 0, 0)
		gridWith3Ships, _ := PlaceShip(gridWith2Ships, 6, 6)
		gridWith4Ships, _ := PlaceShip(gridWith3Ships, 6, 0)

		//act
		_, _, _, got := CurrentPlayerTakeShot(1, gridWith4Ships, coordinates.row, coordinates.col)

		//assert
		if got != nil {
			t.Errorf("got %v, want no error", got)
		}
	}
}
