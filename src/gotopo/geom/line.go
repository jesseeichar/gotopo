package geom

type Line struct {
	coordinates Coords
}

var _ Geometry = Line{}       // Verify that coordsJoined implements Geometry.
var _ Geometry = (*Line)(nil) // Verify that *coordsJoined implements Geometry

func NewLine(coordinates Coords) Line {
	if coordinates.NumCoords() < 2 {
		panic("At least two coordinates are needed to make a line")
	}
	return Line{coordinates}
}

func (this Line) Coords() Coords {
	return this.coordinates;
}

func (this Line) Visit(visitor GeometryVisitor) {
	visitor(this)
}

func (this Line) Equals(other Geometry) bool {
	switch other.(type) {
	default:
		return false
	case Line:
		return EqualsCoords(this.coordinates, other.Coords())
	}
}

func (this Line) Bounds() Bounds {
	min := this.coordinates.Get(0).ToArray()
	max := append(make([]float64, len(min)), min...)

	for i := uint32(1); i < this.coordinates.NumCoords() - 1; i++ {
		coord := this.coordinates.Get(i)
		for dim := uint8(0); dim < coord.NumDim(); dim++ {
			ord := coord.Ord(dim)
			if ord < min[dim] {
				min[dim] = ord
			}
			if ord > max[dim] {
				max[dim] = ord
			}
		}
	}
	boundsCoords := this.coordinates.Factory().CreateFromRawData(min...)
	boundsCoords.InsertRaw(1, max)
	return Bounds{min:boundsCoords.Get(0), max:boundsCoords.Get(1)}

}
