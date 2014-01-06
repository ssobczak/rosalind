// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package graph

import (
	"fmt"
)

type Node struct {
	id      int
	links   []Node
	payload interface{}
}

type Graph map[int]Node

func FromAdjList(string list) Node {
	lines := strings.Split(list, "\n")
	len := 0
	fmt.Sscanf(lines[0], "%d", len)

	from, to := 0, 0
	for i := 1; i != len(lines); i++ {
		fmt.Sscanf(lines[i], "%d %d", from, to)
	}
}
