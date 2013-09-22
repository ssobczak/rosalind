// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"testing"
)

func TestIev(t *testing.T) {
	res := iev([...]int{1, 0, 0, 1, 0, 1})

	if res != 3.5 {
		t.Errorf("nope - got %d instead of %d", res, 3.5)
	}

	res = iev([...]int{19073, 16008, 17191, 18098, 16403, 16395})
	if res != 148094 {
		t.Errorf("nope - got %d instead of %d", res, 148094)
	}
}
