package main

import (
	"fmt"
	"math"
)

func maxLayers(blocks int) int {
	return int(math.Sqrt(float64(blocks)))
}
func ballsUsed(layers int) int {
	return (layers * (1 + (2*layers - 1))) / 2
}
func main() {

	blocktests := []int{
		2,
	}
	for _, v := range blocktests {
		l := maxLayers(v)
		tb := ballsUsed(l)
		bl := v - tb
		fmt.Println(LastPosOfBall(l, bl).String())
	}

}

type Pos struct {
	y int
	x int
}

func (p Pos) String() string {
	return fmt.Sprintf("%d %d", p.x, p.y)
}

func LastPosOfBall(layers int, ballsleft int) Pos {
	if ballsleft == 0 {
		return Pos{
			x: 0,
			y: layers,
		}
	}

	p := Pos{
		x: layers - ballsleft + 1,
	}
	p.y = layers - p.x + 1

	if ballsleft > layers {
		p.x = -2*layers + ballsleft - 1
		p.y = layers - (p.x * -1) + 1
	}

	return p
}
