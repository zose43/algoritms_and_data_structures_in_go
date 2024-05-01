package main

import (
	"fmt"
	"time"
)

func main() {
	const Val = 100000
	On(Val)
	On_v2(Val)
	On_v3(Val)
	On2(Val)
	logN(Val)
}

func On_v3(n int) {
	t := time.Now()
	var count, j int
	for i := 0; i < n; i++ {
		for ; j < n; j++ {
			count++
		}
	}
	fmt.Println("O(n)", time.Since(t).Truncate(time.Nanosecond).String(), count)
}

func On_v2(n int) {
	t := time.Now()
	var count int
	for i := 0; i < n; i++ {
		count++
	}
	for i := 0; i < n; i++ {
		count++
	}
	fmt.Println("O(n)", time.Since(t).Truncate(time.Nanosecond).String(), count)
}

func logN(n int) {
	t := time.Now()
	var count int
	i := 1
	for i < n {
		count++
		i *= 3
	}
	fmt.Println("O(log n", time.Since(t).Truncate(time.Nanosecond).String(), count)
}

func On2(n int) {
	t := time.Now()
	var count int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count++
		}
	}
	fmt.Println("O(n^2)", time.Since(t).Truncate(time.Nanosecond).String(), count)
}

func On(n int) {
	t := time.Now()
	var count int
	for i := 0; i < n; i++ {
		count++
	}
	fmt.Println("O(n)", time.Since(t).Truncate(time.Nanosecond).String(), count)
}
