// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"github.com/GoLangsam/tees/list"
)

// NQueensR returns 'only' the rows
func NQueensR(N int) *list.List {

	var r = list.NewList("Ranks")
	var f = list.NewList("Files")
	var a = list.NewList("DiagA")
	var b = list.NewList("DiagB")

	for _, id := range IDs("R-", N) {
		r.AddBeam(id)
	}
	for _, id := range IDs("F-", N) {
		f.AddBeam(id)
	}
	for _, id := range IDs("A-", 2*N-1) {
		a.AddBeam(id)
	}
	for _, id := range IDs("B-", 2*N-1) {
		b.AddBeam(id)
	}

	// Note: not really relevant, just for clarification :-)
	//	var p = list.NewList( "Primary" )
	var s = list.NewList("Secondary")

	//	p.AddList(r)
	s.AddList(f)
	s.AddList(a)
	s.AddList(b)

	// for index-calculations
	var rs = r.Elements()
	var fs = f.Elements()

	var as = a.Elements()
	var bs = b.Elements()

	var ae, be *list.Element

	var rows = list.NewList("Rows")
	var rowi = 0
	for fi, fe := range fs {
		for ri, re := range rs {
			//			fmt.Println( fi, ri, fi+ri, N-1-fi+ri)
			ae = as[fi+ri]
			be = bs[N-1-fi+ri]
			rows.AddJunk(rowi, re.AwayList(), fe.AwayList(), ae.AwayList(), be.AwayList())
			rowi++
		}
	}
	rows.Join(r)
	return r
}
