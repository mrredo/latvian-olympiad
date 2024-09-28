package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numT := strings.Split(scanner.Text(), " ")
	numL := []int{}
	for _, v := range numT {
		n, _ := strconv.Atoi(v)
		numL = append(numL, n)
	}
	fmt.Println(TrisParDivuCenu(numL))
}
func TrisParDivuCenu(l []int) int {
	slices.Sort(l)
	p1:= 0
	p2:= len(l)-1
	n := 0
	if len(l) == 1 {
		return l[0]
	}
	for p1 <= p2  {
		if p1 == p2 {
			n += l[p1]
			break;
		}
		if p2-1==p1 {
			n += l[p1] + l[p2]
			break;
		}	
		n += l[p1] + l[p2]
		p1++
		p2-=2
	}
	return n
}