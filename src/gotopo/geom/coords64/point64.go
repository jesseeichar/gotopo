package coords64
import "gotopo/geom"
type point64 struct {
	data  coords64
	index uint32
}

var _ geom.Point = point64{}       // Verify that point64 implements Point.
var _ geom.Point = (*point64)(nil) // Verify that *point64 implements Point

func NewPoint(ordinals ...float64) geom.Point {
	coords := NewCoords64FromSlice(uint8(len(ordinals)), ordinals)
	return coords.Get(0)
}
func (this point64) X() float64 {
	return this.data.data[this.index]
}
func (this point64) Y() float64 {
	return this.data.data[this.index + 1]

}
func (this point64) Ord(dimIdx uint8) float64 {
	return this.data.data[this.index + uint32(dimIdx)]
}
func (this point64) NumDim() uint8 {
	return this.data.dimensions
}
func (this point64) ToArray() []float64 {
	endIdx := this.index + uint32(this.data.dimensions)
	return this.data.data[this.index:endIdx]
}
func (this point64) String() string {
	return geom.CoordString(this)
}

func (this point64) Coords() geom.Coords {
	return this.data
}

func (this point64) Equals(other geom.Geometry) bool {
	switch point := other.(type) {
	default:
		return false
	case geom.Point:
		if this.NumDim() != point.NumDim() {
			return false
		}
		for i := uint8(0); i < this.NumDim(); i++ {
			if this.Ord(i) != point.Ord(i) {
				return false
			}
		}
		return true;
	}
}

func (this point64) Visit(visitor geom.GeometryVisitor) {
	visitor(this)
}


