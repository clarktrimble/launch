package launch_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/clarktrimble/launch"
)

var _ = Describe("Redact", func() {

	var (
		data   []byte
		err    error
		redact Redact
	)

	Describe("marshalling a redact string", func() {

		JustBeforeEach(func() {
			data, err = json.Marshal(redact)
		})

		When("value is set", func() {
			BeforeEach(func() {
				redact = Redact("password_or_token")
			})

			It("shows redacted", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(data).To(Equal([]byte(`"--redacted--"`)))
			})
		})

		When("value is not set", func() {
			BeforeEach(func() {
				redact = Redact("")
			})

			It("indicates the value is not set", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(data).To(Equal([]byte(`"--unset--"`)))
			})
		})

	})
})
