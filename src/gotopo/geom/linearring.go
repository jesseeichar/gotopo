package geom

type LinearRing struct {
	line Line
}

var _ Geometry = LinearRing{}       // Verify that coordsJoined implements Geometry.
var _ Geometry = (*LinearRing)(nil) // Verify that *coordsJoined implements Geometry

func NewLinearRingFromLine(line Line) LinearRing {
	assertValidLinearRingValues(line.coordinates)
	return LinearRing{line}

}
func NewLinearRingFromCoords(coordinates Coords) LinearRing {
	assertValidLinearRingValues(coordinates)
	return LinearRing{NewLine(coordinates)}
}

func (this LinearRing) Coords() Coords {
	return this.line.Coords()
}

func assertValidLinearRingValues(coordinates Coords) {
	if coordinates.NumCoords() < 3 {
		panic("At least three coordinates are needed to make a line")
	}

	if !coordinates.Get(0).Equals(coordinates.Get(coordinates.NumCoords() - 1)) {
		panic("First and last coordinate must be the same")
	}
}

func (this LinearRing) Visit(visitor GeometryVisitor) {
	visitor(this)
}

func (this LinearRing) Bounds() Bounds {
	return this.line.Bounds()
}

func (this LinearRing) Equals(other Geometry) bool {
	switch ring := other.(type) {
	default:
		return false
	case LinearRing:
		return this.line.Equals(ring.line)
	}
}