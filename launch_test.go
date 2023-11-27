package launch_test

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/clarktrimble/launch"
)

func TestLaunch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Launch Suite")
}

var _ = Describe("Launch", func() {

	var (
		ctx context.Context
	)
	BeforeEach(func() {
		ctx = context.Background()
	})

	Describe("checking an error", func() {
		var (
			lgr Logger
			err error
		)

		JustBeforeEach(func() {
			Check(ctx, lgr, err)
		})

		When("err is nil", func() {
			BeforeEach(func() {
				lgr = nil
				err = nil
			})
			It("does nothing", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
		// Todo: more interesting when it is not, yeah
	})

	Describe("loading config", func() {
		type Config struct {
			Testa string
		}
		var (
			cfg    Config
			prefix string
			val    string
		)

		JustBeforeEach(func() {
			os.Setenv("TST_TESTA", val)
			Load(&cfg, prefix, "")
		})

		When("env var is set", func() {
			BeforeEach(func() {
				prefix = "tst"
				val = "borbu"
			})
			It("populates cfg struct", func() {
				Expect(cfg).To(Equal(Config{Testa: "borbu"}))
			})
		})
	})

})
