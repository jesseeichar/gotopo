package geom
import (
	"gotopo/util"
)

type Bounds struct {
	min Point
	max Point
}

func (this Bounds) Intersects(pt Point) bool {
	intersects := true
	dims := util.MinUint8(pt.NumDim(), this.max.NumDim())
	for dim := uint8(0); dim < dims; dim++ {
		ord := pt.Ord(dim)
		intersects = intersects && ord > this.min.Ord(dim) && ord < this.max.Ord(dim)
	}
	return intersects
}

func (this *Bounds) expandToInclude(pt Point) {
	dims := util.MinUint8(pt.NumDim(), this.max.NumDim())
	for dim := uint8(0); dim < dims; dim++ {
		ord := pt.Ord(dim)
		if (ord < this.min.Ord(dim)) {
			this.min = derivePoint(dim, ord, this.min)
		} else if ord > this.max.Ord(dim) {
			this.max = derivePoint(dim, ord, this.max)
		}
	}
}

func derivePoint(ordIdx uint8, newValue float64, original Point) Point {
	ptData := make([]float64, original.NumDim(), original.NumDim())
	for i := uint8(0); i < original.NumDim(); i++ {
		if i == ordIdx {
			ptData[i] = newValue
		} else {
			ptData[i] = original.Ord(i)
		}
	}
	coords := original.Coords().Factory().CreateFromRawData(ptData...)

	return coords.Get(0)
}