package strand

func ToRNA(dna string) string {
	var rna string

	for _, v := range dna {
		switch v {
		case 'G':
			rna += "C"
		case 'C':
			rna += "G"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}

	return rna
}
