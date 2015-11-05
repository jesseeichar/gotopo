package geom
import (
	"bytes"
	"fmt"
)
const PRECISION = 0.00001

type Point interface {
	Geometry
	X() float64
	Y() float64
	Ord(dimIdx uint8) float64
	NumDim() uint8
	ToArray() []float64
	String() string
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