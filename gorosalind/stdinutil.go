// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package gorosalind

import (
	"bufio"
	"os"
)

func FileFromStdin() (ret string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ret += scanner.Text() + "\n"
	}

	return
}
