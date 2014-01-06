// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"math"
)

func gc_count(s string) int {
	res := 0

	for _, c := range s {
		if c == 'C' || c == 'G' {
			res++
		}
	}

	return res
}

func rstr(n int, gc_c float64, motif string) float64 {
	gc := gc_count(motif)
	at := len(motif) - gc

	p_gc := math.Pow(gc_c/2, float64(gc))
	p_at := math.Pow((1-gc_c)/2, float64(at))

	return 1 - math.Pow(1-p_gc*p_at, float64(n))
}
