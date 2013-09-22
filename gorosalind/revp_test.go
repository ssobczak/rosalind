// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"testing"
)

func TestRevp(t *testing.T) {
	res := revp("TCAATGCATGCGGGTCTATATGCAT")

	// for _, r := range res {
	// 	fmt.Printf("%d %d\n", r.start, r.start+r.length)
	// }

	if len(res) != 8 {
		t.Errorf("nope - %d", len(res))
	}
}
