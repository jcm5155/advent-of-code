package solutions

// Day15 solution
func (h *Handler) Day15() (uint32, uint32) {
	// not bothering reading this in since it's so short.
	pzl := []uint32{6, 19, 0, 5, 7, 13, 1}
	var histories = make(map[uint32][]uint32)
	var spoken = make([]uint32, 0, 30000000)
	for i := uint32(1); i <= 30000000; i++ {
		if i > uint32(len(pzl)) {
			var diff uint32
			history := histories[spoken[i-2]]
			if len(history) > 1 {
				diff = history[len(history)-1] - history[len(history)-2]
			}
			histories[diff] = append(histories[diff], i)
			spoken = append(spoken, diff)
		} else {
			histories[pzl[i-1]] = append(histories[pzl[i-1]], i)
			spoken = append(spoken, pzl[i-1])
		}
	}
	return spoken[2019], spoken[29999999]
}
