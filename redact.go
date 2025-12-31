package launch

import (
	"os"
	"strings"
)

// Redact is a type for private strings.
// When loaded via envconfig, if the value starts with "/" it is treated
// as a file path and the contents are read from the file.
type Redact string

// Decode implements the envconfig.Decoder interface.
func (redact *Redact) Decode(value string) error {

	if strings.HasPrefix(value, "/") {
		data, err := os.ReadFile(value)
		if err != nil {
			return err
		}
		*redact = Redact(strings.TrimSpace(string(data)))
		return nil
	}

	*redact = Redact(value)
	return nil
}

// MarshalJSON implements the Marshaler interface
func (redact Redact) MarshalJSON() ([]byte, error) {

	if redact == "" {
		return []byte(`"--unset--"`), nil
	}

	return []byte(`"--redacted--"`), nil
}
