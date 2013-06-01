// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

// match regexp "N[^P][ST][^P]" with overlapping
func match(text string) []int {
	ret := make([]int, 0)

	for i := 0; i < len(text)-3; i++ {
		good := true
		for j := 0; j != 4; j++ {
			switch j {
			case 0:
				if text[i+j] != 'N' {
					good = false
					break
				}
			case 1:
				if text[i+j] == 'P' {
					good = false
					break
				}
			case 2:
				if text[i+j] != 'S' && text[i+j] != 'T' {
					good = false
					break
				}
			case 3:
				if text[i+j] == 'P' {
					good = false
					break
				}
			}
		}
		if good {
			ret = append(ret, i+1)
		}
	}

	return ret
}

func fetch_fasta(uniprot_id string) (body string) {
	resp, err := http.Get("http://www.uniprot.org/uniprot/" + uniprot_id + ".fasta")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)

	scanner.Scan() // discard comment line
	for scanner.Scan() {
		body += scanner.Text()
	}

	return
}

type protein struct {
	id         int
	uniprot_id string
	matches    []int
}

func solve(solved chan protein, count chan int) {
	scanner := bufio.NewScanner(os.Stdin)

	var i int
	for i = 0; scanner.Scan(); i++ {
		pr := protein{
			id:         i,
			uniprot_id: scanner.Text(),
			matches:    []int{},
		}

		go func() {
			seq := fetch_fasta(pr.uniprot_id)
			pr.matches = match(seq)
			solved <- pr
		}()
	}

	count <- i
}

func Mprt() {
	solved := make(chan protein)
	count := make(chan int)

	go solve(solved, count)

	cnt := <-count
	results := make([]protein, cnt)
	for i := 0; i != cnt; i++ {
		res := <-solved
		results[res.id] = res
	}

	for _, prot := range results {
		if len(prot.matches) > 0 {
			fmt.Println(prot.uniprot_id)
			for _, pos := range prot.matches {
				fmt.Printf("%d ", pos)
			}
			fmt.Println()
		}
	}
}
