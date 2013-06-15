// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
)

var atoi = map[byte]int{
	'A': 0,
	'C': 1,
	'G': 2,
	'T': 3,
}

var itoa = []int{
	'A',
	'C',
	'G',
	'T',
}

func index_of_max(a []int) (ret int) {
	ret = 0
	for k, v := range a {
		if a[ret] < v {
			ret = k
		}
	}

	return
}

func cons(dna []Dna) (profile [][]int) {
	profile = make([][]int, len(dna[0].sequence))

	for k := range profile {
		profile[k] = make([]int, 4)
		for i := range profile[k] {
			profile[k][i] = 0
		}
	}

	for _, d := range dna {
		for i := range d.sequence {
			profile[i][atoi[d.sequence[i]]]++
		}
	}

	return
}

func Cons() {
	fasta := FileFromStdin()
	dna := DnaFromFasta(fasta)
	profile := cons(dna)

	for _, pr := range profile {
		fmt.Printf("%c", itoa[index_of_max(pr)])
	}
	fmt.Printf("\n")

	for j := range profile[0] {
		fmt.Printf("%c: ", itoa[j])
		for i := range profile {
			fmt.Printf("%d ", profile[i][j])
		}
		fmt.Printf("\n")
	}
}
