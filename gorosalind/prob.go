// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
	"math"
)

func prob1(dna string, cg_ratio float64) float64 {
	res := 0.0
	lcg, lat := math.Log10(cg_ratio/2), math.Log10((1.0-cg_ratio)/2)

	for i := range dna {
		if dna[i] == 'C' || dna[i] == 'G' {
			res += lcg
		} else {
			res += lat
		}
	}

	return res
}

func Prob(dna string, cg_ratios []float64) {
	for _, r := range cg_ratios {
		fmt.Printf("%.3f ", prob1(dna, r))
	}
	fmt.Printf("\n")
}
