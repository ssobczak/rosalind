// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"testing"
)

func TestGrph_example(t *testing.T) {
	fasta := ">Rosalind_0498\n" +
		"AAATAAA\n" +
		">Rosalind_2391\n" +
		"AAATTTT\n" +
		">Rosalind_2323\n" +
		"TTTTCCC\n" +
		">Rosalind_0442\n" +
		"AAATCCC\n" +
		">Rosalind_5013\n" +
		"GGGTGGG"

	edges := map[string]bool{
		"Rosalind_0498 Rosalind_2391": true,
		"Rosalind_0498 Rosalind_0442": true,
		"Rosalind_2391 Rosalind_2323": true,
	}

	res := grph(DnaFromFasta(fasta), 3)

	for _, e := range res.e {
		key := e.from.name + " " + e.to.name
		if _, present := edges[key]; present {
			delete(edges, key)
		} else {
			t.Error("unexpected: ", key)
		}
	}

	if len(edges) != 0 {
		t.Error("expected, but not found: ", edges)
	}
}
