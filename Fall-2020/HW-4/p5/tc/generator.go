package main

import (
	"fmt"
	"math"
)

// primes contains all 32-bit primes, which is sufficient for computing the prime factors of a 64-bit number using the trial division method
var primes = primeSieve(1000000008)

func main() {
	fmt.Printf("4 7;125 1000000007;")
	for i := 3; i < 500; i += 3 {
		fmt.Printf("%d %d;", i, primes[i-1])
	}
}

// Modified from Kim Walisch's implementation: https://gist.github.com/kimwalisch/3dc39786fab8d5b34fee#file-segmented_sieve-cpp
func primeSieve(limit int) []int {
	segmentSize := int(math.Sqrt(float64(limit)))

	// Generation of small primes
	notPrime := make([]bool, segmentSize+1)
	for i := 0; i*i < segmentSize && i < 2; i++ {
		notPrime[i] = true
	}

	for i := 2; i*i < segmentSize; i++ {
		if !notPrime[i] {
			for j := i * i; j <= segmentSize; j += i {
				notPrime[j] = true
			}
		}
	}

	sieve := make([]bool, segmentSize)
	var allPrimes []int
	for i := 2; i < len(notPrime); i++ {
		if !notPrime[i] {
			allPrimes = append(allPrimes, i)
		}
	}

	var next []int
	s := 3
	n := segmentSize + 1
	if n%2 == 0 {
		n--
	}

	for low := 0; low <= limit; low += segmentSize {
		for i := 0; i < segmentSize; i++ {
			sieve[i] = true
		}

		// Current segment on the interval [low, high]
		high := low + segmentSize - 1
		if high > limit {
			high = limit
		}

		// Add new sieving primes <= √high
		for ; s*s <= high; s += 2 {
			if !notPrime[s] {
				next = append(next, s*s-low)
			}
		}

		// Sieve the current segment
		for i := 0; i < len(next); i++ {
			j := next[i]
			for k := allPrimes[i+1] * 2; j < segmentSize; j += k {
				sieve[j] = false
			}

			next[i] = j - segmentSize
		}

		// Skip the first set of primes becasue allPrimes already has them
		if low > 0 {
			for ; n <= high; n += 2 {
				if sieve[n-low] {
					allPrimes = append(allPrimes, n)
				}
			}
		}
	}

	return allPrimes
}
