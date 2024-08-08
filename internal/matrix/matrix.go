package matrix

func TransposeString(matrix [][]string) [][]string {
	out := make([][]string, len(matrix[0]))

	for i := 0; i < len(matrix); i += 1 {
		for j := 0; j < len(matrix[0]); j += 1 {
			out[j] = append(out[j], matrix[i][j])
		}
	}
	return out
}
