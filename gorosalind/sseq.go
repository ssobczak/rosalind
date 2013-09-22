// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
)

func sseq(what, where string) []int {
	res := make([]int, 1)
	res[0] = 0

	for _, l := range what {
		for i := res[len(res)-1]; i != len(where); i++ {
			if where[i] == uint8(l) {
				res = append(res, i+1)
				break
			}
		}
	}

	return res[1:len(res)]
}

func Sseq() {
	dna := DnaFromFasta(FileFromStdin())

	for _, v := range sseq(dna[1].sequence, dna[0].sequence) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
