// Package sendkeys simulates user input in a terminal.
package sendkeys

import (
	"os"
	"time"

	"golang.org/x/sys/unix"
)

// Target represents the tty to which the keys are sent. Create one by calling
// Open and send keys using Target.SendRunes or Target.SendBytes. The Target
// must be closed after use.
type Target struct {
	f *os.File
	d time.Duration
}

// Open returns a Target that sends keys to the specified TTY. The delay is the
// time to wait before sending each separate string when SendRunes or SendBytes
// is called.
func Open(ttyPath string, delay time.Duration) (*Target, error) {
	f, err := os.Open(ttyPath)
	if err != nil {
		return nil, err
	}
	return &Target{f: f, d: delay}, nil
}

// Close releases the resources used by the Target.
func (t *Target) Close() error {
	return t.f.Close()
}

// SendRunes sends each rune of each string the the Target. All runes of a
// given string are sent without delay, and the target's delay is applied
// between each string.
func (t *Target) SendRunes(strs ...string) (int, error) {
	var n int
	for _, str := range strs {
		for _, r := range str {
			if err := unix.IoctlSetPointerInt(int(t.f.Fd()), unix.TIOCSTI, int(r)); err != nil {
				return n, err
			}
			n++
		}
		time.Sleep(t.d)
	}
	return n, nil
}

// SendBytes is like SendRunes except it sends each byte separately, instead if
// each rune.
func (t *Target) SendBytes(strs ...string) (int, error) {
	var n int
	for _, str := range strs {
		for i := 0; i < len(str); i++ {
			v := int(str[i])
			if err := unix.IoctlSetPointerInt(int(t.f.Fd()), unix.TIOCSTI, v); err != nil {
				return n, err
			}
			n++
		}
		time.Sleep(t.d)
	}
	return n, nil
}
