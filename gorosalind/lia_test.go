// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"math"
	"testing"
)

func TestLia_example(t *testing.T) {
	const n, k, out = 2, 1, 0.684

	if ret := lia(n, k); math.Abs(ret-out) > 0.001 {
		t.Error("lia(%d, %d) = %f, expected %f", n, k, ret, out)
	}
}

func TestLia(t *testing.T) {
	const n, k, out = 6, 17, 0.433352

	if ret := lia(n, k); math.Abs(ret-out) > 0.001 {
		t.Error("lia(%d, %d) = %f, expected %f", n, k, ret, out)
	}
}
