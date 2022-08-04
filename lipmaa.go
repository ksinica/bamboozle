package bamboozle

func Lipmaa(n uint64) uint64 {
	m, po3, x := uint64(1), uint64(3), n
	for m < n {
		po3 *= 3
		m = (po3 - 1) / 2
	}
	po3 /= 3
	if m != n {
		for x != 0 {
			m = (po3 - 1) / 2
			po3 /= 3
			x %= m
		}
		if m != po3 {
			po3 = m
		}
	}
	return n - po3
}
