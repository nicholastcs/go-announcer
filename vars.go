package announcer

import "github.com/fatih/color"

var (
	tab = "  "

	defaultSymbolConfig = map[level]symbolConfig{
		info: {
			topEnd:    "╷",
			middle:    "│",
			bottomEnd: "╵",
		},
		warn: {
			topEnd:    "╓",
			middle:    "║",
			bottomEnd: "╙",
		},
		err: {
			topEnd:    "╓",
			middle:    "║",
			bottomEnd: "╙",
		},
	}

	emphasizeField = color.New(color.Bold, color.FgHiRed).SprintFunc()
	defaultField   = color.New(color.Bold, color.FgHiWhite).SprintFunc()
	hiYellow       = color.New(color.FgHiYellow).SprintFunc()
	hiWhite        = color.New(color.FgHiWhite).SprintFunc()
	hiRed          = color.New(color.FgRed).SprintFunc()
)

type symbolConfig struct {
	topEnd, middle, bottomEnd string
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
