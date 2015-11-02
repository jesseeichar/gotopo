package geom
import (
	"bytes"
	"fmt"
)

type Point interface {
	X() float64
	Y() float64
	Ord(dimIdx uint8) float64
	NumDim() uint8
	ToArray() []float64
	String() string
}

func NewPoint(ordinals ...float64) Point {
	coords := NewCoords64WithCapacityAndDimensions(1, uint8(len(ordinals)))
	coords.InsertRaw(0, ordinals)
	return coords.Get(0);
}

func CoordString(coord Point) string {
	buffer := bytes.Buffer{};
	for i := uint8(0); i < coord.NumDim(); i++ {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(fmt.Sprintf("%.3f", coord.Ord(i)))
	}
	return "(" + buffer.String() + ")"
}