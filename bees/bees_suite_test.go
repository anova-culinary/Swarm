package bees_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBees(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bees Suite")
}
