package cfcli

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCfcli(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cfcli Suite")
}
