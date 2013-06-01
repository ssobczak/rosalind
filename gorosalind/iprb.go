// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"fmt"
	"log"
)

func iprb() {
	var k, m, n float64

	if _, err := fmt.Scanf("%f %f %f", &k, &m, &n); err != nil {
		log.Fatal(err)
	}

	all := k + m + n

	x := k/all +
		m/(2*all)*(1+k/(all-1)+(m-1)/(2*(all-1))) +
		n/all*(k/(all-1)+m/(2*(all-1)))
	fmt.Printf("%f\n", x)
}
