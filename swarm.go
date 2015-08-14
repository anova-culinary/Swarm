package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/anova/swarm/bees"
)

func printUsage() {
	fmt.Printf("Usage: ./swarm <IP>:<port> <count> <tcp/ws>")
}

func main() {
	hive := bees.NewHive()

	if len(os.Args) < 4 {
		printUsage()
		return
	}

	beeCount, err := strconv.Atoi(os.Args[2])

	protocol := os.Args[3]

	var beeConstructor bees.BeeFactoryFunc

	switch protocol {
	case "tcp":
		beeConstructor = bees.NewTcpHoneyBee
	case "ws":
		beeConstructor = bees.NewWebSocketWasp
	default:
		printUsage()
		return
	}

	if err != nil {
		printUsage()
		return
	}

	hive.ReleaseTheBees(beeConstructor, os.Args[1], beeCount)
}
