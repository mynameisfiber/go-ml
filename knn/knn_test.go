package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestKNearestNeighboors(t *testing.T) {
	space := NewSpace(1)

	p1 := Point{Location: []float64{1.}, Category: 4}
	space.AddPoint(p1)

	p2 := Point{Location: []float64{5.}, Category: 3}
	space.AddPoint(p2)

	knn := space.KNearestNeighboors([]float64{2.0}, 2)
	assert.Equal(t, len(knn), 2)
	assert.Equal(t, knn[0], &p1)
	assert.Equal(t, knn[1], &p2)
}

func TestClassify(t *testing.T) {
	space := NewSpace(1)

	p1 := Point{Location: []float64{0.}, Category: 4}
	space.AddPoint(p1)

	p2 := Point{Location: []float64{5.}, Category: 3}
	space.AddPoint(p2)

	category := space.Classify([]float64{2.5}, 2)
	assert.Equal(t, category, 3.5)
}

func TestCategoryPDF(t *testing.T) {
	space := NewSpace(1)

	p1 := Point{Location: []float64{0.}, Category: 4}
	space.AddPoint(p1)

	p2 := Point{Location: []float64{1.}, Category: 3}
	space.AddPoint(p2)

	p3 := Point{Location: []float64{2.}, Category: 3}
	space.AddPoint(p3)

	p4 := Point{Location: []float64{3.}, Category: 1}
	space.AddPoint(p4)

	categoryPDF := space.CategoryPDF([]float64{1.5}, 4)

	actual := map[int]float64{
		4: 0.25,
		3: 0.5,
		1: 0.25,
	}
	assert.Equal(t, len(categoryPDF), len(actual))
	for cat, prob := range categoryPDF {
		assert.Equal(t, prob, actual[cat])
	}
}

func TestMaxPosterioriCategory(t *testing.T) {
	space := NewSpace(1)

	p1 := Point{Location: []float64{0.}, Category: 4}
	space.AddPoint(p1)

	p2 := Point{Location: []float64{1.}, Category: 3}
	space.AddPoint(p2)

	p3 := Point{Location: []float64{2.}, Category: 3}
	space.AddPoint(p3)

	p4 := Point{Location: []float64{3.}, Category: 1}
	space.AddPoint(p4)

	category, prob := space.MaxPosterioriCategory([]float64{1.5}, 4)

	assert.Equal(t, category, 3)
	assert.Equal(t, prob, 0.5)
}
