package midas

import (
	"math"
	"math/rand"
)

type EdgeHash struct {
	numRows    int
	numBuckets int
	m          int
	hashA      []int
	hashB      []int
	count      [][]float64
}

func randomIntSlice(min int, max int, size int) []int {
	a := make([]int, size)
	for i, _ := range a {
		a[i] = rand.Intn(max-min) + min
	}
	return a
}

func NewEdgeHash(numRows int, numBuckets int, m int) *EdgeHash {
	count := make([][]float64, numRows)
	for i, _ := range count {
		count[i] = make([]float64, numBuckets)
	}

	return &EdgeHash{
		numRows:    numRows,
		numBuckets: numBuckets,
		m:          m,
		hashA:      randomIntSlice(1, numBuckets, numRows),
		hashB:      randomIntSlice(0, numBuckets, numRows),
		count:      count,
	}
}

func (e *EdgeHash) Hash(a int, b int, i int) int {
	resid := ((a+e.m*b)*e.hashA[i] + e.hashB[i]) % e.numBuckets
	if resid < 0 {
		return resid + e.numBuckets
	} else {
		return resid
	}
}

func (e *EdgeHash) Insert(a int, b int, weight float64) {
	for i := 0; i < e.numRows; i++ {
		bucket := e.Hash(a, b, i)
		e.count[i][bucket] += weight
	}
}

func (e *EdgeHash) GetCount(a int, b int) float64 {
	bucket := e.Hash(a, b, 0)
	minCount := e.count[0][bucket]
	for i := 1; i < e.numRows; i++ {
		bucket = e.Hash(a, b, i)
		minCount = math.Min(minCount, e.count[i][bucket])
	}

	return minCount
}

func (e *EdgeHash) Clear() {
	for i, row := range e.count {
		for j, _ := range row {
			e.count[i][j] = 0
		}
	}
}

func (e *EdgeHash) Lower(factor float64) {
	for i, row := range e.count {
		for j, _ := range row {
			e.count[i][j] *= factor
		}
	}
}
