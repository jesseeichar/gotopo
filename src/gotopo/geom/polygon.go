package geom

type Polygon struct {
	shell LinearRing
	holes []LinearRing
}

var _ Geometry = Polygon{}       // Verify that coordsJoined implements Geometry.
var _ Geometry = (*Polygon)(nil) // Verify that *coordsJoined implements Geometry

func NewPolygon(shell LinearRing, holes []LinearRing) Polygon {
	return Polygon{shell, holes}
}

func (this Polygon) Coords() Coords {
	coordsJoined := NewCoordsJoined(this.shell.Coords())
	for _, hole := range this.holes {
		coordsJoined.allCoords = append(coordsJoined.allCoords, hole.Coords())
	}
	return coordsJoined;
}
func (this Polygon) Shell() LinearRing {
	return this.shell
}

func (this Polygon) Hole(holeIdx uint32) LinearRing {
	return this.holes[holeIdx]
}

func (this Polygon) Visit(visitor GeometryVisitor) {
	visitor(this)
}

func (this Polygon) Equals(other Geometry) bool {
	switch polygon := other.(type) {
	default:
		return false
	case Polygon:
		if !this.shell.Equals(polygon.shell) {
			return false
		}

		if len(this.holes) != len(polygon.holes) {
			return false
		}
		for i, hole := range this.holes {
			if !hole.Equals(polygon.holes[i]) {
				return false
			}
		}
		return false
	}
}
