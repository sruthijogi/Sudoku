# Sudoku
Sudoku puzzle solver and generator - https://sujoku.herokuapp.com/   (Herokuapp.com is an open-source to create apps)

Overview
Sudoku web app: solves, generates, grades, and validates sudoku puzzles.

The algorithm implements two solving functions:

QuickFill - Called so because it quickly checks horizontally, vertically and in the nine grid box for possible options. If only one possible solution remains, it adds the value to the square.

Guess - It fills an empty square with a possible, non-conflicting value. If the puzzle is solved (completely filled), it validates to see if it correct. If not, it goes back and fills the square with another possible value. It recursively fills in and backtracks until the puzzle is complete.

SlowSolve is the faster function and therefore the one used for the API.

Technologies
Go
Gin
Semantic UI
jQuery
Heroku
Godep
API
Get
Board - returns a puzzle board

https://sujoku.herokuapp.com/board

Arguments -

Difficulty:
easy
medium
hard
random
Example:

https://sujoku.herokuapp.com/board?difficulty=easy

Solve - returns the solved puzzle, along with difficulty and status
Grade - returns the difficulty of the puzzle
Validate - returns the status of the puzzle

For more, check out the [api.js](https://github.com/sruthijogi/Sudoku/blob/master/public/api.js) file
