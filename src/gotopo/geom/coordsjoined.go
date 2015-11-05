package geom
import (
	"fmt"
)

type coordsJoined struct {
	allCoords []Coords
}

var _ Coords = coordsJoined{}       // Verify that coordsJoined implements Coords.
var _ Coords = (*coordsJoined)(nil) // Verify that *coordsJoined implements Coords

func NewCoordsJoined(allCoords ...Coords) coordsJoined {
	if len(allCoords) < 1 {
		panic("At least one element is required in the slice of coords")
	}
	verifyConsistentNumDims(allCoords)
	return coordsJoined{allCoords}
}

func verifyConsistentNumDims(allCoords []Coords) {
	dim := map[uint8]bool{}
	for _, c := range allCoords {
		dim[c.NumDim()] = true;
	}
	if len(dim) != 1 {
		allDims := []int{}
		for d := range dim {
			allDims = append(allDims, int(d))
		}
		panic(fmt.Sprintf("All coordinate objects must have the same number of dimensions.  Found: %v", allDims))
	}
}


func (this coordsJoined) NumDim() uint8 {
	return this.allCoords[0].NumDim()
}
func (this coordsJoined) NumCoords() uint32 {
	count := uint32(0)
	for _, c := range this.allCoords {
		count += c.NumCoords()
	}
	return count
}
func (this coordsJoined) IsEmpty() bool {
	for _, c := range this.allCoords {
		if !c.IsEmpty() {
			return false;
		}
	}
	return true
}
func (this coordsJoined) Get(coordIdx uint32) Point {
	count := uint32(0)
	for _, c := range this.allCoords {
		if count + c.NumCoords() > coordIdx {
			return c.Get(coordIdx - count)
		}
		count += c.NumCoords()
	}

	panic(fmt.Sprintf("%d is out of bounds. There are only %s coordinates total", coordIdx, count))
}

func (this coordsJoined) Factory() CoordsFactory {
	return this.allCoords[0].Factory()
}
