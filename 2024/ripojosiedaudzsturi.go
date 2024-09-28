package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, v int
	fmt.Scanf("%d %d\n", &n, &v)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nR := scanner.Text()
	scanner.Scan()
	vR := scanner.Text()
	nlS := strings.Split(nR, " ")
	vlS := strings.Split(vR, " ")
	nl := []int{0}
	vl := []int{}
	for _, v := range nlS {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("failed input3")
			return
		}
		nl = append(nl, num)
	}
	for _, v := range vlS {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("failed input4")
			return
		}
		vl = append(vl, num)
	}
	calc := Calculate(nl, vl)
	result := strings.Join(calc, "\n")
	fmt.Println(result)
}

func Calculate(nl, vl []int) []string {
	//rezultātu punkti
	p := []string{}
	//kopā punktu skaits
	k := 0
	//lokācija katram punktam uz nogriežņa
	nloc := []int{}
	for _, nog := range nl {

		k += nog

		nloc = append(nloc, k)
	}
vll:
	for _, d := range vl {

		s := d % k
		//works
		for j, _ := range nl {
			if s == nloc[j] {
				p = append(p, fmt.Sprint((j)%len(nl)+1))
				continue vll
			}
		}
		for j, nlv := range nloc {
			p2 := (j + 1) % (len(nloc))
			v := nloc[p2]
			if p2 == 0 {
				v += k
			}
			if nlv < s && nloc[p2] > s {
				p = append(p, fmt.Sprintf("%d-%d", j+1, p2%(len(nloc)-1)+1))
				continue vll
			}
		}

	}
	return p
}
