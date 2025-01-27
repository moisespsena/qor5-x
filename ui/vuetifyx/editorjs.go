package vuetifyx

import (
	"github.com/qor5/x/v3/ui/vuetify"
	h "github.com/theplant/htmlgo"
)

type EditorJsBuilder struct {
	vuetify.VTagBuilder[*EditorJsBuilder]
}

func EditorJS() *EditorJsBuilder {
	return vuetify.VTag(&EditorJsBuilder{}, "vx-editorjs")
}

func (b *EditorJsBuilder) Label(text string) *EditorJsBuilder {
	return b.Attr("label", text)
}

func (b *EditorJsBuilder) ErrorMessages(msgs []string) *EditorJsBuilder {
	if len(msgs) > 0 {
		return b.Attr(":error-messages", h.JSONString(msgs))
	}
	return b
}
