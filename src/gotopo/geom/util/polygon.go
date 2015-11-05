package util
import "gotopo/geom"

func ExtractPolygons(src geom.Geometry) []geom.Polygon {

	switch typedGeom := src.(type) {
	case geom.Polygon:
		return []geom.Polygon{typedGeom}
	case geom.GeometryCollection:
		polygons := make([]geom.Polygon, 0, len(typedGeom.Geometries()))

		for _, geom2 := range typedGeom.Geometries() {
			polygons = append(polygons, ExtractPolygons(geom2)...)
		}
		return polygons
	}
	return []geom.Polygon{}
}
