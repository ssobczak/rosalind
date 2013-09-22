// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
)

var counters = map[byte]int{
	'A': 4,
	'C': 2,
	'D': 2,
	'E': 2,
	'F': 2,
	'G': 4,
	'H': 2,
	'I': 3,
	'K': 2,
	'L': 6,
	'M': 1,
	'N': 2,
	'R': 6,
	'P': 4,
	'Q': 2,
	'S': 6,
	'T': 4,
	'V': 4,
	'W': 1,
	'Y': 2,
	's': 3,
}

func mrna(protein string, mod int) int {
	ret := counters['s']
	for i := range protein {
		ret = (ret * counters[protein[i]]) % mod
	}
	return ret
}

func Mrna() {
	var protein string
	fmt.Scanf("%s", &protein)
	fmt.Println(mrna(protein, 1000000))
}
