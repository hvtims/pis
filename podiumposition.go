package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium)-1; i++ {
		for range podium {
			if podium[i][0] < podium[i+1][0] {
				a := podium[i]
				podium[i] = podium[i+1]
				podium[i+1] = a
			}
		}
	}

	return podium
}
