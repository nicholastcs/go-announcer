package test

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"

	"github.com/nicholastcs/go-announcer"
	"github.com/stretchr/testify/assert"
)

const (
	loremIpsum        = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed placerat magna nec augue convallis, ut aliquam augue malesuada. In quis nunc hendrerit, vestibulum diam eget, congue libero. Sed vitae pellentesque metus. Cras scelerisque libero lectus, ut accumsan mi vulputate vitae. Etiam ullamcorper risus feugiat sodales auctor. Pellentesque ultrices convallis nulla vel ultricies. Morbi at odio non elit varius porta ac ut risus. Proin sit amet dolor elementum, aliquet arcu id, lobortis quam. Suspendisse cursus non lectus nec ullamcorper. Etiam viverra leo vitae vulputate consequat. In bibendum tristique massa. Donec nisi urna, rhoncus a aliquam nec, sollicitudin et mi. Integer finibus.`
	loremIpsumContext = "In quis nunc hendrerit, vestibulum diam eget, congue libero. Sed vitae pellentesque metus. Cras scelerisque libero lectus, ut accumsan mi vulputate vitae. Etiam ullamcorper risus feugiat sodales auctor. Pellentesque ultrices convallis nulla vel ultricies. "
)

var normal, _ = template.New("normal").Parse(fmt.Sprint(
	"{{ .TopBar }} \n",
	"{{ .MiddleBar }}   {{ .Field }}: In quis nunc hendrerit, vestibulum diam eget, congue\n",
	"{{ .MiddleBar }} libero. Sed vitae pellentesque metus. Cras scelerisque libero lectus, ut\n",
	"{{ .MiddleBar }} accumsan mi vulputate vitae. Etiam ullamcorper risus feugiat sodales\n",
	"{{ .MiddleBar }} auctor. Pellentesque ultrices convallis nulla vel ultricies. \n",
	"{{ .MiddleBar }} \n",
	"{{ .MiddleBar }} Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed placerat magna\n",
	"{{ .MiddleBar }} nec augue convallis, ut aliquam augue malesuada. In quis nunc hendrerit,\n",
	"{{ .MiddleBar }} vestibulum diam eget, congue libero. Sed vitae pellentesque metus. Cras\n",
	"{{ .MiddleBar }} scelerisque libero lectus, ut accumsan mi vulputate vitae. Etiam\n",
	"{{ .MiddleBar }} ullamcorper risus feugiat sodales auctor. Pellentesque ultrices convallis\n",
	"{{ .MiddleBar }} nulla vel ultricies. Morbi at odio non elit varius porta ac ut risus. Proin\n",
	"{{ .MiddleBar }} sit amet dolor elementum, aliquet arcu id, lobortis quam. Suspendisse\n",
	"{{ .MiddleBar }} cursus non lectus nec ullamcorper. Etiam viverra leo vitae vulputate\n",
	"{{ .MiddleBar }} consequat. In bibendum tristique massa. Donec nisi urna, rhoncus a aliquam\n",
	"{{ .MiddleBar }} nec, sollicitudin et mi. Integer finibus.\n",
	"{{ .BottomBar }} \n",
))

var contextOnly, _ = template.New("contextOnly").Parse(fmt.Sprint(
	"{{ .TopBar }} \n",
	"{{ .MiddleBar }}   {{ .Field }}: In quis nunc hendrerit, vestibulum diam eget, congue\n",
	"{{ .MiddleBar }} libero. Sed vitae pellentesque metus. Cras scelerisque libero lectus, ut\n",
	"{{ .MiddleBar }} accumsan mi vulputate vitae. Etiam ullamcorper risus feugiat sodales\n",
	"{{ .MiddleBar }} auctor. Pellentesque ultrices convallis nulla vel ultricies. \n",
	"{{ .BottomBar }} \n",
))

var announcementOnly, _ = template.New("announcementOnly").Parse(fmt.Sprint(
	"{{ .TopBar }} \n",
	"{{ .MiddleBar }} Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed placerat magna\n",
	"{{ .MiddleBar }} nec augue convallis, ut aliquam augue malesuada. In quis nunc hendrerit,\n",
	"{{ .MiddleBar }} vestibulum diam eget, congue libero. Sed vitae pellentesque metus. Cras\n",
	"{{ .MiddleBar }} scelerisque libero lectus, ut accumsan mi vulputate vitae. Etiam\n",
	"{{ .MiddleBar }} ullamcorper risus feugiat sodales auctor. Pellentesque ultrices convallis\n",
	"{{ .MiddleBar }} nulla vel ultricies. Morbi at odio non elit varius porta ac ut risus. Proin\n",
	"{{ .MiddleBar }} sit amet dolor elementum, aliquet arcu id, lobortis quam. Suspendisse\n",
	"{{ .MiddleBar }} cursus non lectus nec ullamcorper. Etiam viverra leo vitae vulputate\n",
	"{{ .MiddleBar }} consequat. In bibendum tristique massa. Donec nisi urna, rhoncus a aliquam\n",
	"{{ .MiddleBar }} nec, sollicitudin et mi. Integer finibus.\n",
	"{{ .BottomBar }} \n",
))

var emptyAnnouncement, _ = template.New("emptyAnnouncement").Parse(fmt.Sprint(
	"{{ .TopBar }} \n",
	"{{ .MiddleBar }} empty announcement\n",
	"{{ .BottomBar }} \n",
))

type coloredReplacements struct {
	Field     string
	TopBar    string
	MiddleBar string
	BottomBar string
	template  *template.Template
}

func tellColoredReplacments(t *testing.T, field string, template *template.Template) *coloredReplacements {
	if template == nil {
		t.Fatalf("template is empty")
	}

	return &coloredReplacements{
		Field:     field,
		TopBar:    "\x1b[97m╷\x1b[0m",
		MiddleBar: "\x1b[97m│\x1b[0m",
		BottomBar: "\x1b[97m╵\x1b[0m",
		template:  template,
	}
}

func warnColoredReplacments(t *testing.T, field string, template *template.Template) *coloredReplacements {
	if template == nil {
		t.Fatalf("template is empty")
	}

	return &coloredReplacements{
		Field:     field,
		TopBar:    "\x1b[93m╓\x1b[0m",
		MiddleBar: "\x1b[93m║\x1b[0m",
		BottomBar: "\x1b[93m╙\x1b[0m",
		template:  template,
	}
}

func errorColoredReplacments(t *testing.T, field string, template *template.Template) *coloredReplacements {
	if template == nil {
		t.Fatalf("template is empty")
	}

	return &coloredReplacements{
		Field:     field,
		TopBar:    "\x1b[31m╓\x1b[0m",
		MiddleBar: "\x1b[31m║\x1b[0m",
		BottomBar: "\x1b[31m╙\x1b[0m",
		template:  template,
	}
}

type test struct {
	message string
	args    *announcer.AnnouncementArgs
	want    *coloredReplacements
}

func comparisonTest(t *testing.T, tests []test, action func(sut *announcer.Announcer, test test)) {
	gotBuffer := new(bytes.Buffer)
	wantBuffer := new(bytes.Buffer)
	announcer.Redirect(gotBuffer) //nolint

	sut := announcer.New()

	for _, test := range tests {
		// construct announcement got from SUT
		// sut.Tell(test.message, test.args)
		action(sut, test)

		got := gotBuffer.String()

		// construct announcement want from tmpl
		err := test.want.template.Execute(wantBuffer, test.want)
		if err != nil {
			t.Fatalf("malformed template %s", err)
		}
		want := wantBuffer.String()

		okay := assert.Equal(t, want, got)

		if okay {
			fmt.Print(want)
		}

		gotBuffer.Reset()
		wantBuffer.Reset()
	}
}
