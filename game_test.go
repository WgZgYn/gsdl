package main

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	data := NewModel(100, 100)
	data.random()
	count := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if data.curr[i][j] {
				//fmt.Print("1 ")
				count++
			} else {
				//fmt.Print("0 ")
			}
		}
		//fmt.Println()
	}
	fmt.Println(count, float64(count)/10000.0)
}

func TestUpdate(t *testing.T) {
	data := NewModel(5, 5)
	data.random()
	fmt.Println("The initial data")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if data.curr[i][j] {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------")
	fmt.Println("update, the lives")
	data.update()
	fmt.Println("The last: ")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if data.last[i][j] {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------")
	fmt.Println("The curr")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if data.curr[i][j] {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}

func BenchmarkRandomForEach(b *testing.B) {
	d := NewModel(100, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.random()
	}
}
