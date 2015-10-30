package geom
const (
	DEFAULT_NUM_DIMENSIONS = uint8(2)
)
/*
An interface to allow access to the underlying coordinates of a geometry.

In order to improve performance or
 */
type Coords interface {
	NumDim() uint8
	NumCoords() uint32
	Get(coordIdx uint32) Coord
	Set(coordIdx uint32, newValue Coord) error
	Add(newValue Coord) error
	Insert(idx uint32, newValue Coord) error
	// Insert the ordinals at location of idx-th coordinate.  The ordinals can be data for multiple coordinates
	// but must be of the correct length for the number of dimensions of the Coords object.
	// For example if the dimension is 2 then ordinals must be a product of 2 (2,4,6,etc...)
	InsertRaw(idx uint32, ordinals []float64) error
}

type Coord interface {
	X() float64
	Y() float64
	Ord(dimIdx uint8) float64
	NumDim() uint8
	ToArray() []float64
}

func NewCoord(ordinals ...float64) Coord {
	coords := NewCoords64WithCapacityAndDimensions(1, uint8(len(ordinals)))
	coords.InsertRaw(0, ordinals)
	return coords.Get(0);
}