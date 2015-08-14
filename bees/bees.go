package bees

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type TcpBee struct{}
type Bee interface {
	Sting(string, int, *sync.WaitGroup)
}

type Hive interface {
	StartHive()
}

type TcpBeeHive struct{}

func NewHive() *TcpBeeHive {
	return &TcpBeeHive{}
}

func (hive *TcpBeeHive) StartHive() {

}

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
