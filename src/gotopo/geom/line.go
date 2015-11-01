package geom

type Line struct {
	coordinates Coords
}

func NewLine(coordinates Coords) Line {
	if coordinates.NumCoords() < 2 {
		panic("At least two coordinates are needed to make a line")
	}
	return Line{coordinates}
}

