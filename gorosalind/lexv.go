// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
	"github.com/ssobczak/rosalind/permutable"
	"math"
	"sort"
	"strconv"
	"strings"
)

const MAXPROCS = 4

func sum(s string) int {
	sum := 0
	for _, cnt := range s {
		sum += int(cnt - '0')
	}
	return sum
}

func gen_selectors(alphabet string, start, end, length int, res chan string, counts chan int) {
	count := 0
	for selector := start; selector < end; selector++ {
		str_selector := strconv.FormatInt(int64(selector), length+1)
		if sum(str_selector) <= length && selector != 0 {
			res <- str_selector
			count++
		}
	}
	counts <- count
}

type LexSort struct {
	alphabet string
	content  []string
}

func (l LexSort) Len() int      { return len(l.content) }
func (l LexSort) Swap(i, j int) { l.content[i], l.content[j] = l.content[j], l.content[i] }
func (l LexSort) Less(i, j int) bool {
	if len(l.content[i]) < len(l.content[j]) {
		return !l.Less(j, i)
	}

	for k := 0; k != len(l.content[i]); k++ {
		if len(l.content[j]) == k {
			return false
		}

		if l.content[i][k] != l.content[j][k] {
			return strings.Index(l.alphabet, string(l.content[i][k])) <
				strings.Index(l.alphabet, string(l.content[j][k]))
		}
	}

	return false
}

func (l LexSort) Print() {
	for _, s := range l.content {
		fmt.Println(s)
	}
}

func lexv(alphabet string, length int) {
	var (
		a []uint8
	)

	selectors := make(chan string, 10)
	counts := make(chan int)

	max := int(math.Pow(float64(length+1), float64(len(alphabet))))
	for i := 0; i != MAXPROCS; i++ {
		go gen_selectors(alphabet, i*(max/MAXPROCS), (i+1)*(max/MAXPROCS), length, selectors, counts)
	}

	count, counts_rcvd, processed := 0, 0, 0
	res := make([]string, 0)

	for {
		select {
		case selector := <-selectors:
			a = make([]uint8, 0)

			for j := 0; j != len(selector); j++ {
				for k := 0; k != int(selector[j]-'0'); k++ {
					a = append(a, alphabet[len(selector)-j-1])
				}
			}

			sort.Sort(permutable.PermChars(a))

			for {
				res = append(res, string(a))
				if !permutable.NextPermutation(permutable.PermChars(a)) {
					break
				}
			}

			processed++
		case c := <-counts:
			count += c
			counts_rcvd++

			if counts_rcvd == MAXPROCS {
				close(counts)
				counts = nil
			}
		}

		if counts_rcvd == MAXPROCS && count == processed {
			close(selectors)
			break
		}
	}

	ls := LexSort{alphabet, res}
	sort.Sort(ls)
	ls.Print()
}

func Lexv() {
	lexv("HKIVWJSEZURQ", 3)
}
