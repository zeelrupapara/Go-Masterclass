package prime

// Prime Numbers Generator: Generate all prime numbers up to N using different approaches (e.g., Sieve of Eratosthenes, basic iteration)
func Prime(n int) []int {
	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Prime numbers properties
// 1. All prime numbers are greater than 1.
// 2. Prime numbers are divisible by themselves and 1 only.
// 3. Prime numbers are not divisible by any other number than themselves and 1.