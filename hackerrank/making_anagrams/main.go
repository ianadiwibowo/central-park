package main

func MakeAnagram(a, b string) int32 {
	charmap := make(map[rune]int32)

	for _, c := range a {
		charmap[c]++
	}

	for _, c := range b {
		charmap[c]--
	}

	var result int32
	for _, value := range charmap {
		result += absolute(value)
	}

	return result
}

func absolute(number int32) int32 {
	if number < 0 {
		return -number
	}

	return number
}

func main() {

}
