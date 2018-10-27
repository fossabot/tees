// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package main
package walk_test

import (
	"github.com/GoLangsam/tees/list"

	"github.com/GoLangsam/tees/list/math"

	"github.com/GoLangsam/tees/list/walk"
	"github.com/GoLangsam/tees/list/walk/move"

	"fmt"
)

var testakass = []struct {
	akas walk.Akas
	name string
}{
	{move.ChessRook, "Chess Rook"},
	{move.ChessKnight, "Chess Knight"},
	{move.ChessBishop, "Chess Bishop"},
	{move.ChessQueen, "Chess Queen"},
}

// func main() {
func ExampleAkas_PrintWalker() {
	// Create a new list
	var chess = list.NewList("Chess")

	// Create new lists
	var cols = chess.AddBeam("Spalten", "A", "B", "C", "D", "E")
	var rows = chess.AddBeam("Zeilen", 1, 2, 3, 4, 5)
	fmt.Println("Starting")

	var board = math.Xross(cols, rows)
	board.PrintAways("Schach")
	board.Root().Away().List().PrintAways("Brett")
	var s = board.Front().Away().List().Front()         // A|1
	var m = s.Next().Next().Away().Next().Next().Away() // C|3
	var e = board.Back().Away().List().Back()           // E|5
	if s.IsRoot() || m.IsRoot() || e.IsRoot() {
	}

	fmt.Println("Akas:")
	for _, jump := range testakass {
		jump.akas.PrintWalker(jump.name, m)
	}
	// Output:
	// Starting
	// Schach: List=Spalten | Total=5
	// => : List=A	A|1 | A|2 | A|3 | A|4 | A|5 | Total=5
	// => : List=B	B|1 | B|2 | B|3 | B|4 | B|5 | Total=5
	// => : List=C	C|1 | C|2 | C|3 | C|4 | C|5 | Total=5
	// => : List=D	D|1 | D|2 | D|3 | D|4 | D|5 | Total=5
	// => : List=E	E|1 | E|2 | E|3 | E|4 | E|5 | Total=5
	// Brett: List=Zeilen | Total=5
	// => : List=1	1|A | 1|B | 1|C | 1|D | 1|E | Total=5
	// => : List=2	2|A | 2|B | 2|C | 2|D | 2|E | Total=5
	// => : List=3	3|A | 3|B | 3|C | 3|D | 3|E | Total=5
	// => : List=4	4|A | 4|B | 4|C | 4|D | 4|E | Total=5
	// => : List=5	5|A | 5|B | 5|C | 5|D | 5|E | Total=5
	// Akas:
	// Chess Rook	=>: C|3 ->: C|2 ->: C|1 ->: C|4 ->: C|5 ->: B|3 ->: A|3 ->: D|3 ->: E|3
	// Chess Knight	=>: C|3 ->: B|5 ->: D|5 ->: B|1 ->: D|1 ->: A|2 ->: A|4 ->: E|2 ->: E|4
	// Chess Bishop	=>: C|3 ->: B|4 ->: A|5 ->: D|4 ->: E|5 ->: B|2 ->: A|1 ->: D|2 ->: E|1
	// Chess Queen	=>: C|3 ->: C|2 ->: C|1 ->: C|4 ->: C|5 ->: B|3 ->: A|3 ->: D|3 ->: E|3 ->: B|4 ->: A|5 ->: D|4 ->: E|5 ->: B|2 ->: A|1 ->: D|2 ->: E|1
}