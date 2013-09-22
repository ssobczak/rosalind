// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

func iev(a [6]int) float32 {
	return 2 * (float32(a[0]+a[1]+a[2]) + 0.75*float32(a[3]) + 0.5*float32(a[4]))
}
