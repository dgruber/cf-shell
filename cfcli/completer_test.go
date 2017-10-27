package cfcli

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/c-bata/go-prompt"
)

func find(cmds []prompt.Suggest, cmd string) bool {
	for i := range cmds {
		if cmds[i].Text == cmd {
			return true
		}
	}
	return false
}

var _ = Describe("Completer", func() {

	Context("completions", func() {

		It("should find basic cf commands", func() {
			all := createCompletions()
			Ω(find(all, "push")).Should(BeTrue())
			Ω(find(all, "apps")).Should(BeTrue())
			Ω(find(all, "login")).Should(BeTrue())
		})

		It("should find all plugin extionsions", func() {
			all := createCompletions()
			Ω(find(all, "quit")).Should(BeTrue())
			Ω(find(all, "exit")).Should(BeTrue())
		})

	})

	Context("completions helper functions", func() {

		It("should be able to create a usageMap", func() {
			usageMap := createUsageCompletionsMap()
			Ω(usageMap["push"]).Should(ContainSubstring("[-b BUILDPACK_NAME]"))
		})

		It("should be able to create a usage completion", func() {
			completion := createUsageCompletion("push")
			Ω(len(completion)).Should(BeNumerically("==", 1))
			Ω(completion[0].Description).Should(ContainSubstring("[-b BUILDPACK_NAME]"))
			// "cf push" must be filtered out
			Ω(completion[0].Description).ShouldNot(ContainSubstring("push"))
		})

	})

})
