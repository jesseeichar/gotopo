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
