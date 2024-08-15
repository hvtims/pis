package piscine

func StrRev(s string) string {
	if len(s) <= 1 {
		return s
	}

	return StrRev(s[1:]) + string(s[0])
}
