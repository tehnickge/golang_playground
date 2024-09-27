package main

import (
	"fmt"
	"runtime"
	"time"
)

func printNumbers(prefix string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%s: %d\n", prefix, i)
		time.Sleep(1 * time.Millisecond) // Имитация длительной работы
	}
}

func main() {
	runtime.GOMAXPROCS(1) // Ограничение использования одним процессорным ядром
	go printNumbers("Goroutine1")
	go printNumbers("Goroutine2")
	time.Sleep(100 * time.Millisecond) // Дать время для завершения Goroutines
}
