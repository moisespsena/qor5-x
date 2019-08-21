package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSystemBarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSystemBar(children ...h.HTMLComponent) (r *VSystemBarBuilder) {
	r = &VSystemBarBuilder{
		tag: h.Tag("v-system-bar").Children(children...),
	}
	return
}

func (b *VSystemBarBuilder) Absolute(v bool) (r *VSystemBarBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VSystemBarBuilder) App(v bool) (r *VSystemBarBuilder) {
	b.tag.Attr(":app", fmt.Sprint(v))
	return b
}

func (b *VSystemBarBuilder) Color(v string) (r *VSystemBarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSystemBarBuilder) Dark(v bool) (r *VSystemBarBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VSystemBarBuilder) Fixed(v bool) (r *VSystemBarBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VSystemBarBuilder) Height(v int) (r *VSystemBarBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VSystemBarBuilder) Light(v bool) (r *VSystemBarBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VSystemBarBuilder) LightsOut(v bool) (r *VSystemBarBuilder) {
	b.tag.Attr(":lights-out", fmt.Sprint(v))
	return b
}

func (b *VSystemBarBuilder) Status(v bool) (r *VSystemBarBuilder) {
	b.tag.Attr(":status", fmt.Sprint(v))
	return b
}

func (b *VSystemBarBuilder) Window(v bool) (r *VSystemBarBuilder) {
	b.tag.Attr(":window", fmt.Sprint(v))
	return b
}

func (b *VSystemBarBuilder) Class(names ...string) (r *VSystemBarBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSystemBarBuilder) ClassIf(name string, add bool) (r *VSystemBarBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSystemBarBuilder) On(name string, value string) (r *VSystemBarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSystemBarBuilder) Bind(name string, value string) (r *VSystemBarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSystemBarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
