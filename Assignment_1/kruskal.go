package main

import (
	rand2 "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"

	"github.com/soniakeys/graph"
)

func generateConnectGraph(nnum int, enum int) graph.Undirected {
	n, _ := rand2.Int(rand2.Reader, big.NewInt(65536))
	r := rand.New(rand.NewSource(n.Int64()))
	g := graph.GnmUndirected(nnum, enum, r)
	if !g.IsConnected() {
		generateConnectGraph(nnum, enum)
	}
	return g
}

func generateWEdges(g graph.Undirected, enum int) [][]int {
	weight := make([]int, enum)

	for i := 0; i < enum; i++ {
		n, _ := rand2.Int(rand2.Reader, big.NewInt(int64(20)))
		if n.Int64() == 0 {
			weight[i] = int(n.Int64()) + 1
		} else {
			weight[i] = int(n.Int64())
		}
	}
	edgeCollect := [][]int{}
	for fr, to := range g.AdjacencyList {
		for _, to := range to {
			if graph.NI(fr) < to {
				tmp := []int{}
				tmp = append(tmp, fr)
				tmp = append(tmp, int(to))
				edgeCollect = append(edgeCollect, tmp)
			}
		}
	}
	for i := 0; i < enum; i++ {
		edgeCollect[i] = append(edgeCollect[i], weight[i])
	}
	return edgeCollect
}

func kruskal(nnum int, enum int, edgeCollect [][]int) {
	parent := []int{}
	for i := 0; i <= nnum+1; i++ {
		parent = append(parent, i)
	}
	// O(n)
	res := 0

	for i := 0; i < enum; i++ {
		pf := find(parent, edgeCollect[i][0])
		pt := find(parent, edgeCollect[i][1])

		if pf != pt {
			parent[pt] = pf
			res += edgeCollect[i][2]
		}
	}
	//O(m)

	if res == -1 {
		fmt.Println("can't find MST")
	} else {
		fmt.Println("MST is:", res)
	}
	//return res
}
