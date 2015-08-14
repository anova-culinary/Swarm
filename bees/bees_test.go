package bees_test

import (
	"fmt"
	"net"

	bees "github.com/anova/tcp-bees/bees"

	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func handleConnection(conn net.Conn) {

}

func AcceptLoop() {
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

var _ = Describe("The load tester", func() {
	var conn net.Conn
	BeforeSuite(func() {
		conn, err = net.Listen("tcp", ":3000")
		if err != nil {
			fmt.Printf("Error in BeforeSuite: %v\n", err)
		}

		go AcceptLoop(func(result string) {
			response = result
		})

	})

	It("initializes a hive", func() {
		hive := bees.NewHive()

		Expect(hive).ToNot(BeNil())
	})

	It("connects to the ip address defined in args", func() {
		os.Args = []string{"", "1", "127.0.0.1:3000"}

		hive.ReleaseTheBees()
		Expect(result).To(Equal("Hi 1"))
	})
})
