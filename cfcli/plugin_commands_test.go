package cfcli

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PluginCommands", func() {

	Context("plugin extensions", func() {
		It("should find registered extensions", func() {
			exit := isPluginCommand("exit")
			Ω(exit).Should(BeTrue())
			quit := isPluginCommand("quit")
			Ω(quit).Should(BeTrue())
			ls := isPluginCommand("ls")
			Ω(ls).Should(BeTrue())
			ls = isPluginCommand("ls -lisa")
			Ω(ls).Should(BeTrue())
			dir := isPluginCommand("dir")
			Ω(dir).Should(BeTrue())
			pwd := isPluginCommand("pwd")
			Ω(pwd).Should(BeTrue())
			xyz := isPluginCommand("xyz")
			Ω(xyz).Should(BeFalse())
		})
	})

})
