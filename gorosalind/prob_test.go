// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"math"
	"testing"
)

type test struct {
	*testing.T
}

func (t *test) expect_close(got, expected float64, test_name string) {
	if math.Abs(got-expected) > 1e-3 {
		t.Errorf("Got %f instead of %f in %s.", got, expected, test_name)
	}
}

func TestProb(testing *testing.T) {
	t := test{testing}
	t.expect_close(prob1("ACGATACAA", 0.129), -5.737, "example1")
}
