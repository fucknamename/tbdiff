package utils

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func MaxUInt8(a, b uint8) uint8 {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MinUInt8(a, b uint8) uint8 {
	if a > b {
		return b
	}
	return a
}
