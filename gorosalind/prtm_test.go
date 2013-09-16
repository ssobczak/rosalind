// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"math"
	"testing"
)

func TestPrtm(t *testing.T) {
	if math.Abs(Prtm("SKADYEK")-821.39192) > 1e-5 {
		t.Errorf("got %f", Prtm("SKADYEK"))
	}

	t.Error(Prtm("SKADYEK"))
}
