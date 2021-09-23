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

type Announcer struct {
	mu        *sync.Mutex
	formatter announcementFormatter
}

func New() *Announcer {
	return &Announcer{
		formatter: getDefaultFormatter(),
		mu:        &sync.Mutex{},
	}
}

func (ann *Announcer) Tell(msg string, cb ...*AnnouncementArgs) {
	ann.write(info, msg, cb...)
}

func (ann *Announcer) Warn(msg string, cb ...*AnnouncementArgs) {
	ann.write(warn, msg, cb...)
}

func (ann *Announcer) Error(msg string, cb ...*AnnouncementArgs) {
	ann.write(err, msg, cb...)
}

func (ann *Announcer) write(level level, msg string, cb ...*AnnouncementArgs) {
	ann.mu.Lock()

	// populate empty arg to prevent nil reference panic
	if len(cb) == 0 {
		cb = append(cb, Args())
	}

	text := ann.formatter.compose(level, msg, cb[0].contexes)
	fmt.Fprint(outputs[level], text+fmt.Sprintln())

	ann.mu.Unlock()
}

func Redirect(w io.Writer) error {
	if w == nil {
		return errors.New("w io.Writer should not be nil")
	}

	for level := range outputs {
		outputs[level] = w
	}

	return nil
}
