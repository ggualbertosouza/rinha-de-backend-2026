package fraud

func Distance(
	a Vector,
	b Vector,
) float32 {
	var sum float32

	for i := 0; i < 14; i++ {

		diff := a[i] - b[i]

		sum += diff * diff
	}

	return sum
}
