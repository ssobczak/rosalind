// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
)

type Combination struct {
	chars        string
	gen          []int
	curr, length int
}

func NewCombination(chars string, length int) Combination {
	return Combination{chars, make([]int, length), 0, length}
}

func (p *Combination) To_s() string {
	ret := make([]byte, p.length)

	for pos, i := range p.gen {
		ret[p.length-pos-1] = p.chars[i]
	}

	return string(ret)
}

func (p *Combination) Next() bool {
	if p.gen[p.curr]+1 != len(p.chars) {
		p.gen[p.curr]++
		return true
	}

	for p.gen[p.curr] == len(p.chars)-1 {
		p.curr++
		if p.curr == p.length {
			return false
		}
	}
	p.gen[p.curr]++
	for p.curr--; p.curr != 0; p.curr-- {
		p.gen[p.curr] = 0
	}
	p.gen[p.curr] = 0

	return true
}

func lexf(chars string, length int) {
	p := NewCombination(chars, length)
	fmt.Println(p.To_s())

	for p.Next() {
		fmt.Println(p.To_s())
	}
}
