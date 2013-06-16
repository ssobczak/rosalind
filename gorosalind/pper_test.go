// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"testing"
)

func pper_test(n, k, expected int, t *testing.T) {
	ret := pper(n, k)
	if ret != expected {
		t.Error("Expected ", expected, ", got", ret)
	}
}

func TestPper_example(t *testing.T) {
	pper_test(21, 7, 51200, t)
}

func TestPper_big(t *testing.T) {
	pper_test(95, 8, 926400, t)
}
