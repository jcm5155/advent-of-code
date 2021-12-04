package common

func ClearBit(n uint, pos uint) uint {
	return n &^ (1 << pos)
}

func IsBitSet(n uint, pos int) bool {
	return n&(1<<pos) > 0
}

func SetBit(n uint, pos int) uint {
	n |= 1 << pos
	return n
}
