package geom

type MultiPolygon interface {
	Polygons() []Polygon
}