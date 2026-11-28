func hammingWeight(n int) int {
	oneBit := 0

	for n>=1 {
		floorLog2 := bits.Len(uint(n)) - 1
		oneBit += 1
		n -= (1 << floorLog2)
	}
	return oneBit
}
