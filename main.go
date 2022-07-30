package main

import (
	"fmt"
	"sync"
)

func gr(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	ss := []string{
		"A",
		"B",
		"C",
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(ss))

	for i, s := range ss {
		go gr(fmt.Sprintf("%d-%s", i, s), wg)
	}

	wg.Wait()

	println("Finished ...")
}
