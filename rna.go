package strand

func ToRNA(dna string) string {
	// panic("Please implement t÷he ToRNA function")
	runes := []rune(dna)
	for i := 0; i < len(dna); i++ {
		if runes[i] == 'A' {
			runes[i] = 'U'
		} else if runes[i] == 'C' {
			runes[i] = 'G'
		} else if runes[i] == 'G' {
			runes[i] = 'C'
		} else if runes[i] == 'T' {
			runes[i] = 'A'
		}
	}
	ans := string(runes)
	return ans
}
