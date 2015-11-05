package op
import (
	"gotopo/geom"
	"gotopo/util"
)

const (
	INSIDE_AREA = -1
)
type geomloc struct {
	geom     geom.Geometry
	segIndex int
	pt       geom.Point
}

type distanceOp struct {
	minDistance, terminateDistance float64
}

func Distance(geom1, geom2 geom.Geometry) {
	geomLocs := [2]geomloc{}
}

func computeContainmentDistance(result distanceOp, geom1, geom2 geom.Geometry, geomLocs [2]geomloc) {
	computeContainmentDistanceForGeom(result, geom1, geomLocs)
	computeContainmentDistanceForGeom(result, geom2, geomLocs)
}
func computeContainmentDistanceForGeom(result distanceOp, geom geom.Geometry, geomLocs [2]geomloc) {
	polys := util.ExtractPolygons(geom)
	if len(polys) > 0 {
		insideLocs := getGeometryLocations(geom)

		computeContainmentDistanceFromListOfLocsAndPolygons(result, insideLocs, polys, geomLocs);
		computeContainmentDistanceFromListOfLocsAndPolygons(result, insideLocs, polys, geomLocs);
	}
}

func computeContainmentDistanceFromListOfLocsAndPolygons(result distanceOp, locs []geomloc, polys []geom.Polygon, geomLocs [2]geomloc) {
	for loc := range locs {
		for poly := range polys {
			computeContainmentDistance(result, loc, poly, geomLocs);
			if result.minDistance <= result.terminateDistance {
				return
			}
		}
	}
}

func computeContainmentDistanceToPolygon(result distanceOp, loc geomloc, poly geom.Polygon, geomLocs [2]geomloc) {
	// if pt is not in exterior, distance to geom is 0
	if Location.EXTERIOR != ptLocator.locate(loc.pt, poly) {
		result.minDistance = 0.0
		geomLocs[0] = loc
		geomLocs[1] = geomloc{poly, INSIDE_AREA, loc.pt}
	}
}

func getGeometryLocations(geom geom.Geometry) []geomloc {
	locations := []geomloc{}
	geom.Visit(func(geom) {
		switch geom.(type) {
		case geom.Line:
		case geom.Polygon:
		case geom.Point:
			append(locations, geomloc{geom.Coords().Get(0), 0, geom})
		}
	})
	return locations
}
