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


type MidasModel struct {
	curCount *EdgeHash
	totalCount *EdgeHash
	curT int
}

// Creates a new Midas struct that will enable the use of
// Fit and FitPredict API.
func NewMidasModel(numRows int, numBuckets int, m int) *MidasModel {
	return &MidasModel{
		totalCount: NewEdgeHash(numRows, numBuckets, m),
		curCount: NewEdgeHash(numRows, numBuckets, m),
		curT: 1,
	}
}

// Fit the source, destination and time to the MidasModel struct
// similar to the sklearn api
func (m *MidasModel) Fit(src, dst, time int){
	if time > m.curT {
		m.curCount.Clear()
		m.curT = time
	}
	m.curCount.Insert(src, dst, 1)
	m.totalCount.Insert(src, dst, 1)
}

// Fit the source, destination and time to the MidasModel struct and
// calculate the anomaly score
func (m *MidasModel) FitPredict(src, dst, time int) float64{
	m.Fit(src, dst, time)
	curMean := m.totalCount.GetCount(src, dst) / float64(m.curT)
	sqerr := math.Pow(m.curCount.GetCount(src, dst)-curMean, 2)
	var curScore float64
	if m.curT == 1 {
		curScore = 0
	} else {
		curScore = sqerr/curMean + sqerr/(curMean*(float64(m.curT)-1))
	}
	return curScore
}

// Takes in a list of source, destination and times to do anomaly score of each edge
// This function mirrors the implementation of https://github.com/bhatiasiddharth/MIDAS
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
