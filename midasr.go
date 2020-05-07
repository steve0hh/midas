package midas

import (
	"math"
)

func countsToAnom(tot float64, cur float64, curT int) float64 {
	curMean := tot / cur
	sqerr := math.Pow(max(0, cur-curMean), 2)
	return (sqerr/curMean + sqerr/(curMean*max(1.0, float64(curT-1.0))))
}

type MidasRModel struct {
	curCount   *EdgeHash
	totalCount *EdgeHash
	srcScore   *NodeHash
	dstScore   *NodeHash
	srcTotal   *NodeHash
	dstTotal   *NodeHash
	curT       int
	factor     float64
}

func NewMidasRModel(numRows int, numBuckets int, m int, factor float64) *MidasRModel {
	return &MidasRModel{
		curCount:   NewEdgeHash(numRows, numBuckets, m),
		totalCount: NewEdgeHash(numRows, numBuckets, m),
		srcScore:   NewNodeHash(numRows, numBuckets),
		dstScore:   NewNodeHash(numRows, numBuckets),
		srcTotal:   NewNodeHash(numRows, numBuckets),
		dstTotal:   NewNodeHash(numRows, numBuckets),
		curT:       1,
		factor:     factor,
	}
}

func (m *MidasRModel) Fit(src, dst, time int) {
	if time > m.curT {
		m.curCount.Lower(m.factor)
		m.srcScore.Lower(m.factor)
		m.dstScore.Lower(m.factor)
		m.curT = time
	}
	m.curCount.Insert(src, dst, 1)
	m.totalCount.Insert(src, dst, 1)
	m.srcScore.Insert(src, 1)
	m.dstScore.Insert(src, 1)
	m.srcTotal.Insert(src, 1)
	m.dstTotal.Insert(dst, 1)
}

func (m *MidasRModel) FitPredict(src, dst, time int) float64 {
	m.Fit(src, dst, time)
	score := countsToAnom(m.totalCount.GetCount(src, dst), m.curCount.GetCount(src, dst), time)
	scoreSrc := countsToAnom(m.srcTotal.GetCount(src), m.srcScore.GetCount(src), time)
	scoreDst := countsToAnom(m.dstTotal.GetCount(dst), m.dstScore.GetCount(dst), time)
	combinedScore := max(max(scoreSrc, scoreDst), score)
	return math.Log(1 + combinedScore)
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
