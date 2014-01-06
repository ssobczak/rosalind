// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Eval() {
	n, s := 0, ""
	if _, err := fmt.Scanf("%d\n%s\n", &n, &s); err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	if gcs, err := reader.ReadString('\n'); err != nil {
		log.Fatal(err)
	} else {
		gcs = strings.TrimSpace(gcs)
		for _, gc := range strings.Split(gcs, " ") {
			if f, err := strconv.ParseFloat(gc, 64); err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("%f ", eval(n, f, s))
			}
		}
	}
}

func eval(n int, gc_c float64, motif string) float64 {
	gc := gc_count(motif)
	at := len(motif) - gc

	pg := math.Pow(gc_c/2, float64(gc)) * math.Pow((1-gc_c)/2, float64(at))

	return float64(n-len(motif)+1) * pg
}
