package launch

import "context"

// Logger is a logging interface
type Logger interface {
	Error(ctx context.Context, msg string, err error, kv ...interface{})
}
