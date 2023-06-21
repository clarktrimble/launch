// Package minlog implements a (sub)minimal logger
// see https://github.com/clarktrimble/sabot for a featureful implementation
package minlog

import (
	"context"
	"fmt"
)

type MinLog struct{}

func (ml *MinLog) Info(ctx context.Context, msg string, kv ...any) {

	fmt.Printf("msg > %s\n", msg)
}

func (ml *MinLog) Error(ctx context.Context, msg string, err error, kv ...any) {

	fmt.Printf("err > %s %+v\n", msg, err)
}

func (ml *MinLog) WithFields(ctx context.Context, kv ...any) context.Context {

	return ctx
}
