package gorosalind

import (
	"fmt"
	"log"
	"math"
)

func combinations(n, k int) (ret uint64) {
	ret = 1

	if n-k > k {
		k = n - k
	}

	for i := n; i > k; i-- {
		ret = ret * uint64(i) / uint64(n-i+1)
	}

	return
}

func mulmul(start, step float64, steps int) float64 {
	for i := 0; i < steps; i++ {
		start *= step
	}

	return start
}

func lia(n, k int) float64 {
	var ret float64 = 1
	const px = 0.25

	n = 1 << uint(n)

	if n-k < k {
		k = n - k
		ret = 0
	}

	for i := 0; i < k; i++ {
		ret -= mulmul(mulmul(float64(combinations(n, i)), px, i), 1-px, n-i)
	}

	return math.Abs(ret)
}

func Lia() {
	n, k := 0, 0

	if _, err := fmt.Scanf("%d %d", &n, &k); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%.40f\n", lia(n, k))
}
