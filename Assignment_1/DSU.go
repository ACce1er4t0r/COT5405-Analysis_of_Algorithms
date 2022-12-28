package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/soniakeys/graph"
)

func generateUndirected(nnum int, enum int) graph.Undirected {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return graph.GnmUndirected(nnum, enum, r)
}

func generateEdges(g graph.Undirected) [][]int {
	edgeCollect := [][]int{}
	for fr, to := range g.AdjacencyList {
		for _, to := range to {
			if graph.NI(fr) < to {
				tmp := []int{}
				tmp = append(tmp, fr)
				tmp = append(tmp, int(to))
				edgeCollect = append(edgeCollect, tmp)
				//fmt.Println(fr, "-", to)
			}
		}
	}
	return edgeCollect
}

func FindCycleDSU(edgeCollect [][]int) bool {

	n := len(edgeCollect)
	if n < 3 {
		fmt.Println("no cycle")
		return false
	}

	parent := make([]int, n+1)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	for _, e := range edgeCollect {
		f, t := e[0], e[1]
		pf := find(parent, f)
		pt := find(parent, t)
		if pf == pt {
			fmt.Println("find cycle")
			return true
		}
		parent[pt] = pf
	}
	fmt.Println("no cycle")
	return false
}
