package main

import (
	"fmt"
	"github.com/sruthijogi/Sudoku/board"
	"github.com/sruthijogi/Sudoku/puzzle"
	"testing"
)

// Command: go test

func Test_puzzleSetUp(t *testing.T) {
	input := board.Basic()
	newBoard := puzzle.New(input)
	b, _, _ := newBoard.Display()
	if b != board.Basic() {
		fmt.Println(b)
		t.Error("should equal: ", board.Basic())
	}
}

func Test_solve(t *testing.T) {
	input := board.Basic()
	newBoard := puzzle.New(input)
	newBoard.Solve()
	b, _, _ := newBoard.Display()
	if b != board.Solved() {
		fmt.Println(b)
		t.Error("should equal solved board: ", board.Solved())
	}
}

func Test_slowSolve(t *testing.T) {
	input := board.Basic()
	newBoard := puzzle.New(input)
	newBoard.Solve()
	b, _, _ := newBoard.Display()
	if b != board.Solved() {
		fmt.Println(b)
		t.Error("should equal solved board: ", board.Solved())
	}
}

func Test_validate(t *testing.T) {
	input := board.Basic()
	// import standard input
	newBoard := puzzle.New(input)
	newBoard.Solve()
	v := newBoard.Validate()
	if !v {
		fmt.Println(v)
		t.Error("should equal true")
	}
	// import unsolvable input
	input = board.Broken()
	newBoard.FillPuzzle(input)
	newBoard.Solve()
	v = newBoard.Validate()
	if v {
		fmt.Println(v)
		t.Error("should equal false")
	}
}
