package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	stingerCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Missing number of stingers:", os.Args)
	}
	for i := 0; i < stingerCount; i++ {
		dest := os.Args[2]
		wg.Add(1)
		go Stinger(dest, i, &wg)
	}

	fmt.Println("Waiting")
	wg.Wait()
}
