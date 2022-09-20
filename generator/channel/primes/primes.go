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

// GetFirstN return the next prime number in the sequence
func GetFirstN(n int) chan int {
	s := make(chan int, n)

	go func() {
		defer close(s)

		var nsent int
		for i := 2; nsent < n && i <= max; i++ {
			if isPrime[i] {
				s <- i
				nsent++
			}
		}
	}()
	return s
}
