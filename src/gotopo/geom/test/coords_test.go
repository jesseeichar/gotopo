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

func CoordsInsertRawTestImpl(t *testing.T, simpleCoords geom.ReadWriteCoords) {
	assert.True(t, simpleCoords.IsEmpty())

	simpleCoords.InsertRaw(0, []float64{2, 3})
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")

	assert.False(t, simpleCoords.IsEmpty())

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

func CoordsInsertTestImpl(t *testing.T, simpleCoords geom.ReadWriteCoords) {
	assert.True(t, simpleCoords.IsEmpty())

	assert.Panics(t, func() {simpleCoords.Insert(1, geom.NewPoint(66, 32))})
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")

	assert.True(t, simpleCoords.IsEmpty())

	simpleCoords.Insert(0, geom.NewPoint(1, 2))
	assert.Equal(t, int(simpleCoords.NumCoords()), 1, "Number of coordinates should be the 1")
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	assert.False(t, simpleCoords.IsEmpty())

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

func CoordsAddTestImpl(t *testing.T, simpleCoords geom.ReadWriteCoords) {
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

func CoordsSetTestImpl(t *testing.T, coordsFunc func([]float64) geom.ReadWriteCoords) {
	simpleCoords := coordsFunc([]float64{1, 2, 3, 4, 5, 6})

	assert.False(t, simpleCoords.IsEmpty())

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

func CoordsEqualsTestImpl(t *testing.T, coordsFunc2D, coordsFunc3D func([]float64) geom.Coords) {
	c1 := coordsFunc2D([]float64{1, 1, 2, 2, 3, 3})
	c2 := coordsFunc2D([]float64{1, 1, 2, 2, 3, 3})
	c3 := coordsFunc2D([]float64{1, 1, 2, 2})
	c4 := coordsFunc2D([]float64{1, 1, 2, 2, 4, 4})
	c5 := coordsFunc3D([]float64{1, 1, 1, 2, 2, 2, 4, 4, 4})
	c6 := coordsFunc3D([]float64{1, 1, 1, 2, 2, 2, 4, 4, 4})
	c7 := coordsFunc3D([]float64{1, 1, 1, 2, 2, 2, 4, 4, 5})

	assert.True(t, geom.EqualsCoords(c1, c1))
	assert.True(t, geom.EqualsCoords(c1, c2))
	assert.False(t, geom.EqualsCoords(c1, c3))
	assert.False(t, geom.EqualsCoords(c1, c4))
	assert.False(t, geom.EqualsCoords(c1, c5))
	assert.True(t, geom.EqualsCoords(c5, c5))
	assert.True(t, geom.EqualsCoords(c5, c6))
	assert.False(t, geom.EqualsCoords(c5, c7))
}

func CoordsGetTestImpl(t *testing.T, coordsFunc2D func([]float64) geom.Coords) {
	c1 := coordsFunc2D([]float64{1, 2, 3, 4, 5, 6})

	assert.Equal(t, float64(1), c1.Get(0).X())
	assert.Equal(t, float64(2), c1.Get(0).Y())

	assert.Equal(t, float64(3), c1.Get(1).X())
	assert.Equal(t, float64(4), c1.Get(1).Y())

	assert.Equal(t, float64(5), c1.Get(2).X())
	assert.Equal(t, float64(6), c1.Get(2).Y())

	assert.Panics(t, func() {c1.Get(3)})
	assert.Panics(t, func() {c1.Get(4)})
}
func CoordsIsEmptyTestImpl(t *testing.T, coordsFunc func([]float64) geom.Coords) {
	assert.True(t, coordsFunc([]float64{}).IsEmpty())
	assert.False(t, coordsFunc([]float64{1, 2}).IsEmpty())
}
func CoordsNumCoordsTestImpl(t *testing.T, coordsFunc2D func([]float64) geom.Coords) {
	assert.Equal(t, 3, int(coordsFunc2D([]float64{1, 2, 3, 4, 5, 6}).NumCoords()))
	assert.Equal(t, 2, int(coordsFunc2D([]float64{1, 2, 3, 4}).NumCoords()))
	assert.Equal(t, 1, int(coordsFunc2D([]float64{3, 4}).NumCoords()))
	assert.Equal(t, 0, int(coordsFunc2D([]float64{}).NumCoords()))


}
func CoordsNumDimTestImpl(t *testing.T, coordsFunc2D, coordsFunc3D func([]float64) geom.Coords) {
	assert.Equal(t, 2, int(coordsFunc2D([]float64{1, 2, 3, 4, 5, 6}).NumDim()))
	assert.Equal(t, 2, int(coordsFunc2D([]float64{1, 2, 3, 4}).NumDim()))
	assert.Equal(t, 3, int(coordsFunc3D([]float64{3, 4, 3}).NumDim()))
	assert.Equal(t, 3, int(coordsFunc3D([]float64{}).NumDim()))
}