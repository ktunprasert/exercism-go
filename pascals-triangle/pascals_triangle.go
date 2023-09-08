package pascal

func Triangle(n int) [][]int {
	out := make([][]int, n)

	for i := range out {
		out[i] = make([]int, i+1)
		out[i][0], out[i][i] = 1, 1

		for j := 1; j < i; j++ {
			out[i][j] = out[i-1][j-1] + out[i-1][j]
		}
	}

	return out
}
