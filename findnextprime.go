package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	for I := nb; ; I++ {
		if IsPrime(I) {
			return I
		}
	}
}
