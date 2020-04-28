package midas

import (
	"math"
)

func max(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func biggest(values []int) int {
	b := values[0]
	for _, v := range values {
		if b < v {
			b = v
		}
	}
	return b
}

func countsToAnom(tot float64, cur float64, curT int) float64 {
	curMean := tot / cur
	sqerr := math.Pow(max(0, cur-curMean), 2)
	return (sqerr/curMean + sqerr/(curMean*max(1.0, float64(curT-1.0))))
}

func Midas(src []int, dst []int, times []int, numRows int, numBuckets int) []float64 {
	m := biggest(src)
	curCount := NewEdgeHash(numRows, numBuckets, m)
	totalCount := NewEdgeHash(numRows, numBuckets, m)
	anomScore := make([]float64, len(src))
	curT := 1
	for i, _ := range src {
		if i == 0 || times[i] > curT {
			curCount.Clear()
			curT = times[i]
		}

		curSrc := src[i]
		curDst := dst[i]
		curCount.Insert(curSrc, curDst, 1)
		totalCount.Insert(curSrc, curDst, 1)
		curMean := totalCount.GetCount(curSrc, curDst) / float64(curT)
		sqerr := math.Pow(curCount.GetCount(curSrc, curDst)-curMean, 2)

		var curScore float64
		if curT == 1 {
			curScore = 0
		} else {
			curScore = sqerr/curMean + sqerr/(curMean*(float64(curT)-1))
		}
		anomScore[i] = curScore
	}
	return anomScore
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
