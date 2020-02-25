package hamming

import "fmt"

// Distance takes two strings and returns the hamming distance between them.
func Distance(a, b string) (int, error) {
	// default distance to 0
	distance := 0
	// default err as an error type set to nil
	var err error = nil

	// convert to runes so that utf8 characters can be compared.
	runesA := []rune(a)
	runesB := []rune(b)
	fmt.Println(runesA, runesB)
	// if lengths dont match throw an error.
	if len(runesA) != len(runesB) {
		err = fmt.Errorf("Distance(%q, %q), string lengths are not equal so this is not a valid input for this function", a, b)
	} else if a != b {
		// loop over the runes in runesA and compare the same index in runesB
		// if it does not match increment distance
		for index, runeValueA := range runesA {
			if runeValueA != runesB[index] {
				distance++
			}
		}
	}

	return distance, err
}
