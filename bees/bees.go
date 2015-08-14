package bees

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	ws "github.com/gorilla/websocket"
)

type Bee interface {
	Sting(string, int, *sync.WaitGroup)
}

type TcpHoneyBee struct{}
type WebSocketWasp struct{}

type BeeFactoryFunc func() Bee

type Hive struct{}

func NewTcpHoneyBee() Bee {
	return &TcpHoneyBee{}
}

func NewWebSocketWasp() Bee {
	return &WebSocketWasp{}
}

func NewHive() *Hive {
	return &Hive{}
}

func (tcpbee TcpHoneyBee) Sting(addr string, id int, wg *sync.WaitGroup) {
	fmt.Printf("Stinger: %d\n", id)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Bad connection - %v\n", err)
	}

	fmt.Printf("	Dialed: %d\n", id)

	for i := 0; i < 10; i++ {
		hiMessage := fmt.Sprintf("hi from %d: %d\n", id, i)

		io.WriteString(conn, hiMessage)

		time.Sleep(time.Second)
	}
	wg.Done()
}

func (wasp WebSocketWasp) Sting(addr string, id int, wg *sync.WaitGroup) {
	fmt.Printf("Stinger: %d\n", id)

	conn, _, err := ws.DefaultDialer.Dial(addr, http.Header{})
	if err != nil {
		fmt.Println("Bad connection - %v\n", err)
	}

	fmt.Printf("	Dialed: %d\n", id)
	hiMessage := fmt.Sprintf("hi %d", id)
	conn.WriteMessage(ws.TextMessage, []byte(hiMessage))

	time.Sleep(time.Second)

	wg.Done()
}

func (hive *Hive) ReleaseTheBees(queenBee BeeFactoryFunc, dest string, stingerCount int) {
	var wg sync.WaitGroup
	bee := queenBee()
	for i := 0; i < stingerCount; i++ {
		wg.Add(1)
		go bee.Sting(dest, i, &wg)
	}

	fmt.Println("Waiting")
	wg.Wait()
}
