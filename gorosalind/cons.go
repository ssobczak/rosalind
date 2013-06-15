package gorosalind

import (
	"bufio"
	"fmt"
	"os"
)

func read_sequences() (seq []string) {
	seq = make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '>' {
			seq = append(seq, "")
		} else {
			seq[len(seq)-1] += line
		}
	}

	return
}

var atoi = map[byte]int{
	'A': 0,
	'C': 1,
	'G': 2,
	'T': 3,
}

var itoa = []int{
	'A',
	'C',
	'G',
	'T',
}

func index_of_max(a []int) (ret int) {
	ret = 0
	for k, v := range a {
		if a[ret] < v {
			ret = k
		}
	}

	return
}

func Cons() {
	seqences := read_sequences()

	profile := make([][]int, len(seqences[0]))
	for k := range profile {
		profile[k] = make([]int, 4)
		for i := range profile[k] {
			profile[k][i] = 0
		}
	}

	for _, seq := range seqences {
		for i := range seq {
			profile[i][atoi[seq[i]]]++
		}
	}

	for _, pr := range profile {
		fmt.Printf("%c", itoa[index_of_max(pr)])
	}
	fmt.Printf("\n")

	for j := range profile[0] {
		fmt.Printf("%c: ", itoa[j])
		for i := range profile {
			fmt.Printf("%d ", profile[i][j])
		}
		fmt.Printf("\n")
	}
}
