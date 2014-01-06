// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package permutable

import (
	"testing"
)

func TestPermutable(t *testing.T) {
	perm := []uint8{'A', 'D', 'D', 'D', 'D', 'D'}
	NextPermutation(PermChars(perm))

	if string(perm) != "DADDDD" {
		t.Errorf("nope - got %s instead of DAD", string(perm))
	}

	perm = []uint8{'A', 'D', 'C', 'B'}
	NextPermutation(PermChars(perm))

	if string(perm) != "BACD" {
		t.Errorf("nope - got %s instead of BACD", string(perm))
	}
}
