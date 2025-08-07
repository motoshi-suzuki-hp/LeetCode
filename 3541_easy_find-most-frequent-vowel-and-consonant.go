func maxFreqSum(s string) int {
    vowels := map[string]int{"a": 0, "e": 0, "i": 0, "o": 0, "u": 0}
    consonants := make(map[string]int)

    for _, c := range s {
        cStr := string(c)
        if _, vowelExist := vowels[cStr]; vowelExist {
            vowels[cStr]++
        } else {
            if _, consonantsExist := consonants[cStr]; consonantsExist {
                consonants[cStr]++
            } else {
                consonants[cStr] = 1
            }
        }
    }

    var mostFrequentVowel int
    for _, value := range vowels {
		if value > mostFrequentVowel {
            mostFrequentVowel = value
        }
	}

    var mostFrequentConsonant int
    for _, value := range consonants {
		if value > mostFrequentConsonant {
            mostFrequentConsonant = value
        }
	}

    return mostFrequentVowel + mostFrequentConsonant
}
