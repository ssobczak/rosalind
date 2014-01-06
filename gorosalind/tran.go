// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
)

func isTransition(a, b byte) bool {
	q := a == 'A' && b == 'G'
	w := a == 'G' && b == 'A'
	e := a == 'C' && b == 'T'
	r := a == 'T' && b == 'C'

	return q || w || e || r
}

func Tran() {
	dna := DnaFromFasta(FileFromStdin())
	transitions, transversions := 0, 0

	for i, _ := range dna[0].sequence {
		if dna[0].sequence[i] != dna[1].sequence[i] {
			if isTransition(dna[0].sequence[i], dna[1].sequence[i]) {
				transitions += 1
			} else {
				transversions += 1
			}
		}
	}

	fmt.Println(float64(transitions) / float64(transversions))
}
