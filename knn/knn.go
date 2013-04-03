package main

import (
	"container/list"
	"fmt"
	"math"
)

var (
	IncorrectDimensions = fmt.Errorf("Specified point has the incorrect dimentionality")
)

type Point struct {
	Location []float64
	Category int
}

func (p Point) Distance(other Point) float64 {
	if len(p.Location) != len(other.Location) {
		return -1
	}
	dist := 0.0
	for i := 0; i < len(p.Location); i++ {
		dist += (p.Location[i] - other.Location[i]) * (p.Location[i] - other.Location[i])
	}
	return math.Sqrt(dist)
}

type Space struct {
	Points *list.List
	dim    int
}

func NewSpace(dim int) Space {
	return Space{Points: list.New(), dim: dim}
}

func (s *Space) AddPoint(p Point) error {
	if len(p.Location) != s.dim {
		return IncorrectDimensions
	}
	s.Points.PushBack(p)
	return nil
}

type tmpResult struct {
	point *Point
	dist  float64
}

func (s *Space) KNearestNeighboors(location []float64, K int) []*Point {
	if s.Points.Len() < K {
		return nil
	}

	this := &Point{location, -1}
	buf := make([]tmpResult, 0, K)
	for test := s.Points.Front(); test != nil; test = test.Next() {
		testPoint := test.Value.(Point)
		if len(buf) < K {
			buf = append(buf, tmpResult{&testPoint, this.Distance(testPoint)})
			continue
		}
		dist := this.Distance(testPoint)
		for i, k := range buf {
			if dist < k.dist {
				buf[i] = tmpResult{&testPoint, this.Distance(testPoint)}
				break
			}
		}
	}

	points := make([]*Point, K)
	for i := 0; i < K; i++ {
		points[i] = buf[i].point
	}
	return points
}

func (s *Space) Classify(location []float64, K int) float64 {
	points := s.KNearestNeighboors(location, K)

	category := 0.0
	for _, p := range points {
		category += float64(p.Category)
	}
	category /= float64(K)

	return category
}

func (s *Space) CategoryPDF(location []float64, K int) map[int]float64 {
	points := s.KNearestNeighboors(location, K)
	categories := make(map[int]float64)

	for _, p := range points {
		categories[p.Category] += 1
	}

	Kf := float64(K)
	for k, _ := range categories {
		categories[k] /= Kf
	}

	return categories
}

func (s *Space) MaxPosterioriCategory(location []float64, K int) (int, float64) {
	categoryPDF := s.CategoryPDF(location, K)

	mapCategory := 0
	mapProb := 0.0
	for category, prob := range categoryPDF {
		if prob > mapProb {
			mapCategory, mapProb = category, prob
		}
	}
	return mapCategory, mapProb
}
