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
	simpleCoords := geom.NewCoords64FromSlice(3, []float64{1, 2, 3, 4, 5, 6})

	assert.Equal(t, 3, int(simpleCoords.NumDim()), "Number of dimensions should be the 3")
	assert.Equal(t, 2, int(simpleCoords.NumCoords()), "Number of coordinates should be the 2")
	assert.Equal(t, []float64{1, 2, 3}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{4, 5, 6}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	assert.Panics(t, func() {geom.NewCoords64FromSlice(3, []float64{1, 2, 3, 4, 5})})

	simpleCoords = geom.NewCoords64FromSlice(3, []float64{})
	assert.Equal(t, 3, int(simpleCoords.NumDim()), "Number of dimensions should be the 3")
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")
}

func TestCoords64InsertRaw(t *testing.T) {
	CoordsInsertRawTestImpl(t, geom.NewCoords64())
}

func TestCoords64Insert(t *testing.T) {
	CoordsInsertTestImpl(t, geom.NewCoords64())
}

func TestCoords64Add(t *testing.T) {
	CoordsAddTestImpl(t, geom.NewCoords64())
}

func TestCoords64Set(t *testing.T) {
	CoordsSetTestImpl(t, func(values []float64) geom.Coords {
		return geom.NewCoords64FromSlice(2, values)
	})
}