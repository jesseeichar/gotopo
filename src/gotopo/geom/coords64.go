package geom
import (
	"fmt"
)

type point64 struct {
	data  coords64
	index uint32
}

func (this point64) X() float64 {
	return this.data.data[this.index]
}
func (this point64) Y() float64 {
	return this.data.data[this.index + 1]

}
func (this point64) Ord(dimIdx uint8) float64 {
	return this.data.data[this.index + uint32(dimIdx)]
}
func (this point64) NumDim() uint8 {
	return this.data.dimensions
}
func (this point64) ToArray() []float64 {
	endIdx := this.index + uint32(this.data.dimensions)
	return this.data.data[this.index:endIdx]
}
func (this point64) String() string {
	return CoordString(this)
}


// ===============================================================================================================
type coords64 struct {
	data       []float64
	dimensions uint8
}

func NewCoords64() Coords {
	return NewCoords64WithDimensions(DEFAULT_NUM_DIMENSIONS)
}

func NewCoords64WithCapacity(capacity uint32) Coords {
	return NewCoords64WithCapacityAndDimensions(capacity, DEFAULT_NUM_DIMENSIONS)
}
func NewCoords64WithDimensions(dimensions uint8) Coords {
	return NewCoords64WithCapacityAndDimensions(0, dimensions)
}

func NewCoords64WithCapacityAndDimensions(capacity uint32, dimensions uint8) Coords {
	sliceCapacity := capacity * uint32(dimensions)
	if sliceCapacity < uint32(dimensions) {
		return &coords64{
			data:[]float64{},
			dimensions:dimensions}
	} else {
		return &coords64{
			data:make([]float64, 0, sliceCapacity),
			dimensions:dimensions}
	}
}

func NewCoords64FromSlice(dimensions uint8, data []float64) Coords {
	if len(data) % int(dimensions) != 0 {
		panic(fmt.Sprintf("The number of eleements in the data array must be divisible by the number of dimensions." +
		" Array size '%d'.  Dimensions: '%d'", len(data), dimensions))
	}
	return &coords64{
		data:data,
		dimensions:dimensions}
}

func (this coords64) NumDim() uint8 {
	return this.dimensions
}

func (this coords64) NumCoords() uint32 {
	return uint32(len(this.data)) / uint32(this.dimensions)
}

func (this coords64) Get(coordIdx uint32) Point {
	return point64{this, coordIdx * uint32(this.dimensions)}
}
func (this *coords64) Set(coordIdx uint32, newValue Point) {
	if newValue.NumDim() != this.dimensions {
		panic(fmt.Sprintf("Number of dimensions in coordinate(%d) do not match those in this coords object (%d)",
			newValue.NumDim(), this.dimensions))
	}

	setIdx := coordIdx * uint32(this.dimensions)

	if setIdx > uint32(len(this.data)) - uint32(this.dimensions) {
		panic(fmt.Sprintf("Insert index is out of bounds.  Legal bounds are: 0 -> %d but was %d", this.NumCoords(), coordIdx))
	}

	for i := uint8(0); i < newValue.NumDim(); i++ {
		this.data[setIdx + uint32(i)] = newValue.Ord(i)
	}
}
func (this *coords64) Add(newValue Point) {
	if newValue.NumDim() != this.dimensions {
		panic(fmt.Sprintf("Number of dimensions in coordinate(%d) do not match those in this coords object (%d)",
			newValue.NumDim(), this.dimensions))
	}
	this.data = append(this.data, newValue.ToArray()...)
}
func (this *coords64) Insert(idx uint32, newValue Point) {
	this.InsertRaw(idx, newValue.ToArray())
}
func (this *coords64) InsertRaw(idx uint32, ordinals []float64) {
	mod := len(ordinals) % int(this.dimensions)

	if mod != 0 {
		panic(fmt.Sprintf("The number of ordinals provided for insert must be exactly divisible by the number of " +
		"dimensions but ordinals: %d is not divisible by %d, there is a remainder of: %d", len(ordinals), this.dimensions, mod))
	}

	insertIdx := idx * uint32(this.dimensions)

	if insertIdx > uint32(len(this.data)) {
		panic(fmt.Sprintf("Insert index is out of bounds.  Legal bounds are: 0 -> %d but was %d", this.NumCoords(), idx))
	}

	this.data = append(this.data, ordinals...)
	copy(this.data[insertIdx + uint32(this.dimensions):], this.data[insertIdx:])
	for i, o := range ordinals {
		this.data[insertIdx + uint32(i)] = o
	}
}
