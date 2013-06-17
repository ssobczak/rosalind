// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

// This is a prallel algorithm for processing the sequence and producing the results ONLINE.
// I'm proud of it even tough it's unfinished. But it's the first one of this kind written by me.
// TODO(ssobczak) Implemenr processing of reversed complimentary sequence

package gorosalind

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const CODON_LEN = 3

func cloneChannel(ch chan byte) (r1, r2 chan byte) {
	r1, r2 = make(chan byte), make(chan byte)
	go func() {
		for b := range ch {
			r1 <- b
			r2 <- b
		}
		close(r1)
		close(r2)
	}()
	return
}

func makeParentAndSibling(ch chan byte) (parent, sibling chan byte) {
	parent, sibling = make(chan byte), make(chan byte)
	go func() {
		for b := range parent {
			ch <- b
			sibling <- b
		}
		close(ch)
		close(sibling)
	}()
	return
}

func groundChannel(ch chan byte) {
	go func() {
		for ok := true; ok; _, ok = <-ch {
		}
	}()
}

func dnaChanToProtein(in chan byte) {
	out := make(chan byte)
	groundChannel(out)

	buf := make([]byte, CODON_LEN)

	go func() {
		i := 0
		var buildingProteinChannel chan byte

		for b := range in {
			buf[i] = b
			i++

			if i == CODON_LEN {
				i = 0
				if isStop(buf) {
					close(out)
					out = make(chan byte)
					groundChannel(out)
				} else {
					if isStart(buf) {
						out, buildingProteinChannel = makeParentAndSibling(out)
						go buildProtein(buildingProteinChannel)
					}

					if protein, ok := dnaCodonToProtein(buf); ok {
						out <- protein
					} else {
						log.Fatal("Codon not found: " + string(buf))
					}
				}
			}
		}
		close(out)
	}()
	return
}

func print(ch chan byte) {
	for b := range ch {
		fmt.Printf("%c\n", b)
	}
}

func buildProtein(ch chan byte) {
	var protein string
	for b := range ch {
		protein += string(b)
	}
	fmt.Println(protein)
}

func OrfOnline() {
	channels := make([]chan byte, CODON_LEN)
	chClones := make([]chan byte, CODON_LEN)

	chClones[0] = make(chan byte)
	for i := 1; i < CODON_LEN; i++ {
		channels[i-1], chClones[i] = cloneChannel(chClones[i-1])
	}
	channels[CODON_LEN-1] = chClones[CODON_LEN-1]

	for i := range channels {
		dnaChanToProtein(channels[i])
	}

	maxFed := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		seq := scanner.Text()

		for i := 0; i < len(seq); i++ {
			chClones[CODON_LEN-maxFed-1] <- seq[i]

			if maxFed < CODON_LEN-1 {
				maxFed++
			}
		}
	}

	// closing the channels will trigger flush of unterminated transcriptions
	// close(chClones[0])
}
