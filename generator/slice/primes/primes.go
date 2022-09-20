package primes

const max = 10000

var isPrime []bool

func init() {
	isPrime = make([]bool, max+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i*i <= max; i++ {
		if isPrime[i] {
			for j := i * 2; j <= max; j += i {
				isPrime[j] = false
			}
		}

	}
}

// GetFirstN return the first N prime numbers
func GetFirstN(n int) []int {
	s := []int{}
	for i := 2; len(s) < n; i++ {
		if isPrime[i] {
			s = append(s, i)
		}
	}
	return s
}
