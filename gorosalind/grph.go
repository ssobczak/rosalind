package gorosalind

import (
	"fmt"
)

type Edge struct {
	from *Dna
	to   *Dna
}

type Graph struct {
	v []Dna
	e []Edge
}

func (g *Graph) printEdges() {
	for _, e := range g.e {
		fmt.Println(e.from.name, e.to.name)
	}
}

func (g *Graph) print() {
	fmt.Println("vertices:")
	for _, v := range g.v {
		fmt.Println(v.name, v.sequence)
	}

	fmt.Println("edges:")
	g.printEdges()
}

func overlap(u, v string, k int) bool {
	return u[len(u)-k:len(u)] == v[:k]
}

func grph(dna []Dna, k int) (graph Graph) {
	graph.v = dna
	graph.e = make([]Edge, 0)

	for i, u := range graph.v {
		for j, v := range graph.v {
			if i != j && overlap(u.sequence, v.sequence, k) {
				graph.e = append(graph.e, Edge{&graph.v[i], &graph.v[j]})
			}
		}
	}

	return
}

func Grph() {
	var dna = DnaFromFasta(FileFromStdin())
	graph := grph(dna, 3)
	graph.printEdges()
}
