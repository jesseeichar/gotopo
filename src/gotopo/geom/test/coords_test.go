package test
import (
	"testing"
	"gotopo/geom"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestCoords64EmptyConstructor(t *testing.T) {
	simpleCoords := geom.NewCoords64()
	assert.Equal(t, int(geom.DEFAULT_NUM_DIMENSIONS), int(simpleCoords.NumDim()), "Number of dimensions should be the default number")
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")
}

func TestCoords64CapacityConstructor(t *testing.T) {
	simpleCoords := geom.NewCoords64WithCapacity(27)

	assert.Equal(t, int(geom.DEFAULT_NUM_DIMENSIONS), int(simpleCoords.NumDim()), "Number of dimensions should be the default number")
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")
}

func TestCoords64DimensionsConstructor(t *testing.T) {
	simpleCoords := geom.NewCoords64WithDimensions(4)

	assert.Equal(t, 4, int(simpleCoords.NumDim()), "Number of dimensions should be the 4")
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")
}

func TestCoords64CapacityAndDimensionsConstructor(t *testing.T) {
	simpleCoords := geom.NewCoords64WithCapacityAndDimensions(100, 3)

	assert.Equal(t, 3, int(simpleCoords.NumDim()), "Number of dimensions should be the 3")
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")
}
func TestCoords64FromSlice(t *testing.T) {
	simpleCoords, err := geom.NewCoords64FromSlice(3, []float64{1, 2, 3, 4, 5, 6})

	assert.NoError(t, err)
	assert.Equal(t, 3, int(simpleCoords.NumDim()), "Number of dimensions should be the 3")
	assert.Equal(t, 2, int(simpleCoords.NumCoords()), "Number of coordinates should be the 2")
	assert.Equal(t, []float64{1, 2, 3}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{4, 5, 6}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	simpleCoords, err = geom.NewCoords64FromSlice(3, []float64{1, 2, 3, 4, 5})
	assert.Error(t, err)
	assert.Nil(t, simpleCoords)

	simpleCoords, err = geom.NewCoords64FromSlice(3, []float64{})
	assert.NoError(t, err)
	assert.Equal(t, 3, int(simpleCoords.NumDim()), "Number of dimensions should be the 3")
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")
}

func TestCoords64InsertRaw(t *testing.T) {
	simpleCoords := geom.NewCoords64()

	err := simpleCoords.InsertRaw(0, []float64{2, 3})
	assert.NoError(t, err, "Did not expect an error")
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")


	err = simpleCoords.InsertRaw(0, []float64{2, 3, 4})
	assert.Error(t, err, "Did not expect an error")
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")

	err = simpleCoords.InsertRaw(0, []float64{2})
	assert.Error(t, err)
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")

	err = simpleCoords.InsertRaw(0, []float64{2})
	assert.Error(t, err)
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")

	err = simpleCoords.InsertRaw(2, []float64{2, 3})
	assert.Error(t, err)
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")

	err = simpleCoords.InsertRaw(1, []float64{4, 5})
	assert.NoError(t, err)
	assert.Equal(t, 2, int(simpleCoords.NumCoords()), "Number of coordinates should be the 2")
	assert.Equal(t, []float64{2, 3}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{4, 5}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect second coordinate in %v", simpleCoords))

	err = simpleCoords.InsertRaw(1, []float64{3, 4})
	assert.NoError(t, err)
	assert.Equal(t, 3, int(simpleCoords.NumCoords()), "Number of coordinates should be the 3")
	assert.Equal(t, []float64{2, 3}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{3, 4}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect second coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{4, 5}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect second coordinate in %v", simpleCoords))

	err = simpleCoords.InsertRaw(0, []float64{-1, 0})
	assert.NoError(t, err)
	assert.Equal(t, 4, int(simpleCoords.NumCoords()), "Number of coordinates should be the 4")
	assert.Equal(t, []float64{-1, 0}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{2, 3}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{3, 4}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect second coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{4, 5}, simpleCoords.Get(3).ToArray(), fmt.Sprintf("Incorrect second coordinate in %v", simpleCoords))
}

func TestCoords64Insert(t *testing.T) {
	simpleCoords := geom.NewCoords64()

	err := simpleCoords.Insert(1, geom.NewCoord(66, 32))
	assert.Error(t, err)
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")

	err = simpleCoords.Insert(0, geom.NewCoord(1, 2))
	assert.NoError(t, err)
	assert.Equal(t, int(simpleCoords.NumCoords()), 1, "Number of coordinates should be the 1")
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	err = simpleCoords.Insert(2, geom.NewCoord(66, 32))
	assert.Error(t, err)
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	err = simpleCoords.Insert(0, geom.NewCoord(-1, 0))
	assert.NoError(t, err)
	assert.Equal(t, 2, int(simpleCoords.NumCoords()), "Number of coordinates should be the 2")
	assert.Equal(t, []float64{-1, 0}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	err = simpleCoords.Insert(1, geom.NewCoord(1, 1))
	assert.NoError(t, err)
	assert.Equal(t, 3, int(simpleCoords.NumCoords()), "Number of coordinates should be the 3")
	assert.Equal(t, []float64{-1, 0}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{1, 1}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	err = simpleCoords.Insert(1, geom.NewCoord(1, 1, 4))
	assert.Error(t, err)
	assert.Equal(t, 3, int(simpleCoords.NumCoords()), "Number of coordinates should be the 3")
	assert.Equal(t, []float64{-1, 0}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{1, 1}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
}

func TestCoords64Add(t *testing.T) {
	simpleCoords := geom.NewCoords64()

	err := simpleCoords.Add(geom.NewCoord(1, 2))
	assert.NoError(t, err)
	assert.Equal(t, 1, int(simpleCoords.NumCoords()), "Number of coordinates should be the 1")
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	err = simpleCoords.Add(geom.NewCoord(3, 4))
	assert.NoError(t, err)
	assert.Equal(t, 2, int(simpleCoords.NumCoords()), "Number of coordinates should be the 2")
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{3, 4}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	err = simpleCoords.Add(geom.NewCoord(3, 4, 5))
	assert.Error(t, err)
	assert.Equal(t, 2, int(simpleCoords.NumCoords()), "Number of coordinates should be the 2")
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{3, 4}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
}

func TestCoords64Set(t *testing.T) {
	simpleCoords, _ := geom.NewCoords64FromSlice(2, []float64{1, 2, 3, 4, 5, 6})

	err := simpleCoords.Set(1, geom.NewCoord(30, 40))
	assert.NoError(t, err)
	assert.Equal(t, 3, int(simpleCoords.NumCoords()), "Number of coordinates should be the 3")
	assert.Equal(t, []float64{1, 2}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{30, 40}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{5, 6}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	err = simpleCoords.Set(0, geom.NewCoord(10, 20))
	assert.NoError(t, err)
	assert.Equal(t, int(simpleCoords.NumCoords()), 3, "Number of coordinates should be the 3")
	assert.Equal(t, []float64{10, 20}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{30, 40}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{5, 6}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	err = simpleCoords.Set(2, geom.NewCoord(50, 60))
	assert.NoError(t, err)
	assert.Equal(t, int(simpleCoords.NumCoords()), 3, "Number of coordinates should be the 3")
	assert.Equal(t, []float64{10, 20}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{30, 40}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{50, 60}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	err = simpleCoords.Set(3, geom.NewCoord(99, 9))
	assert.Error(t, err)
	assert.Equal(t, int(simpleCoords.NumCoords()), 3, "Number of coordinates should be the 3")
	assert.Equal(t, []float64{10, 20}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{30, 40}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{50, 60}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	err = simpleCoords.Set(2, geom.NewCoord(99, 99, 99))
	assert.Error(t, err)
	assert.Equal(t, int(simpleCoords.NumCoords()), 3, "Number of coordinates should be the 3")
	assert.Equal(t, []float64{10, 20}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{30, 40}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{50, 60}, simpleCoords.Get(2).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
}