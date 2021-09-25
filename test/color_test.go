package test

import (
	"testing"

	"github.com/fatih/color"
	"github.com/nicholastcs/go-announcer"
)

func TestTell(t *testing.T) {
	color.NoColor = false

	tests := []test{
		{
			message: "",
			args:    announcer.Args(),
			want:    tellColoredReplacments(t, "", emptyAnnouncement),
		},
		{
			message: "",
			args:    announcer.Args().AddContext("Context", loremIpsumContext),
			want:    tellColoredReplacments(t, "\x1b[1;97mContext\x1b[0m", contextOnly),
		},
		{
			message: "",
			args:    announcer.Args().AddContext("Context", loremIpsumContext, true),
			want:    tellColoredReplacments(t, "\x1b[1;91mContext\x1b[0m", contextOnly),
		},
		{
			message: loremIpsum,
			args:    announcer.Args().AddContext("Context", loremIpsumContext),
			want:    tellColoredReplacments(t, "\x1b[1;97mContext\x1b[0m", normal),
		},
		{
			message: loremIpsum,
			args:    announcer.Args().AddContext("Context", loremIpsumContext, true),
			want:    tellColoredReplacments(t, "\x1b[1;91mContext\x1b[0m", normal),
		},
		{
			message: loremIpsum,
			args:    announcer.Args().AddContext("Context", loremIpsumContext, true),
			want:    tellColoredReplacments(t, "\x1b[1;91mContext\x1b[0m", normal),
		},
		{
			message: loremIpsum,
			want:    tellColoredReplacments(t, "\x1b[1;91mContext\x1b[0m", announcementOnly),
		},
	}

	comparisonTest(t, tests, func(sut *announcer.Announcer, test test) {
		sut.Tell(test.message, test.args)
	})
}

func TestWarn(t *testing.T) {
	color.NoColor = false

	tests := []test{
		{
			message: "",
			args:    announcer.Args(),
			want:    warnColoredReplacments(t, "", emptyAnnouncement),
		},
		{
			message: "",
			args:    announcer.Args().AddContext("Context", loremIpsumContext),
			want:    warnColoredReplacments(t, "\x1b[1;97mContext\x1b[0m", contextOnly),
		},
		{
			message: "",
			args:    announcer.Args().AddContext("Context", loremIpsumContext, true),
			want:    warnColoredReplacments(t, "\x1b[1;91mContext\x1b[0m", contextOnly),
		},
		{
			message: loremIpsum,
			args:    announcer.Args().AddContext("Context", loremIpsumContext),
			want:    warnColoredReplacments(t, "\x1b[1;97mContext\x1b[0m", normal),
		},
		{
			message: loremIpsum,
			args:    announcer.Args().AddContext("Context", loremIpsumContext, true),
			want:    warnColoredReplacments(t, "\x1b[1;91mContext\x1b[0m", normal),
		},
		{
			message: loremIpsum,
			want:    warnColoredReplacments(t, "\x1b[1;91mContext\x1b[0m", announcementOnly),
		},
	}

	comparisonTest(t, tests, func(sut *announcer.Announcer, test test) {
		sut.Warn(test.message, test.args)
	})
}

func TestError(t *testing.T) {
	color.NoColor = false

	tests := []test{
		{
			message: "",
			args:    announcer.Args(),
			want:    errorColoredReplacments(t, "", emptyAnnouncement),
		},
		{
			message: "",
			args:    announcer.Args().AddContext("Context", loremIpsumContext),
			want:    errorColoredReplacments(t, "\x1b[1;97mContext\x1b[0m", contextOnly),
		},
		{
			message: "",
			args:    announcer.Args().AddContext("Context", loremIpsumContext, true),
			want:    errorColoredReplacments(t, "\x1b[1;91mContext\x1b[0m", contextOnly),
		},
		{
			message: loremIpsum,
			args:    announcer.Args().AddContext("Context", loremIpsumContext),
			want:    errorColoredReplacments(t, "\x1b[1;97mContext\x1b[0m", normal),
		},
		{
			message: loremIpsum,
			args:    announcer.Args().AddContext("Context", loremIpsumContext, true),
			want:    errorColoredReplacments(t, "\x1b[1;91mContext\x1b[0m", normal),
		},
		{
			message: loremIpsum,
			want:    errorColoredReplacments(t, "\x1b[1;91mContext\x1b[0m", announcementOnly),
		},
	}

	comparisonTest(t, tests, func(sut *announcer.Announcer, test test) {
		sut.Error(test.message, test.args)
	})
}
