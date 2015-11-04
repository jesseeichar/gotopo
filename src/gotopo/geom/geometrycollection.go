package geom

type GeometryCollection interface {
	Geometries() []Geometry
}