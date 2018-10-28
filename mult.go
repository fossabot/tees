// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

import (
	"github.com/GoLangsam/tees/list"
)

type Calcer interface {
	CanIter
	CVs() *list.ComposedValue
	With(*list.List) *list.ComposedValue
}

// Times returns a new list: the cross product of l with some lists...
// ( recursively as [[[ l * l ] * l ] ... ] )
// Note: Times( l, nil ) returns a new empty list
// the root of which carries the CVs of the original l.Root()
// and the elements carry the CVs of the original elements
// Note: The Away's in the new list point to nil - thus, the new list is isolated.
func Times(l Calcer, lists ...*list.List) *list.List {
	n := len(lists)
	switch {
	case n == 0:
		return times(l, nil)
	case n == 1:
		return times(l, lists[0])
	default:
		return times(l, Times(lists[0], lists[1:]...))
	}
}

// ===========================================================================

// times returns a new list with len(X) * len(Y) Elements
// representing the cross-product of the list X * Y
// Note: l.times( nil ) returns a new list with no elements
// Note: The Away's in the new list point to nil - thus, the new list is isolated.
func times(X Calcer, Y *list.List) *list.List {
	if X == nil {
		return New(nil)
	}
	newl := New(X.CVs())
	if Y != nil {
		for x := X.Front(); x != nil; x = x.Next() {
			for y := Y.Front(); y != nil; y = y.Next() {
				newl.PushBack(x.With(y))
			}
		}
		newl.Root().Value = X.With(Y)
	}
	return newl
}