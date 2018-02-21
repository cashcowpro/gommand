// +build spec

package gommand

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGommand(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gommand Suite")
}
