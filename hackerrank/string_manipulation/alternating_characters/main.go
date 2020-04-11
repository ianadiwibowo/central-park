package main

func AlternatingCharacters(s string) int32 {
	var removalCount int32
	var prevChar rune
	for _, r := range s {
		if r == prevChar {
			removalCount++
		}
		prevChar = r
	}

	return removalCount
}

func main() {

}
