// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
)

func pper(n, k int) (ret int) {
	ret = 1
	for i := n; i != n-k; i-- {
		ret = (ret * i) % 1e6
	}
	return
}

func Pper() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	fmt.Println(pper(n, k))
}
