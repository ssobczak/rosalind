// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"bufio"
	"log"
	"os"
)

func readSeq() (seq string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		seq += scanner.Text()
	}
	return
}

func reverse(seq string) (rev string) {
	for i := len(seq) - 1; i >= 0; i-- {
		rev += string(seq[i])
	}
	return
}

func compliment(seq string) (compl string) {
	for _, v := range seq {
		switch v {
		case 'C':
			compl += "G"
		case 'G':
			compl += "C"
		case 'A':
			compl += "T"
		case 'T':
			compl += "A"
		default:
			log.Fatal("Unknown char ", v, " in sequence ", seq)
		}
	}
	return
}

func findProteins(seq string) {
	ch := make(chan byte)
	dnaChanToProtein(ch)

	for i := range seq {
		ch <- seq[i]
	}
}

func Orf() {
	seq := readSeq()
	rc := reverse(compliment(seq))

	for i := 0; i < 3; i++ {
		findProteins(seq[i : len(seq)-1])
		findProteins(rc[i : len(rc)-1])
	}
}
