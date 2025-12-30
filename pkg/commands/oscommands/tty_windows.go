//go:build windows

package oscommands

import (
	"errors"
)

// getTtyName returns an error on Windows as TTY detection is not supported.
// GPG_TTY is not typically needed on Windows as GPG uses different mechanisms.
func getTtyName(fd int) (string, error) {
	return "", errors.New("TTY name detection not supported on Windows")
}
