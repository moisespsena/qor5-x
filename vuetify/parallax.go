package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VParallaxBuilder struct {
	tag *h.HTMLTagBuilder
}

func VParallax(children ...h.HTMLComponent) (r *VParallaxBuilder) {
	r = &VParallaxBuilder{
		tag: h.Tag("v-parallax").Children(children...),
	}
	return
}

func (b *VParallaxBuilder) Alt(v string) (r *VParallaxBuilder) {
	b.tag.Attr("alt", v)
	return b
}

func (b *VParallaxBuilder) Height(v string) (r *VParallaxBuilder) {
	b.tag.Attr("height", v)
	return b
}

func (b *VParallaxBuilder) Src(v string) (r *VParallaxBuilder) {
	b.tag.Attr("src", v)
	return b
}

func (b *VParallaxBuilder) Class(names ...string) (r *VParallaxBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VParallaxBuilder) ClassIf(name string, add bool) (r *VParallaxBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VParallaxBuilder) On(name string, value string) (r *VParallaxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VParallaxBuilder) Bind(name string, value string) (r *VParallaxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VParallaxBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
