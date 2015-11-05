package coords64
import (
	"testing"
	"gotopo/geom"
	"github.com/stretchr/testify/assert"
	"fmt"
	"gotopo/geom/test"
)

func TestCoords64EmptyConstructor(t *testing.T) {
	simpleCoords := NewCoords64()
	assert.Equal(t, int(geom.DEFAULT_NUM_DIMENSIONS), int(simpleCoords.NumDim()), "Number of dimensions should be the default number")
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")
}

func TestCoords64CapacityConstructor(t *testing.T) {
	simpleCoords := NewCoords64WithCapacity(27)

	assert.Equal(t, int(geom.DEFAULT_NUM_DIMENSIONS), int(simpleCoords.NumDim()), "Number of dimensions should be the default number")
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")
}

func TestCoords64DimensionsConstructor(t *testing.T) {
	simpleCoords := NewCoords64WithDimensions(4)

	assert.Equal(t, 4, int(simpleCoords.NumDim()), "Number of dimensions should be the 4")
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")
}

func TestCoords64CapacityAndDimensionsConstructor(t *testing.T) {
	simpleCoords := NewCoords64WithCapacityAndDimensions(100, 3)

	assert.Equal(t, 3, int(simpleCoords.NumDim()), "Number of dimensions should be the 3")
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")
}
func TestCoords64FromSlice(t *testing.T) {
	simpleCoords := NewCoords64FromSlice(3, []float64{1, 2, 3, 4, 5, 6})

	assert.Equal(t, 3, int(simpleCoords.NumDim()), "Number of dimensions should be the 3")
	assert.Equal(t, 2, int(simpleCoords.NumCoords()), "Number of coordinates should be the 2")
	assert.Equal(t, []float64{1, 2, 3}, simpleCoords.Get(0).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))
	assert.Equal(t, []float64{4, 5, 6}, simpleCoords.Get(1).ToArray(), fmt.Sprintf("Incorrect first coordinate in %v", simpleCoords))

	assert.Panics(t, func() {NewCoords64FromSlice(3, []float64{1, 2, 3, 4, 5})})

	simpleCoords = NewCoords64FromSlice(3, []float64{})
	assert.Equal(t, 3, int(simpleCoords.NumDim()), "Number of dimensions should be the 3")
	assert.Equal(t, 0, int(simpleCoords.NumCoords()), "Number of coordinates should be the 0")
}

func TestCoords64InsertRaw(t *testing.T) {
	test.CoordsInsertRawTestImpl(t, NewCoords64())
}

func TestCoords64Insert(t *testing.T) {
	test.CoordsInsertTestImpl(t, NewCoords64(), NewPoint)
}

func TestCoords64Add(t *testing.T) {
	test.CoordsAddTestImpl(t, NewCoords64(), NewPoint)
}
func rwCoords642D(values []float64) geom.ReadWriteCoords {
	return NewCoords64FromSlice(2, values)
}
func rwCoords643D(values []float64) geom.ReadWriteCoords {
	return NewCoords64FromSlice(3, values)
}
func coords642D(values []float64) geom.Coords {
	return geom.Coords(rwCoords642D(values))
}
func coords643D(values []float64) geom.Coords {
	return geom.Coords(rwCoords643D(values))
}
func TestCoords64Set(t *testing.T) {
	test.CoordsSetTestImpl(t, NewPoint, rwCoords642D)
}

func TestCoords64Equals(t *testing.T) {
	test.CoordsEqualsTestImpl(t, coords642D, coords643D)
}
func TestCoords64Get(t *testing.T) {
	test.CoordsGetTestImpl(t, coords642D)
}
func TestIsEmpty64Get(t *testing.T) {
	test.CoordsIsEmptyTestImpl(t, coords642D)
}
func TestNumCoords64Get(t *testing.T) {
	test.CoordsNumCoordsTestImpl(t, coords642D)
}
func TestNumDim64Get(t *testing.T) {
	test.CoordsNumDimTestImpl(t, coords642D, coords643D)
}