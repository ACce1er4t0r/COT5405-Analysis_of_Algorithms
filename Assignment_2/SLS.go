package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/xcltapestry/xclpkg/algorithm"
)

type points [][]float64

func (p points) Len() int {
	return len(p)
}

func (p points) Less(i, j int) bool {
	for k := 0; k < len(p[i]); k++ {
		if p[i][k] < p[j][k] {
			return true
		} else if p[i][k] == p[j][k] {
			continue
		} else {
			return false
		}
	}
	return true
}

func (p points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func SLS(ps points, penalty float64) ([]string, float64) {
	var result []string
	n := len(ps)
	t := []float64{-1, -1}
	ps = append(ps, t)
	sort.Sort(ps)
	cumulativeX := make([]float64, n+1)
	cumulativeY := make([]float64, n+1)
	cumulativeXY := make([]float64, n+1)
	cumulativeXsqr := make([]float64, n+1)
	OPT := make([]float64, n+1)
	optSegment := make([]int, n+1)
	cumulativeX[0], cumulativeY[0], cumulativeXY[0], cumulativeXsqr[0] = 0, 0, 0, 0
	e := make([][]float64, n+1)
	slope := make([][]float64, n+1)
	intercept := make([][]float64, n+1)
	for i := 0; i < n+1; i++ {
		e[i] = make([]float64, n+1)
		slope[i] = make([]float64, n+1)
		intercept[i] = make([]float64, n+1)
	}
	for j := 1; j <= n; j++ {
		cumulativeX[j] = cumulativeX[j-1] + ps[j][0]
		cumulativeY[j] = cumulativeY[j-1] + ps[j][1]
		cumulativeXY[j] = cumulativeXY[j-1] + ps[j][0]*ps[j][1]
		cumulativeXsqr[j] = cumulativeXsqr[j-1] + ps[j][0]*ps[j][0]
		for i := 1; i <= j; i++ {
			interval := float64(j - i + 1)
			xSum := cumulativeX[j] - cumulativeX[i-1]
			ySum := cumulativeY[j] - cumulativeY[i-1]
			xySum := cumulativeXY[j] - cumulativeXY[i-1]
			xsqrSum := cumulativeXsqr[j] - cumulativeXsqr[i-1]

			num := interval*xySum - xSum*ySum
			if num == 0 {
				slope[i][j] = 0.0
			} else {
				denom := interval*xsqrSum - xSum*xSum
				if denom == 0 {
					slope[i][j] = math.MaxFloat64
				} else {
					slope[i][j] = num / denom
				}
			}
			intercept[i][j] = (ySum - slope[i][j]*xSum) / interval
			for k := i; k <= j; k++ {
				e[i][j] = 0
				tmp := ps[k][1] - slope[i][j]*ps[k][0] - intercept[i][j]
				e[i][j] += tmp * tmp
			}
		}
	}
	OPT[0], optSegment[0] = 0, 0
	for j := 1; j <= n; j++ {
		mn, k := math.MaxFloat64, 0
		for i := 1; i <= j; i++ {
			tmp := e[i][j] + OPT[i-1]
			if tmp < mn {
				mn = tmp
				k = i
			}
		}
		OPT[j] = mn + penalty
		optSegment[j] = k
	}

	stack := algorithm.NewStack()
	i := n
	j := optSegment[n]
	for i > 0 {
		stack.Push(i)
		stack.Push(j)
		i = j - 1
		j = optSegment[i]
	}
	for !stack.Empty() {
		x := stack.Top()
		err := stack.Pop()
		if err != nil {
			fmt.Printf("%s", err)
			return nil, 0
		}
		y := stack.Top()
		err = stack.Pop()
		if err != nil {
			fmt.Printf("%s", err)
			return nil, 0
		}
		i, _ = x.(int)
		j, _ = y.(int)
		result = append(result, fmt.Sprintf("Segment (y = %f * x + %f) from points %d to %d with square error %f.\n", slope[i][j], intercept[i][j], x, y, e[i][j]))
	}
	return result, OPT[n]
}
