//Note: need to use the "go get" command to fetch testing suite

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)

	array := strings.Fields(s)

	for _, v := range array {

		count := 0

		for _, k := range array {

			if v == k {
				count++
			}

		}

		m[v] = count

	}

	return m

}

func main() {
	wc.Test(WordCount)
}
