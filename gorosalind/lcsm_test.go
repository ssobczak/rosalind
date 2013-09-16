// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"testing"
)

func TestLcsm_example(t *testing.T) {
	fasta := ">Rosalind_1\n" +
		"GATTACA\n" +
		">Rosalind_2\n" +
		"TAGACCA\n" +
		">Rosalind_3\n" +
		"ATACA\n"

	expected := "TA"
	res := lcsm(DnaFromFasta(fasta))

	if expected != res {
		t.Error("oops!", expected, res)
	}
}

// func TestLcsm_Match(t *testing.T) {
// 	l, p := lcsm("qwe", "aqswqwd")

// 	if l != 2 || p != 0 {
// 		t.Error("oops!", l, p)
// 	}

// 	l, p = lcsm("qwe", "aqswqw")

// 	if l != 2 || p != 0 {
// 		t.Error("oops!", l, p)
// 	}

// 	l, p = lcsm("ATACA", "TAGACCA")

// 	if l != 2 || p != 2 {
// 		t.Error("oops!", l, p)
// 	}
// }
