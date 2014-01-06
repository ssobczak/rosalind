// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package permutable

type Permutable interface {
	Len() int
	Less(i, j int) bool
	Eq(i, j int) bool
	Swap(i, j int)
}

func revertEnd(p Permutable, from int) {
	for i := 0; i != (p.Len()-from)/2; i++ {
		p.Swap(from+i, p.Len()-i-1)
	}
}

func NextPermutation(p Permutable) bool {
	for k := p.Len() - 2; k >= 0; k-- {
		if p.Less(k, k+1) {
			for l := p.Len() - 1; l > k; l-- {
				if p.Less(k, l) && !p.Eq(k, k+1) {
					p.Swap(k, l)
					revertEnd(p, k+1)
					return true
				}
			}
		}
	}

	return false
}

type PermInts []int

func (p PermInts) Len() int           { return len(p) }
func (p PermInts) Less(i, j int) bool { return p[i] < p[j] }
func (p PermInts) Eq(i, j int) bool   { return p[i] == p[j] }
func (p PermInts) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type PermChars []uint8

func (p PermChars) Len() int           { return len(p) }
func (p PermChars) Less(i, j int) bool { return p[i] < p[j] }
func (p PermChars) Eq(i, j int) bool   { return p[i] == p[j] }
func (p PermChars) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type PermBools []bool

func (p PermBools) Len() int           { return len(p) }
func (p PermBools) Less(i, j int) bool { return p[i] == false && p[j] == true }
func (p PermBools) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
