package main

import (
	"fmt"
	"math/rand"
)

var (
	num     int
	bottles []int
	target  []int
	ops     = 0
)

func swap(slice *[]int) func(i, j int) {
	return func(i, j int) {
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
}

func initGame() {
	fmt.Println("输入瓶子的数量：")
	fmt.Scanf("%d", &num)
	for i := 0; i < num; i++ {
		bottles = append(bottles, i)
		target = append(target, i)
	}
	for {
		rand.Shuffle(num, swap(&bottles))
		rand.Shuffle(num, swap(&target))
		if check(false) != num || num == 0 {
			break
		}
	}
}

func iter(op func(...int), targets ...(*[]int)) {
	for i := 0; i < num; i++ {
		values := []int{}
		for _, target := range targets {
			values = append(values, (*target)[i])
		}
		op(values...)
	}
}

func check(print bool) int {
	equals := 0
	iter(func(i ...int) {
		if len(i) != 2 {
			panic("error")
		}
		if i[0] == i[1] {
			equals++
		}
	}, &bottles, &target)
	if print {
		fmt.Println("当前正确的瓶子数量是：", equals)
	}
	return equals
}

func printBottles() {
	fmt.Println("当前的瓶子顺序是：")
	fmt.Print(" ")
	for i := 0; i < num; i++ {
		fmt.Printf("%d   ", i)
	}
	fmt.Println()
	iter(func(i ...int) {
		fmt.Printf("%d ", i)
	}, &bottles)
	fmt.Println()
}

func operate() {
	var i, j int
	fmt.Println("输入2个整数，交换2个位置瓶子，以空格分隔：")
	fmt.Scanf("%d %d", &i, &j)
	bottles[i], bottles[j] = bottles[j], bottles[i]
	ops++
}

func main() {
	fmt.Println("\033[H\033[2J")
	initGame()
	for {
		fmt.Println("\033[H\033[2J")
		if check(true) == num {
			fmt.Println("游戏结束！你通过%d次操作完成了%d个瓶子的归位", ops, num)
			break
		}
		printBottles()
		operate()
	}
}
