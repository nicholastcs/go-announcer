package announcer

import "strings"

type Context struct {
	value     string
	emphasize bool
}

type AnnouncementArgs struct {
	contexes map[string]Context
}

func Args() *AnnouncementArgs {
	return &AnnouncementArgs{
		contexes: map[string]Context{},
	}
}

func (cb *AnnouncementArgs) AddContext(field string, value string, emphasize ...bool) *AnnouncementArgs {
	if len(emphasize) == 0 {
		emphasize = append(emphasize, false)
	}

	if strings.TrimSpace(field) == "" || strings.TrimSpace(value) == "" {
		return cb
	}

	cb.contexes[field] = Context{
		value:     value,
		emphasize: emphasize[0],
	}

	return cb
}
