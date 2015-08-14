package bees_test

import (
	"sync"

	"github.com/anova/swarm/bees"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// func handleConnection(conn net.Conn) {

// }

// func AcceptLoop(ln *Listener, acceptanceTestFunc func(string) bool) {
// 	for {
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			fmt.Printf("Error in Accept Loop %v\n", err)
// 		}
// 		go handleConnection(conn)
// 	}
// }

var swarm []*MockBee

func NewMockBee() bees.Bee {
	bee := &MockBee{}
	swarm = append(swarm, bee)
	return bee
}

type MockBee struct {
	StingCalled bool
}

func (mockBee *MockBee) Sting(addr string, id int, wg *sync.WaitGroup) {
	mockBee.StingCalled = true

	wg.Done()
}

var _ = Describe("The load tester", func() {

	var hive *bees.Hive
	BeforeEach(func() {
		hive = bees.NewHive()
	})

	It("initializes a hive", func() {

		Expect(hive).ToNot(BeNil())
	})

	It("spawns up a wave of killer bees", func() {
		hive.ReleaseTheBees(NewMockBee, "127.0.0.1", 10)

		allBeesStung := true
		for _, bee := range swarm {
			allBeesStung = allBeesStung && bee.StingCalled
		}

		Expect(allBeesStung).To(BeTrue())
	})

})
