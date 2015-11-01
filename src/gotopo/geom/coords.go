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
	Get(coordIdx uint32) Point
	Set(coordIdx uint32, newValue Point)
	Add(newValue Point)
	Insert(idx uint32, newValue Point)
	// Insert the ordinals at location of idx-th coordinate.  The ordinals can be data for multiple coordinates
	// but must be of the correct length for the number of dimensions of the Coords object.
	// For example if the dimension is 2 then ordinals must be a product of 2 (2,4,6,etc...)
	InsertRaw(idx uint32, ordinals []float64)
}
