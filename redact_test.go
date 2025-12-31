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

	Describe("decoding a redact string", func() {

		JustBeforeEach(func() {
			err = redact.Decode(string(data))
		})

		When("value is a direct secret", func() {
			BeforeEach(func() {
				data = []byte("secret_value")
			})

			It("stores the value directly", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(string(redact)).To(Equal("secret_value"))
			})
		})

		When("value is a file path", func() {
			BeforeEach(func() {
				data = []byte("/etc/hosts")
			})

			It("reads from the file", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(string(redact)).ToNot(BeEmpty())
			})
		})

		When("value is a non-existent file path", func() {
			BeforeEach(func() {
				data = []byte("/no/such/file")
			})

			It("returns an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

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
