package test
import (
	"testing"
	"gotopo/geom"
	"github.com/stretchr/testify/assert"
)

func TestNewLineInsufficientCoords(t *testing.T) {
	assert.Panics(t, func() {geom.NewLine(geom.NewCoords64())})
}

func TestNewLineCoords(t *testing.T) {
	geom.NewLine(geom.NewCoords64FromSlice(2, []float64{1, 1, 5, 5}))
}