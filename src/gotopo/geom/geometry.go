package geom

type Geometry interface {
	Coords() Coords
	Visit(visitor GeometryVisitor)
	Equals(Geometry) bool
}

type GeometryVisitor func(Geometry)