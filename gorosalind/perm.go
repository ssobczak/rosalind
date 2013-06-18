package gorosalind

import (
	"fmt"
	"log"
)

func factorial(n int) (ret int) {
	ret = 1

	for i := 2; i <= n; i++ {
		ret *= i
	}

	return
}

func swap(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

func revert(a []int) {
	for i := 0; i < len(a)/2; i++ {
		swap(&a[i], &a[len(a)-i-1])
	}
}

func next_permutation(a []int) bool {
	for k := len(a) - 2; k >= 0; k-- {
		if a[k] < a[k+1] {
			for l := len(a) - 1; l > k; l-- {
				if a[k] < a[l] {
					swap(&a[k], &a[l])
					revert(a[k+1 : len(a)])
					return true
				}
			}
		}
	}

	return false
}

func printPerm(a []int) {
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
}

func Perm() {
	n := 0

	if _, err := fmt.Scanf("%d", &n); err != nil {
		log.Fatal(err)
	}

	fmt.Println(factorial(n))

	a := make([]int, n)
	for k := range a {
		a[k] = k + 1
	}

	printPerm(a)
	for next_permutation(a) {
		printPerm(a)
	}
}
