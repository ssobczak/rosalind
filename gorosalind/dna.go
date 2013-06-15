// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"strings"
)

type Dna struct {
	name     string
	sequence string
}

func DnaFromFasta(fasta string) (ret []Dna) {
	ret = make([]Dna, 0)

	lines := strings.Split(fasta, "\n")
	for _, line := range lines {
		if line[0] == '>' {
			ret = append(ret, Dna{line[1 : len(line)-1], ""})
		} else {
			ret[len(ret)-1].sequence += line
		}
	}

	return
}
