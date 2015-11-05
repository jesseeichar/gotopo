package test
import (
	"testing"
	"gotopo/geom"
	"github.com/stretchr/testify/assert"
	"gotopo/geom/coords64"
)

func TestNewLineInsufficientCoords(t *testing.T) {
	assert.Panics(t, func() {geom.NewLine(coords64.NewCoords64())})
}

func TestNewLineCoords(t *testing.T) {
	geom.NewLine(coords64.NewCoords64FromSlice(2, []float64{1, 1, 5, 5}))
}

func TestLineCoords(t *testing.T) {
	line := geom.NewLine(coords64.NewCoords64FromSlice(2, []float64{1, 1, 5, 5}))
	assert.Equal(t, 2, int(line.Coords().NumCoords()))
}