// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"math"
	"testing"
)

func TestEval_example(t *testing.T) {
	n, s, gc_c := 10, "AG", []float64{0.25, 0.5, 0.75}
	profile := []float64{0.422, 0.563, 0.422}

	for i, gc := range gc_c {
		if ret := eval(n, gc, s); math.Abs(ret-profile[i]) > 0.001 {
			t.Error("eval(%f, %f, %s) = %f, expected %f", n, gc, s, ret, profile[i])
		}
	}
}
