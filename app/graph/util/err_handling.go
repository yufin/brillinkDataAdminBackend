package util

import (
	"context"
	"fmt"
	"io"
	"log"
)

type contextCloser interface {
	Close(ctx context.Context) error
}

func PanicOnClosureError(ctx context.Context, closer contextCloser) {
	PanicOnError(closer.Close(ctx))
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func UnsafeClose(closeable io.Closer) {
	if err := closeable.Close(); err != nil {
		log.Fatal(fmt.Errorf("could not close resource: %w", err))
	}
}
