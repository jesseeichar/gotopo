package coordsjoined_test
import (
	"testing"
	"gotopo/geom"
	"github.com/stretchr/testify/assert"
	"gotopo/geom/coords64"
	"gotopo/geom/test"
)

func TestCoordsJoined(t *testing.T) {
	assert.Panics(t, func() {
		geom.NewCoordsJoined(coords64.NewCoords64WithDimensions(2), coords64.NewCoords64WithDimensions(3))
	}, "All functions should have same number of dimensions")
	assert.Panics(t, func() {
		geom.NewCoordsJoined()
	}, "There should be at least on Coords object in the Joined Coords")
}
func coordsJoined2D(values []float64) geom.Coords {
	if len(values) < 2 {
		return coords64.NewCoords64()
	}
	c1 := coords64.NewCoords64FromSlice(2, values[:2])
	c2 := coords64.NewCoords64FromSlice(2, values[2:])
	return geom.NewCoordsJoined(c1, c2)
}
func coordsJoined3D(values []float64) geom.Coords {
	if len(values) < 3 {
		return coords64.NewCoords64WithDimensions(3)
	}
	c1 := coords64.NewCoords64FromSlice(3, values[:3])
	c2 := coords64.NewCoords64FromSlice(3, values[3:])
	return geom.NewCoordsJoined(c1, c2)
}
func TestCoordsJoinedEquals(t *testing.T) {
	test.CoordsEqualsTestImpl(t, coordsJoined2D, coordsJoined3D)
}
func TestCoordsJoinedGet(t *testing.T) {
	test.CoordsGetTestImpl(t, coordsJoined2D)
}
func TestIsEmptyJoinedGet(t *testing.T) {
	test.CoordsIsEmptyTestImpl(t, coordsJoined2D)
}
func TestNumCoordsJoinedGet(t *testing.T) {
	test.CoordsNumCoordsTestImpl(t, coordsJoined2D)
}
func TestNumDimJoinedGet(t *testing.T) {
	test.CoordsNumDimTestImpl(t, coordsJoined2D, coordsJoined3D)
}