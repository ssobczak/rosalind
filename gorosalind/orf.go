// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"bufio"
	"fmt"
	"os"
)

func cloneChannel(c chan byte) (chan byte, chan byte) {
	r1, r2 := make(chan byte), make(chan byte)

	go func() {
		for tmp := range c {
			r1 <- tmp
			r2 <- tmp
		}

		close(r1)
		close(r2)
	}()

	return r1, r2
}

func printChan(ch chan byte) {
	var buff string

	for letter := range ch {
		buff += string(letter)
	}

	fmt.Println(buff)
}

func Orf() {
	var sequence string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // discard comment line
	for scanner.Scan() {
		sequence += scanner.Text()
	}

	c123 := make(chan byte)
	c1, c23 := cloneChannel(c123)
	c2, c3 := cloneChannel(c23)

	go printChan(c1)
	go printChan(c2)
	go printChan(c3)

	c23 <- sequence[0]
	c3 <- sequence[1]

	for i := 2; i < len(sequence); i++ {
		c123 <- sequence[i]
	}
}
