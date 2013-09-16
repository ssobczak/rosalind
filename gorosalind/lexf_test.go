// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"testing"
)

func TestCombination(t *testing.T) {
	p := NewCombination("qwe", 2)

	if p.To_s() != "qq" {
		t.Errorf("got %s, expected qq", p.To_s())
	}

	p.Next()
	if p.To_s() != "qw" {
		t.Errorf("got %s, expected qw", p.To_s())
	}

	for i := 1; i != 8; i++ {
		if !p.Next() {
			t.Errorf("should be able to adnavce %d-th time", i)
		}
	}

	if p.To_s() != "ee" {
		t.Errorf("got %s, expected ee", p.To_s())
	}
	if p.Next() {
		t.Errorf("should NOT be able to adnavce")
	}
}
