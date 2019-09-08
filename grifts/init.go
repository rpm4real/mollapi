package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/rpm4real/mollapi/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
