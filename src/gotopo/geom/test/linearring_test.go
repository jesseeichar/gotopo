package test
import (
	"testing"
	"gotopo/geom"
	"github.com/stretchr/testify/assert"
)

func TestNewLinearRingInsufficientCoords(t *testing.T) {
	c := geom.NewCoords64FromSlice(2, []float64{1, 1, 2, 2})
	assert.Panics(t, func() {geom.NewLinearRingFromCoords(c)})
}

func TestNewLinearRingEndPointMismatchCoords(t *testing.T) {
	c := geom.NewCoords64FromSlice(2, []float64{1, 1, 2, 2, 3, 3})
	assert.Panics(t, func() {geom.NewLinearRingFromCoords(c)})
}

func TestNewLinearRingFromCoords(t *testing.T) {
	geom.NewLinearRingFromCoords(geom.NewCoords64FromSlice(2, []float64{1, 1, 5, 5, 1, 1}))
}

func TestNewLinearRingInsufficientLine(t *testing.T) {
	c := geom.NewCoords64FromSlice(2, []float64{1, 1, 2, 2})
	assert.Panics(t, func() {geom.NewLinearRingFromCoords(c)})
}

func TestNewLinearRingEndPointMismatchLine(t *testing.T) {
	c := geom.NewCoords64FromSlice(2, []float64{1, 1, 2, 2, 3, 3})
	assert.Panics(t, func() {geom.NewLinearRingFromCoords(c)})
}

func TestNewLinearRingFromLine(t *testing.T) {
	c := geom.NewCoords64FromSlice(2, []float64{1, 1, 5, 5, 1, 1})
	geom.NewLinearRingFromCoords(c)
}
