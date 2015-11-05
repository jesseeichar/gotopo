package op
import "gotopo/geom"

type locationEnum int
const (
	location_INTERIOR locationEnum = iota
	location_BOUNDARY
	location_NONE
	location_EXTERIOR
)

type locationResult struct {
	isIn          bool
	numBoundaries int
}

func locate(p geom.Point, geom geom.Geometry) uint8 {
	if len(geom.Coords()) == 0 {
		return location_EXTERIOR
	}

	switch typedGeom := geom.(type) {
	case geom.Line:
		return locateOnLine(p, typedGeom)
	case geom.Polygon:
		return locateOnPolygon(p, typedGeom)
	}

	result := locationResult{}
	computeLocation(result, p, geom)
	if boundaryRule.isInBoundary(result.numBoundaries) {
		return location_BOUNDARY
	}
	if result.numBoundaries > 0 || result.isIn {
		return location_INTERIOR
	}

	return location_EXTERIOR
}

func computeLocation(result locationResult, p geom.Point, geom geom.Geometry) {
	switch typedGeom := locationResult.(type) {
	case geom.Point:
		updateLocationResult(locateOnPoint(p, typedGeom))
	case geom.Line:
		updateLocationResult(locateOnLine(p, typedGeom))
	case geom.Polygon:
		updateLocationResult(locateOnPolygon(p, typedGeom))
	case geom.MultiLine:
		for line := range typedGeom.Lines() {
			updateLocationResult(locateOnLine(p, line))
		}
	case geom.MultiPolygon:
		updateLocationResult(locateOnPolygon(p, typedGeom))
	case geom.GeometryCollection:
		for geom := range typedGeom {
			computeLocation(result, p, geom)
		}
	}
}

func updateLocationResult(result locationResult, loc locationEnum) {
	switch loc {
	case location_INTERIOR:
		result.isIn = true
	case location_BOUNDARY:
		result.numBoundaries++
	}
}

func locateOnPolygon(pt geom.Point, poly geom.Polygon) locationEnum {
	if poly.Coords().IsEmpty() {
		return location_EXTERIOR
	}

	shell := poly.Shell()

	shellLoc := locateInPolygonRing(pt, shell)
	if shellLoc == location_EXTERIOR {
		return location_EXTERIOR
	}
	if shellLoc == location_BOUNDARY {
		return location_BOUNDARY
	}
	// now test if the point lies in or on the holes
	for i := 0; i < poly.NumHoles(); i++ {
		hole := poly.Hole(i)
		holeLoc := locateInPolygonRing(pt, hole)
		if holeLoc == location_INTERIOR {
			return location_EXTERIOR
		}
		if holeLoc == location_BOUNDARY {
			return location_BOUNDARY
		}
	}
	return location_INTERIOR
}

func locateInPolygonRing(pt geom.Point, ring geom.LinearRing) locationEnum {
	if !ring.Bounds().intersects(p)) return Location.EXTERIOR;

	return CGAlgorithms.locatePointInRing(p, ring.getCoordinates());

}