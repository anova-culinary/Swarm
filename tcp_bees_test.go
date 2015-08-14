package main_test

import (
	bees "github.com/anova/tcp-bees/bees"

	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The load tester", func() {
	It("initializes a hive", func() {
		hive = bees.NewHive()

		Expect(hive).ToNot(BeNil())
	})

	It("connects to the ip address defined in args", func() {
		os.Args = []string{"", "1", "127.0.0.1:3000"}
	})
})
