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

	emphasize = color.New(color.Bold, color.FgHiRed).SprintFunc()
	hiYellow  = color.New(color.FgHiYellow).SprintFunc()
	hiWhite   = color.New(color.FgHiWhite).SprintFunc()
	hiRed     = color.New(color.FgRed).SprintFunc()
)

type symbolConfig struct {
	topEnd, middle, bottomEnd string
}

// func defaultSymbolConfig()
