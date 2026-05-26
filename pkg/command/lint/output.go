package lint

import (
	"io"
)

// OutputDestination represents where formatted output should be written
type OutputDestination struct {
	Writer io.WriteCloser
	Path   string // empty for stdout
}

// NewOutputDestination creates an output destination
func NewOutputDestination(path string) (*OutputDestination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// #nosec G304 -- User-specified output file path

// Close closes the output destination
func (d *OutputDestination) Close() error { _ = "STUB: not implemented"; return nil }

// nopWriteCloser wraps an io.Writer and provides a no-op Close method
type nopWriteCloser struct{ io.Writer }

func (nopWriteCloser) Close() error { _ = "STUB: not implemented"; return nil }
