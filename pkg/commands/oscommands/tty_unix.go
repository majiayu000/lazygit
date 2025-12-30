//go:build !windows

package oscommands

import (
	"fmt"
	"os"
)

// getTtyName returns the TTY device name for the given file descriptor.
// This is needed for setting GPG_TTY environment variable.
func getTtyName(fd int) (string, error) {
	// Try to readlink /proc/self/fd/{fd} which works on Linux
	procPath := fmt.Sprintf("/proc/self/fd/%d", fd)
	ttyPath, err := os.Readlink(procPath)
	if err == nil {
		return ttyPath, nil
	}

	// Fallback for macOS and BSD: use /dev/fd/{fd}
	devPath := fmt.Sprintf("/dev/fd/%d", fd)
	ttyPath, err = os.Readlink(devPath)
	if err == nil {
		return ttyPath, nil
	}

	// If both fail, we can't determine the TTY
	return "", err
}
