// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
	"sort"
)

type DnaArray []Dna

func (a DnaArray) Len() int      { return len(a) }
func (a DnaArray) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByLength struct{ DnaArray }

func (a ByLength) Less(i, j int) bool {
	return len(a.DnaArray[i].sequence) < len(a.DnaArray[j].sequence)
}

type subseq struct {
	start, cnt int
}

func is_subseq(start, cnt int, dnas DnaArray, res_chan chan subseq) {
	for i := 1; i != len(dnas); i++ {
		dna_matching := false
		for m_start := 0; m_start != len(dnas[i].sequence)-cnt+1; m_start++ {
			matching := true
			for j := 0; j != cnt; j++ {
				if dnas[0].sequence[start+j] != dnas[i].sequence[m_start+j] {
					matching = false
					break
				}
			}
			if matching {
				dna_matching = true
				break
			}
		}
		if !dna_matching {
			res_chan <- subseq{0, 0}
			return
		}
	}

	res_chan <- subseq{start, cnt}
}

func lcsm(dnas DnaArray) string {
	sort.Sort(ByLength{dnas})

	res_chan := make(chan subseq)
	defer close(res_chan)

	res_cnt := 0
	for start := 0; start < len(dnas[0].sequence); start++ {
		for cnt := 1; cnt < len(dnas[0].sequence)-start+1; cnt++ {
			go is_subseq(start, cnt, dnas, res_chan)
			res_cnt++
		}
	}

	best := subseq{0, 0}
	for i := 0; i != res_cnt; i++ {
		res := <-res_chan
		if best.cnt < res.cnt {
			best = res
		}
	}

	return dnas[0].sequence[best.start : best.start+best.cnt]
}

func Lcsm() {
	fasta := FileFromStdin()
	dna := DnaFromFasta(fasta)

	fmt.Println(lcsm(dna))
}
