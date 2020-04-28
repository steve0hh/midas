package midas

import "math"

type NodeHash struct {
	numRows    int
	numBuckets int
	hashA      []int
	hashB      []int
	count      [][]float64
}

func NewNodeHash(numRows int, numBuckets int) *NodeHash {
	count := make([][]float64, numRows)

	for i, _ := range count {
		count[i] = make([]float64, numBuckets)
	}

	return &NodeHash{
		numRows:    numRows,
		numBuckets: numBuckets,
		hashA:      randomIntSlice(1, numBuckets, numRows),
		hashB:      randomIntSlice(0, numBuckets, numRows),
		count:      count,
	}
}

func (n *NodeHash) Hash(a int, i int) int {
	resid := (a*n.hashA[i] + n.hashB[i]) % n.numBuckets
	if resid < 0 {
		return resid + n.numBuckets
	} else {
		return resid
	}
}

func (n *NodeHash) Insert(a int, weight float64) {
	for i := 0; i < n.numRows; i++ {
		bucket := n.Hash(a, i)
		n.count[i][bucket] += weight
	}
}

func (n *NodeHash) GetCount(a int) float64 {
	bucket := n.Hash(a, 0)
	minCount := n.count[0][bucket]
	for i := 1; i < n.numRows; i++ {
		bucket = n.Hash(a, i)
		minCount = math.Min(minCount, n.count[i][bucket])
	}
	return minCount
}

func (n *NodeHash) Clear() {
	for i, row := range n.count {
		for j, _ := range row {
			n.count[i][j] = 0
		}
	}
}

func (n *NodeHash) Lower(factor float64) {
	for i, row := range n.count {
		for j, _ := range row {
			n.count[i][j] *= factor
		}
	}
}
