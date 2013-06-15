// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"testing"
)

func TestCons_example(t *testing.T) {
	fasta := ">Rosalind_1\n" +
		"ATCCAGCT\n" +
		">Rosalind_2\n" +
		"GGGCAACT\n" +
		">Rosalind_3\n" +
		"ATGGATCT\n" +
		">Rosalind_4\n" +
		"AAGCAACC\n" +
		">Rosalind_5\n" +
		"TTGGAACT\n" +
		">Rosalind_6\n" +
		"ATGCCATT\n" +
		">Rosalind_7\n" +
		"ATGGCACT"

	profile := [][]int{
		{5, 1, 0, 0, 5, 5, 0, 0},
		{0, 0, 1, 4, 2, 0, 6, 1},
		{1, 1, 6, 3, 0, 1, 0, 0},
		{1, 5, 0, 0, 0, 1, 1, 6},
	}

	res := cons(DnaFromFasta(fasta))

	for i := range profile {
		for j := range profile[0] {
			if profile[i][j] != res[j][i] {
				t.Error("oops!", profile, res)
			}
		}
	}
}
