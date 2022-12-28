package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/soniakeys/graph"
)

func main() {
	Test()
	Test2()
}

/** Question 1
 *
 */
func Test() {
	nodes := 10000000
	edges := 20000000
	g := generateUndirected(nodes, edges)
	edgeCollect := generateEdges(g)
	start := time.Now()
	FindCycleDSU(edgeCollect)
	elapsed := time.Since(start)
	fmt.Println("total cost:", elapsed)
}

/** Question 2
 *
 */
func Test2() {
	nodes := 10
	edges := 18
	if edges < nodes-1 {
		fmt.Println("too less edges")
		os.Exit(-1)
	}
	g := generateConnectGraph(nodes, edges)
	for fr, to := range g.AdjacencyList {
		for _, to := range to {
			if graph.NI(fr) < to {
				fmt.Println(fr, "-", to)
			}
		}
	}
	edgeCollect := generateWEdges(g, edges)
	fmt.Println(edgeCollect)
	sort.Slice(edgeCollect[:edges], func(i, j int) bool {
		return edgeCollect[i][2] < edgeCollect[j][2]
	})
	fmt.Println(edgeCollect)
	start := time.Now()
	kruskal(nodes, edges, edgeCollect)
	elapsed := time.Since(start)
	fmt.Println("total cost:", elapsed)

}

func find(parent []int, f int) int {
	if f == parent[f] {
		return f
	}
	tmp := find(parent, parent[f])
	parent[f] = tmp
	return tmp
}
