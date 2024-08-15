package piscine

func Swap(a *int, b *int) {
	flak := *a
	*a = *b
	*b = flak
}
