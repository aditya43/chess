## Chess
Simple program simulating movements of various types of pieces on the empty chessboard.

## Current Status
WIP (Work In Progress)!

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
        cells    map[int]string // Numeric Position to Text Representation. For e.g. 10 -> B2
        pieces   map[int]*Piece // Numeric Position --> Piece. For printing purposes only
        strCells map[string]int // Text Representation to Numeric Position of a cell
    }
    ```
- Piece struct:
    ```go
    type Piece struct {
        name            string          // Name. e.g. king, queen, horse, pawn, rook, bishop
        curPos          int             // Numeric position on board
        symbol          string          // e.g. ♚, ♛, ♜, ♝, ♞
        isHorse         bool            // Is current piece a Horse
        isPawn          bool            // Is current piece a Pawn
        xAllow          bool            // Is piece allowed to traverse in X direction
        yAllow          bool            // Is piece allowed to traverse in Y direction
        crossAllow      bool            // Is piece allowed to traverse in diagonal direction
        maxX            int             // Max number of cells a piece can traverse in X direction
        maxY            int             // Max number of cells a piece can traverse in Y direction
        maxCross        int             // Max number of cells a piece can traverse in diagonal direction
        availPos        map[int]bool    // Available move positions for a piece
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
        - **Since there is only 1 side (as in Black/White), diagonal traversal functionality will be ommitted.**
    * Bishop
        - Can move across the board only diagonally

# APIs
- **ChessBoard**
    * `CreateChessBoard()`: Create cells with their numeric identifiers and return a pointer to the `ChessBoard struct`.
    * `placePiece()`: Place piece on a chessboard and update available move positions for a piece.
    * `print()`: Render chessboard with a piece.
- **Piece**
    * `CreatePiece()`: Create a piece.
    * `updatePositions()`: Calculate and add available move positions for a piece based off it's current position.

# Test Cases
- Is the chessboard created.
- User inputs:
    * Valid piece name is entered.
    * Valid position is entered.
- Can `X` piece move at `Y` location.
- Get current positions of all pieces.

## License
Open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT).