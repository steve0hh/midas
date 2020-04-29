package midas

import (
	"math"
)

func countsToAnom(tot float64, cur float64, curT int) float64 {
	curMean := tot / cur
	sqerr := math.Pow(max(0, cur-curMean), 2)
	return (sqerr/curMean + sqerr/(curMean*max(1.0, float64(curT-1.0))))
}

func MidasR(src []int, dst []int, times []int, numRows int, numBuckets int, factor float64) []float64 {
	m := biggest(src)
	curCount := NewEdgeHash(numRows, numBuckets, m)
	totalCount := NewEdgeHash(numRows, numBuckets, m)

	srcScore := NewNodeHash(numRows, numBuckets)
	dstScore := NewNodeHash(numRows, numBuckets)
	srcTotal := NewNodeHash(numRows, numBuckets)
	dstTotal := NewNodeHash(numRows, numBuckets)

	anomScore := make([]float64, len(src))
	var curT, curSrc, curDst int
	curT = 1
	var curScore, curScoreSrc, curScoreDst, combinedScore float64

	for i, _ := range src {
		if i == 0 || times[i] > curT {
			curCount.Lower(factor)
			srcScore.Lower(factor)
			dstScore.Lower(factor)
			curT = times[i]
		}

		curSrc = src[i]
		curDst = dst[i]
		curCount.Insert(curSrc, curDst, 1)
		totalCount.Insert(curSrc, curDst, 1)
		srcScore.Insert(curSrc, 1)
		dstScore.Insert(curDst, 1)
		srcTotal.Insert(curSrc, 1)
		dstTotal.Insert(curDst, 1)
		curScore = countsToAnom(totalCount.GetCount(curSrc, curDst), curCount.GetCount(curSrc, curDst), curT)
		curScoreSrc = countsToAnom(srcTotal.GetCount(curSrc), srcScore.GetCount(curSrc), curT)
		curScoreDst = countsToAnom(dstTotal.GetCount(curDst), dstScore.GetCount(curDst), curT)
		combinedScore = max(max(curScoreSrc, curScoreDst), curScore)
		anomScore[i] = math.Log(1 + combinedScore)
	}
	return anomScore
}
