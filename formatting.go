package announcer

import (
	"fmt"
	"strings"

	"github.com/mitchellh/go-wordwrap"
)

type announcementFormatter interface {
	compose(level level, msg string, contexts map[string]Context) string
}

type defaultAnnouncementFormatter struct {
	maxWordPerLine uint
}

func getDefaultFormatter() *defaultAnnouncementFormatter {
	return &defaultAnnouncementFormatter{
		maxWordPerLine: 75,
	}
}

func (f *defaultAnnouncementFormatter) compose(level level, msg string, contexts map[string]Context) string {
	return f.compile(level, msg, contexts)
}

func (f *defaultAnnouncementFormatter) compile(level level, msg string, ctxs map[string]Context) string {
	compiledContexts := f.compileContexts(ctxs)

	resultant := msg

	if compiledContexts != "" {
		resultant = compiledContexts + fmt.Sprintln() + resultant
	}

	wrapped := strings.TrimLeft(
		wordwrap.WrapString(
			resultant,
			f.maxWordPerLine,
		), fmt.Sprintln(),
	)

	wrapped = fmt.Sprintln() + wrapped + fmt.Sprintln()

	return prependSymbol(level, wrapped)
}

func (f *defaultAnnouncementFormatter) compileContexts(ctxs map[string]Context) string {
	if len(ctxs) == 0 {
		return ""
	}
	var sb strings.Builder

	for field, ctx := range ctxs {
		colorizedField := field
		if ctx.emphasize {
			colorizedField = emphasize(colorizedField)
		}
		sb.WriteString(fmt.Sprintf("%s%s: %s\n", tab, colorizedField, strings.TrimRight(ctx.value, fmt.Sprintln())))
	}

	return sb.String()
}

func prependSymbol(level level, wrapped string) string {
	lines := strings.Split(wrapped, fmt.Sprintln())

	symbols := bars(level, len(lines))

	for i := range lines {
		lines[i] = symbols[i] + " " + lines[i]
	}

	return strings.Join(lines, fmt.Sprintln())
}

func bars(level level, count int) []string {
	topEnd := defaultSymbolConfig[level].topEnd
	middle := defaultSymbolConfig[level].middle
	bottomEnd := defaultSymbolConfig[level].bottomEnd
	colorFunc := getColorFunc(level)

	bars := make([]string, count)

	for i := range bars {
		bars[i] = colorFunc(middle)
	}

	bars[0] = colorFunc(topEnd)
	bars[count-1] = colorFunc(bottomEnd)

	return bars
}

func getColorFunc(level level) func(a ...interface{}) string {
	if level == info {
		return hiWhite
	}
	if level == warn {
		return hiYellow
	}
	if level == err {
		return hiRed
	}

	return hiWhite
}
