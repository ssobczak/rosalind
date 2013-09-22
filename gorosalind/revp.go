// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
)

type RevPal struct {
	start, length int
}

var rev = map[byte]byte{
	'C': 'G',
	'G': 'C',
	'A': 'T',
	'T': 'A',
}

func isRevPalindrome(dna string) bool {
	for i := range dna {
		if dna[i] != rev[dna[len(dna)-i-1]] {
			return false
		}
	}

	return true
}

func revp(dna string) []RevPal {
	res := make([]RevPal, 0)

	for start := 0; start != len(dna); start++ {
		for length := 4; length <= 12 && start+length <= len(dna); length++ {
			if isRevPalindrome(dna[start : start+length]) {
				res = append(res, RevPal{start, length})
			}
		}
	}

	return res
}

func Revp() {
	dna := DnaFromFasta(FileFromStdin())

	for _, r := range revp(dna[0].sequence) {
		fmt.Printf("%d %d\n", r.start+1, r.length)
	}
}

// TCAATGCATGCGGGTCTATATGCAT
