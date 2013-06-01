// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
	"log"
)

func wabbits(n, m int) uint64 {
	young := make([]uint64, m)
	young[0] = 1
	for i := 1; i < m; i++ {
		young[i] = 0
	}

	var old uint64 = 0

	for i := 1; i < n; i++ {
		tmp := young[i%m]
		young[i%m] = old
		old += young[(m+i-1)%m] - tmp
	}

	return old + young[(n-1)%m]
}

func Fibd() {
	var n, m int
	if _, err := fmt.Scanf("%d %d", &n, &m); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", wabbits(n, m))
}
