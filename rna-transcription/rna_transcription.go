package strand

// ToRNA return the RNA complement string of the given DNA string.
func ToRNA(dna string) string {
	rna := make([]rune, len(dna))

	for i, v := range dna {
		switch v {
		case 'G':
			rna[i] = 'C'
		case 'C':
			rna[i] = 'G'
		case 'T':
			rna[i] = 'A'
		case 'A':
			rna[i] = 'U'
		}
	}

	return string(rna)
}
