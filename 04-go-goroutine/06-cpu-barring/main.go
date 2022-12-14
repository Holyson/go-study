package main

import "time"

func main() {
	running := true
	go func() {
		println("start thread1")
		count := 1
		for running {
			count++
		}
		println("end thread1: count =", count)
	}()
	go func() {
		println("start thread2")
		for {
			running = false
		}
	}()
	time.Sleep(time.Hour)
}
