package main

import (
	"fmt"
	"time"
	"sync"
)

func do(i int, wg *sync.WaitGroup) {
	fmt.Printf("Start Job: %d\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("End Job: %d\n", i)

	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	go do(1, &wg)
	go do(2, &wg)
	go do(3, &wg)

	wg.Wait()
	fmt.Println("Jobs all done")
}
