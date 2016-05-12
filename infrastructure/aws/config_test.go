package aws_test

import (
	"io"
	"os"

	"github.com/datianshi/pcfdeploy/infrastructure/aws"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var _ = Describe("Given a Valid config yml", func() {
		var cfg io.Reader
		BeforeEach(func() {
			cfg, _ = os.Open("./fixtures/aws.yml")
		})
		Context("when parse the config yml", func() {
			It("Should not have an error happened", func() {
				_, err := aws.Parse(cfg)
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
