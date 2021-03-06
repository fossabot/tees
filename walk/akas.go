// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package walk

// Akas (Burmese) is a slice of forms of movement (in martial arts, and dances :-) )
type Akas []Kata // Burmese

// ========================================================

// From returns the trail of non-nil Heres reached from e by jumps
func (jumps Akas) From(e Here) (Trail, Distance) {
	var dist, dnow Distance
	var goal = e
	goals := make(Trail, 0, len(jumps))
	for _, steps := range jumps {
		goal, dnow = steps.From(e)
		if goal == nil {
			continue
		}
		goals = append(goals, goal)
		dist += dnow
	}
	return goals, dist
}

// ========================================================

// Grab returns all Heres reached from e by jumps
// Note: Grab may be useful in debugging, as it returns a full trace
// To Grab is not intended for regular use - Don't be greedy :-)
func (jumps Akas) Grab(e Here) (Trail, Distance) {
	var dist Distance
	goals := make(Trail, 0, len(jumps)*len(jumps))
	for _, steps := range jumps {
		goal, dnow := steps.Grab(e)
		if goal == nil || len(goal) == 0 {
			continue
		}
		goals = append(goals, goal...)
		dist += dnow
	}
	return goals, dist
}

// ========================================================

// Haul returns the Heres (or nil) From e by hauling jumps
// Note: From any new goal, just the current Kata is repeated!
// Not all jumps are done again - this would imply loops.
func (jumps Akas) Haul(e Here) (Trail, Distance) {
	var dist Distance
	goals := make(Trail, 0, len(jumps)*len(jumps)*8)
	for _, steps := range jumps {
		goal, dnow := steps.Haul(e)
		if goal == nil || len(goal) == 0 {
			continue
		}
		goals = append(goals, goal...)
		dist += dnow
	}
	return goals, dist
}

// ========================================================
// Iterator

// Walker returns an iterator walking all Kata.From(e) ...
func (jumps Akas) Walker(e Here) Walk {

	var curr = e
	var akas = jumps
	var maxi = len(akas) - 1
	var aidx int // index of akas
	var next = akas[aidx].Walker(curr)

	var move Walk = func() Here {
	next:
		goal := next()
		if goal == nil && aidx < maxi {
			aidx++
			next = akas[aidx].Walker(curr)
			goto next
		}
		return goal
	}
	return move
}
