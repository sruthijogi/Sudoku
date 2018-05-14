package main

import (
	"fmt"
	"github.com/sruthijogi/Sudoku/board"
	"github.com/sruthijogi/Sudoku/puzzle"
)

// Command: go run cli.go

func main() {
	i := board.Basic()
	sudoku := puzzle.New(i)
	sudoku.Grade()
	err := sudoku.SlowSolve()
	board, status, dif := sudoku.Display()
	if !err {
		fmt.Println(board, status, dif)
	}
}
