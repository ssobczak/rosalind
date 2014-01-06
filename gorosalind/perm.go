// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
	"github.com/ssobczak/rosalind/permutable"
	"log"
)

func factorial(n int) (ret int) {
	ret = 1

	for i := 2; i <= n; i++ {
		ret *= i
	}

	return
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
	for permutable.NextPermutation(permutable.PermInts(a)) {
		printPerm(a)
	}
}
