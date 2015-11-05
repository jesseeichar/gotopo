package coords64
import "gotopo/geom"

type CoordsFactory64 struct {
	dim uint8
}
var _ geom.CoordsFactory = CoordsFactory64{}   // Verify that CoordsFactory64 implements geom.CoordsFactory
var _ geom.CoordsFactory = (*CoordsFactory64)(nil)   // Verify that CoordsFactory64 implements geom.CoordsFactory

func (this CoordsFactory64) Create() geom.ReadWriteCoords {
	return NewCoords64WithDimensions(this.dim)
}
func (this CoordsFactory64) CreateFromRawData(data ...float64) geom.ReadWriteCoords {
	return NewCoords64FromSlice(this.dim, data)
}
func (this CoordsFactory64) CreateFromPoints(points ...geom.Point) geom.ReadWriteCoords {
	coords := NewCoords64()
	for i, pt := range points {
		coords.InsertRaw(uint32(i), pt.ToArray()[:this.dim])
	}
	return coords
}
func (this CoordsFactory64) CopyFromGeometry(geom geom.Geometry) geom.ReadWriteCoords {
	coords := NewCoords64()
	geomCoords := geom.Coords()
	for i := uint32(0); i < geomCoords.NumCoords(); i++ {
		coords.InsertRaw(i, geomCoords.Get(i).ToArray()[:this.dim])
	}
	return coords
}