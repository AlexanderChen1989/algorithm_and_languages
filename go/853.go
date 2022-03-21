package main

import "sort"

func carFleet(target int, position []int, speed []int) int {
	pairs := make([]Pair, len(position))

	for i := 0; i < len(position); i++ {
		pairs[i].Pos = position[i]
		pairs[i].Speed = speed[i]
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Pos > pairs[j].Pos })

	fleets := pairs[:1]

	for i := 1; i < len(pairs); i++ {
		p := pairs[i]
		t := fleets[len(fleets)-1]
		if float64(target-p.Pos)/float64(p.Speed) > float64(target-t.Pos)/float64(t.Speed) {
			fleets = append(fleets, p)
		}
	}

	return len(fleets)
}

type Pair struct {
	Pos   int
	Speed int
}