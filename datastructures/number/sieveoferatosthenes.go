package number

func SieveOfEratosthenes(limit int) []int {
	isPrimes := make([]bool, limit+1)
	for i := 2; i < len(isPrimes); i++ {
		isPrimes[i] = true
	}

	for i := 2; i^2 <= limit; i++ {
		if !isPrimes[i] {
			continue
		}
		for j := i * 2; j <= limit; j += i {
			isPrimes[j] = false
		}
	}

	primeNumbers := []int{}
	for i := 2; i < len(isPrimes); i++ {
		if isPrimes[i] {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return primeNumbers
}
