// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
	"github.com/ssobczak/rosalind/permutable"
	"log"
)

func signPerm(a []int) {
	for i := 0; i != 1<<uint8(len(a)); i++ {
		for k, v := range a {
			if 1<<uint8(k)&i != 0 {
				fmt.Printf("%d ", v)
			} else {
				fmt.Printf("%d ", -v)
			}
		}
		fmt.Print("\n")
	}
}

func sign(n int) {
	fmt.Println(factorial(n) * (1 << uint8(n)))

	a := make([]int, n)
	for k := range a {
		a[k] = k + 1
	}

	signPerm(a)
	for permutable.NextPermutation(permutable.PermInts(a)) {
		signPerm(a)
	}
}

func Sign() {
	n := 0

	if _, err := fmt.Scanf("%d", &n); err != nil {
		log.Fatal(err)
	}

	sign(n)
}
