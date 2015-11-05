package util

func MinUint8(i1, i2 uint8) uint8 {
	if i1 > i2 {
		return i2
	} else {
		return i1
	}
}