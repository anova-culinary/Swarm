package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func Stinger(addr string, id int, wg *sync.WaitGroup) {
	fmt.Printf("Stinger: %d\n", id)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Bad connection - %v\n", err)
	}

	fmt.Printf("	Dialed: %d\n", id)
	hiMessage := fmt.Sprintf("hi %d", id)
	io.WriteString(conn, hiMessage)

	time.Sleep(time.Second)

	wg.Done()

}

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
