package modules_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestModules(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Modules Suite")
}
