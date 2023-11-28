package spinner_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/clarktrimble/launch/spinner"
)

func TestSpinner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spinner Suite")
}

var _ = Describe("Spinner", func() {

	var (
		sp *Spinner
	)

	Describe("creating a new spinner", func() {

		JustBeforeEach(func() {
			sp = New()
		})

		When("all is well", func() {
			It("populates Chars, zero for Start and Count", func() {
				Expect(sp).To(Equal(&Spinner{
					Chars: []string{`-`, `\`, `|`, `/`},
				}))
			})
		})
	})

	Describe("spinning spinner twice", func() {

		JustBeforeEach(func() {
			sp.Spin()
			time.Sleep(time.Millisecond * 9)
			sp.Spin()
		})

		When("all is well", func() {
			BeforeEach(func() {
				sp = New()
			})

			It("has a count of 2 and elapsed of about 10 ms", func() {
				Expect(sp.Count).To(Equal(2))
				Expect(sp.Elapsed()).To(BeNumerically("~", 0.01, 0.001))
			})
		})
	})

})
