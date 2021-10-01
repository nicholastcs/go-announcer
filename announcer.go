package announcer

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

const (
	info level = "info"
	warn level = "warn"
	err  level = "error"
)

var (
	outputs = map[level]io.Writer{
		info: os.Stdout,
		warn: os.Stdout,
		err:  os.Stderr,
	}
)

type level string

// Announcer is the struct that prints announcement to console or terminal.
type Announcer struct {
	mu        *sync.Mutex
	formatter announcementFormatter
}

// New instantiates Announcer struct with sensible defaults.
// Currently the API does not expose any means to customize the
// Formatting or symbols.
func New() *Announcer {
	return &Announcer{
		formatter: getDefaultFormatter(),
		mu:        &sync.Mutex{},
	}
}

// Tell tells an announcement.
func (ann *Announcer) Tell(msg string, cb ...*AnnouncementArgs) {
	ann.write(info, msg, cb...)
}

// Warn tells a warning. Emphasis bar is yellow color.
func (ann *Announcer) Warn(msg string, cb ...*AnnouncementArgs) {
	ann.write(warn, msg, cb...)
}

// Error tells a error. Emphasis bar is red.
// Emits in stderr.
func (ann *Announcer) Error(msg string, cb ...*AnnouncementArgs) {
	ann.write(err, msg, cb...)
}

func (ann *Announcer) write(level level, msg string, cb ...*AnnouncementArgs) {
	ann.mu.Lock()

	// populate empty arg to prevent nil reference panic
	if len(cb) == 0 {
		cb = append(cb, Args())
	}
	if cb[0] == nil {
		cb[0] = Args()
	}

	text := ann.formatter.compose(level, msg, cb[0].contexes)
	fmt.Fprint(outputs[level], text+fmt.Sprintln())

	ann.mu.Unlock()
}

// Redirect stream to a io.Writer instead of stdout or stderr.
func Redirect(w io.Writer) error {
	if w == nil {
		return errors.New("w io.Writer should not be nil")
	}

	for level := range outputs {
		outputs[level] = w
	}

	return nil
}
