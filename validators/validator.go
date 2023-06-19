package validators

import "math"

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	maxDivisor := int(math.Sqrt(float64(number)))
	for i := 2; i <= maxDivisor; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
