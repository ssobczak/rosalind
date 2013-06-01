// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
	"log"
)

func fib(n, m int) uint64 {
	var old, young uint64 = 0, 1

	for i := 1; i < n; i++ {
		tmp := old
		old += young
		young = tmp
	}

	return old
}

func Fib() {
	var n, m int
	if _, err := fmt.Scanf("%d %d", &n, &m); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", fib(n, m))
}
