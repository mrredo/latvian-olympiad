package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Reading input
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // To ignore the first line input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Parsing the input into a list of integers
	numT := strings.Split(scanner.Text(), " ")
	numL := []int{}
	for _, v := range numT {
		n, _ := strconv.Atoi(v)
		numL = append(numL, n)
	}

	// Calling the Dazadie function
	g, h, hi := Dazadie(numL)

	// Printing the results
	fmt.Printf("%d\n%d\n%s\n", g, h, strings.Join(hi, " "))
}

// Dazadie detects the longest sublist with unique elements
// and returns the length of the longest sublist, the count of such sublists, and their starting indices.
func Dazadie(l []int) (g, h int, hi []string) {
	if len(l) == 0 {
		return 0, 0, []string{}
	}

	// Map to store the last seen index of elements
	m := map[int]int{}

	p1 := 0
	p2 := 1

	for p2 <= len(l) {
		v := l[p2-1]
		if _, ok := m[v]; ok && m[v] >= p1 {
			p1 = m[v] + 1
		}
		nl := l[p1:p2]

		if len(nl) > g {
			g = len(nl)
			h = 1
			hi = []string{fmt.Sprint(p1 + 1)} // Store 1-based index
		} else if len(nl) == g {
			h++
			hi = append(hi, fmt.Sprint(p1 + 1)) // Store 1-based index
		}

		// Move p2 forward
		p2++
	}

	return
}
