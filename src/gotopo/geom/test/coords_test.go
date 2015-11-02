package test
import (
	"testing"
	"gotopo/geom"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestCoordString(t *testing.T) {
	coord := geom.NewPoint(1, 2, 3, 4)
	assert.Equal(t, "(1.000, 2.000, 3.000, 4.000)", coord.String())
}

func CoordsInsertRawTestImpl(t *testing.T, simpleCoords geom.Coords) {
	simpleCoords.InsertRaw(0, []float64{2, 3})
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")


	assert.Panics(t, func() {simpleCoords.InsertRaw(0, []float64{2, 3, 4})})
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")

	assert.Panics(t, func() {simpleCoords.InsertRaw(0, []float64{2})})
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")

	assert.Panics(t, func() {simpleCoords.InsertRaw(0, []float64{2})})
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")

	assert.Panics(t, func() {simpleCoords.InsertRaw(2, []float64{2, 3})})
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")

	simpleCoords.InsertRaw(1, []float64{4, 5})
	assert.Equal(t, 2, int(simpleCoords.NumCoords()), "Number of coordinates should be the 2")
	assert.Equal(t, []float64{2, 3}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{4, 5}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect second coordinate in %v", simpleCoords))

	simpleCoords.InsertRaw(1, []float64{3, 4})
	assert.Equal(t, 3, int(simpleCoords.NumCoords()), "Number of coordinates should be the 3")
	assert.Equal(t, []float64{2, 3}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{3, 4}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect second coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{4, 5}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect second coordinate in %v", simpleCoords))

	simpleCoords.InsertRaw(0, []float64{-1, 0})
	assert.Equal(t, 4, int(simpleCoords.NumCoords()), "Number of coordinates should be the 4")
	assert.Equal(t, []float64{-1, 0}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{2, 3}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{3, 4}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect second coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{4, 5}, simpleCoords.Get(3).ToArray(), fmt.Sprintf("Incorrect second coordinate in %v", simpleCoords))
}

func CoordsInsertTestImpl(t *testing.T, simpleCoords geom.Coords) {
	assert.Panics(t, func() {simpleCoords.Insert(1, geom.NewPoint(66, 32))})
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")

	simpleCoords.Insert(0, geom.NewPoint(1, 2))
	assert.Equal(t, int(simpleCoords.NumCoords()), 1, "Number of coordinates should be the 1")
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	assert.Panics(t, func() {simpleCoords.Insert(2, geom.NewPoint(66, 32))})
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	simpleCoords.Insert(0, geom.NewPoint(-1, 0))
	assert.Equal(t, 2, int(simpleCoords.NumCoords()), "Number of coordinates should be the 2")
	assert.Equal(t, []float64{-1, 0}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	simpleCoords.Insert(1, geom.NewPoint(1, 1))
	assert.Equal(t, 3, int(simpleCoords.NumCoords()), "Number of coordinates should be the 3")
	assert.Equal(t, []float64{-1, 0}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{1, 1}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	assert.Panics(t, func() {simpleCoords.Insert(1, geom.NewPoint(1, 1, 4))})
	assert.Equal(t, 3, int(simpleCoords.NumCoords()), "Number of coordinates should be the 3")
	assert.Equal(t, []float64{-1, 0}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{1, 1}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
}

func CoordsAddTestImpl(t *testing.T, simpleCoords geom.Coords) {
	simpleCoords.Add(geom.NewPoint(1, 2))
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	simpleCoords.Add(geom.NewPoint(3, 4))
	assert.Equal(t, 2, int(simpleCoords.NumCoords()), "Number of coordinates should be the 2")
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{3, 4}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	assert.Panics(t, func() {simpleCoords.Add(geom.NewPoint(3, 4, 5))})
	assert.Equal(t, 2, int(simpleCoords.NumCoords()), "Number of coordinates should be the 2")
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{3, 4}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
}

func CoordsSetTestImpl(t *testing.T, coordsFunc func([]float64) geom.Coords) {
	simpleCoords := coordsFunc([]float64{1, 2, 3, 4, 5, 6})

	simpleCoords.Set(1, geom.NewPoint(30, 40))
	assert.Equal(t, 3, int(simpleCoords.NumCoords()), "Number of coordinates should be the 3")
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{30, 40}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{5, 6}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	simpleCoords.Set(0, geom.NewPoint(10, 20))
	assert.Equal(t, int(simpleCoords.NumCoords()), 3, "Number of coordinates should be the 3")
	assert.Equal(t, []float64{10, 20}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{30, 40}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{5, 6}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	simpleCoords.Set(2, geom.NewPoint(50, 60))
	assert.Equal(t, int(simpleCoords.NumCoords()), 3, "Number of coordinates should be the 3")
	assert.Equal(t, []float64{10, 20}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{30, 40}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{50, 60}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	assert.Panics(t, func() {simpleCoords.Set(3, geom.NewPoint(99, 9))})
	assert.Equal(t, int(simpleCoords.NumCoords()), 3, "Number of coordinates should be the 3")
	assert.Equal(t, []float64{10, 20}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{30, 40}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{50, 60}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	assert.Panics(t, func() {simpleCoords.Set(2, geom.NewPoint(99, 99, 99))})
	assert.Equal(t, int(simpleCoords.NumCoords()), 3, "Number of coordinates should be the 3")
	assert.Equal(t, []float64{10, 20}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{30, 40}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{50, 60}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
}