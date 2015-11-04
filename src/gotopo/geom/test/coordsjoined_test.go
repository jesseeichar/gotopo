package test
import (
	"testing"
	"gotopo/geom"
	"github.com/stretchr/testify/assert"
)

func TestCoordsJoined(t *testing.T) {
	assert.Panics(t, func() {
		geom.NewCoordsJoined(geom.NewCoords64WithDimensions(2), geom.NewCoords64WithDimensions(3))
	}, "All functions should have same number of dimensions")
	assert.Panics(t, func() {
		geom.NewCoordsJoined()
	}, "There should be at least on Coords object in the Joined Coords")
}
func coordsJoined2D(values []float64) geom.Coords {
	if len(values) < 2 {
		return geom.NewCoords64()
	}
	c1 := geom.NewCoords64FromSlice(2, values[:2])
	c2 := geom.NewCoords64FromSlice(2, values[2:])
	return geom.NewCoordsJoined(c1, c2)
}
func coordsJoined3D(values []float64) geom.Coords {
	if len(values) < 3 {
		return geom.NewCoords64WithDimensions(3)
	}
	c1 := geom.NewCoords64FromSlice(3, values[:3])
	c2 := geom.NewCoords64FromSlice(3, values[3:])
	return geom.NewCoordsJoined(c1, c2)
}
func TestCoordsJoinedEquals(t *testing.T) {
	CoordsEqualsTestImpl(t, coordsJoined2D, coordsJoined3D)
}
func TestCoordsJoinedGet(t *testing.T) {
	CoordsGetTestImpl(t, coordsJoined2D)
}
func TestIsEmptyJoinedGet(t *testing.T) {
	CoordsIsEmptyTestImpl(t, coordsJoined2D)
}
func TestNumCoordsJoinedGet(t *testing.T) {
	CoordsNumCoordsTestImpl(t, coordsJoined2D)
}
func TestNumDimJoinedGet(t *testing.T) {
	CoordsNumDimTestImpl(t, coordsJoined2D, coordsJoined3D)
}