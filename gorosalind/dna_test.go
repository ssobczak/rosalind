// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"testing"
)

func TestDna_parse(t *testing.T) {
	dna := DnaFromFasta(">first\nAAATAAA\n>second\nAAATTTT")
	if len(dna) != 2 {
		t.Error("length doesn't match")
	}
}
