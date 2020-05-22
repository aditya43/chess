## Chess
Simple program simulating movements of various types of pieces on the empty chessboard.

------

## Current Status
Done (Last Updated: 22 May 2020)

------

## 3rd Party Libraries Used
- `github.com/gookit/color`: For printing colorful text in terminal

------

## APIs
- **ChessBoard**
    * `CreateChessBoard()`: Create cells with their numeric identifiers and return a pointer to the `ChessBoard struct`.
    * `PlacePiece()`: Place piece on a chessboard and update available move positions for a piece.
    * `Print()`: Render chessboard with a piece.
- **Piece**
    * `CreatePiece()`: Create a piece and add available move positions.

------

## Tests
- **[Test Coverage Report (HTML) :rocket:](https://htmlpreview.github.io/?https://github.com/aditya43/chess/blob/master/test_coverage.html#file0)**
- **Unit Tests**
    ```diff
    // User Input | user_input_test.go
    + TestForSingleWordInput
    + TestForInvalidPieceNameInput
    + TestForInvalidPositionInput
    + TestForInvalidPieceTypeAndPositionInput
    + TestForCorrectUserInput
    + TestForCaseInsensitiveUserInput

    // Chessboard | chessboard_test.go
    + TestChessBoardHas64Cells
    + TestStrCellsInChessBoardRepresentsCorrectNumericPositionOnChessBoard
    + TestCreateChessBoardReturnsPointerToChessboardType
    + TestPlacePieceFuncPlacesPieceOnAChessBoard
    + TestPlacePieceFuncUpdatesAvailableMovePositionsForPiece

    // Piece | pieces_test.go
    + TestCreatePieceReturnsPointerToPieceType
    + TestCreatePieceCreatesPiece
    + TestUpdateMovePositionsUpdatesAvailableMovePositionsForPiece
    + TestUpdateMovePositionsUpdateCorrectNumberOfMovePositionsForPiece
    + TestUpdateMovePositionsGenerateCorrectMovePositionsForPiece
    + TestMovePositionsAreGeneratedForDifferentTypesOfPieces

    // Piece Boundries | boundries_test.go
    + TestLeftBoundryIsSetForPieceIfItCanMoveOnXAxis
    + TestRightBoundryIsSetForPieceIfItCanMoveOnXAxis
    + TestCorrectLeftBoundryIsSetForPiece
    + TestCorrectRightBoundryIsSetForPiece
    + TestLeftBoundryIsNotSetForPieceIfItCannotMoveOnXAxis
    + TestRightBoundryIsNotSetForPieceIfItCannotMoveOnXAxis
    + TestTopBoundryIsSetForPieceIfItCanMoveOnYAxis
    + TestBottomBoundryIsSetForPieceIfItCanMoveOnYAxis
    + TestCorrectTopBoundryIsSetForPiece
    + TestCorrectBottomBoundryIsSetForPiece
    + TestTopBoundryIsNotSetForPieceIfItCannotMoveOnYAxis
    + TestBottomBoundryIsNotSetForPieceIfItCannotMoveOnYAxis
    + TestTopLeftBoundryIsSetForPieceIfItCanMoveInDiagonalDirection
    + TestTopRightBoundryIsSetForPieceIfItCanMoveInDiagonalDirection
    + TestCorrectTopLeftBoundryIsSetForPiece
    + TestCorrectTopRightBoundryIsSetForPiece
    + TestTopLeftBoundryIsNotSetForPieceIfItCannotMoveInDiagonalDirection
    + TestTopRightBoundryIsNotSetForPieceIfItCannotMoveInDiagonalDirection
    + TestBottomLeftBoundryIsSetForPieceIfItCanMoveInDiagonalDirection
    + TestBottomRightBoundryIsSetForPieceIfItCanMoveInDiagonalDirection
    + TestCorrectBottomLeftBoundryIsSetForPiece
    + TestCorrectBottomRightBoundryIsSetForPiece
    + TestBottomLeftBoundryIsNotSetForPieceIfItCannotMoveInDiagonalDirection
    + TestBottomRightBoundryIsNotSetForPieceIfItCannotMoveInDiagonalDirection
    + TestCorrectBoundriesAreSetForQueen

    // Move Positions For Piece | positions_test.go
    + TestCorrectMovePositionIsGeneratedForPawnTypePiece
    + TestMovePositionsAreGeneratedInTopDirectionForNonPawnTypePiece
    + TestMovePositionsAreGeneratedInBottomDirectionForNonPawnTypePiece
    + TestMovePositionsAreGeneratedInLeftDirectionForNonPawnTypePiece
    + TestMovePositionsAreGeneratedInRightDirectionForNonPawnTypePiece
    + TestCorrectMovePositionsAreGeneratedForDifferentTypesOfPieces
    ```

- **Benchmark Tests**
    ```diff
    // Benchmark test for validateInput() func
    - BenchmarkUserInputValidator

    // Benchmark test for CreateChessBoard() func
    - BenchmarkCreateChessBoard

    // Benchmark test for moving a piece on chessboard
    - BenchmarkMovePieceOnChessBoard

    // Benchmark create chessboard and create piece
    - BenchmarkCreateChessBoardAndPiece

    // Benchmark create chessboard, create piece and
    // then move piece on a created chessboard
    - BenchmarkCreateChessBoardAndPieceMoveItOnChessBoard
    ```

- **Benchmark Test Results**
    ```
    goos: darwin
    goarch: amd64
    pkg: github.com/aditya43/chess
    BenchmarkUserInputValidator
    BenchmarkUserInputValidator-8                           	30921040	        36.6 ns/op	       0 B/op	       0 allocs/op
    BenchmarkCreateChessBoard
    BenchmarkCreateChessBoard-8                             	   49857	     22715 ns/op	   18126 B/op	     212 allocs/op
    BenchmarkMovePieceOnChessBoard
    BenchmarkMovePieceOnChessBoard-8                        	 5746688	       200 ns/op	     336 B/op	       4 allocs/op
    BenchmarkCreateChessBoardAndPiece
    BenchmarkCreateChessBoardAndPiece-8                     	   50523	     22996 ns/op	   18270 B/op	     213 allocs/op
    BenchmarkCreateChessBoardAndPieceMoveItOnChessBoard
    BenchmarkCreateChessBoardAndPieceMoveItOnChessBoard-8   	   50686	     23099 ns/op	   18607 B/op	     217 allocs/op
    ```

------

## Implementation Flow (Developer Notes)
> This will be modified as the process goes from rough draft to final implementation.
- **User Interaction:**
    - User inputs will be case-insensitive.
    - Input piece name and piece position.
    - Type `exit` to quit program.
- Create 8x8 chessboard with position acronyms (a1-a8)..(h1-h8).
- Chessboard struct:
    ```go
    type ChessBoard struct {
        Cells    map[int]string // Numeric Position to Text Representation. For e.g. 10 -> B2
        Pieces   map[int]*Piece // Numeric Position --> Piece. For printing purposes only
        StrCells map[string]int // Text Representation to Numeric Position of a cell
    }
    ```
- Piece struct:
    ```go
    type Piece struct {
        Kind   string // Kind of piece. e.g. king, queen, horse, pawn, rook, bishop
        CurPos int    // Numeric position on board
        Symbol string // e.g. ♚, ♛, ♜, ♝, ♞

        isHorse bool // Is current piece a Horse
        isPawn  bool // Is current piece a Pawn

        xAllow     bool // Is piece allowed to traverse in X direction
        yAllow     bool // Is piece allowed to traverse in Y direction
        crossAllow bool // Is piece allowed to traverse in diagonal direction

        maxX     int // Max number of cells a piece can traverse in X direction
        maxY     int // Max number of cells a piece can traverse in Y direction
        maxCross int // Max number of cells a piece can traverse in diagonal direction

        AvailPos map[int]bool // Available move positions for a piece

        xBoundry        XBoundry        // Horizontal boundries
        yBoundry        YBoundry        // Vertical boundries
        diagonalBoundry DiagonalBoundry // Diagonal boundries
    }
    ```
- Pieces:
    * King
        - Can move only 1 step at a time in all 8 directions (horizontal, vertical
and diagonal)
    * Queen
        - Can move across the board in all 8 directions
    * Horse
        - Can move across the board only in 2.5 steps (2 vertical steps and 1
horizontal step)
    * Rook
        - Can move across the board only vertically and horizontally
    * Pawn
        - Can move only 1 step at a time, in the forward direction, vertically.
Can also move 1 step forward diagonally, in order to eliminate an opposing
piece.
        - **Since there is only 1 side (as in Black/White), diagonal traversal functionality for a pawn piece will be ommitted.**
    * Bishop
        - Can move across the board only diagonally
- Skeleton for adding move positions for various types of pieces:
    ```go
    // Calculate and add available positions for a piece based off it's current position
    func (p *Piece) updateMovePositions(b *ChessBoard) {
        p.availPos = make(map[int]bool)

        if !p.isHorse {
            if p.yAllow { // Check if a piece can move on y axis
                if p.isPawn {
                    // Y axis
                    // Get positions in top direction for a pawn piece
                }

                if !p.isPawn {
                    // Y axis
                    // Get positions in top direction for a non pawn piece

                    // Y axis
                    // Get positions in bottom direction for a non pawn piece

                }
            }

            // Check if a piece can move on x axis
            if p.xAllow {
                // X axis
                // Get positions in right direction for a non pawn piece

                // X axis
                // Get positions in left direction for a non pawn piece
            }

            // Check if a piece can move in diagonal direction
            if p.crossAllow {
                // Diagonal
                // Get diagonal move positions in top-right direction for a non pawn piece

                // Diagonal
                // Get diagonal move positions in top-left direction for a non pawn piece

                // Diagonal
                // Get diagonal move positions in bottom-right direction for a non pawn piece

                // Diagonal
                // Get diagonal move positions in bottom-left direction for a non pawn piece
            }
        }

        if p.isHorse {
            // Add move positions for a horse piece
        }
    }
    ```
- To calculate move position in various directions from the current position of the piece:
![Move Positions Calculations](/images/move-positions.jpg)
- To generate test coverage:
    ```sh
    go test ./... -coverprofile=test_coverage.txt
    go tool cover -html=test_coverage.txt -o test_coverage.html
    ```

------

## License
Open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT).