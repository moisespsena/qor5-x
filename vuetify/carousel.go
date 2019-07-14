package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCarouselBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCarousel(children ...h.HTMLComponent) (r *VCarouselBuilder) {
	r = &VCarouselBuilder{
		tag: h.Tag("v-carousel").Children(children...),
	}
	return
}

func (b *VCarouselBuilder) ActiveClass(v string) (r *VCarouselBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VCarouselBuilder) Cycle(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":cycle", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) Dark(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) DelimiterIcon(v string) (r *VCarouselBuilder) {
	b.tag.Attr("delimiter-icon", v)
	return b
}

func (b *VCarouselBuilder) Height(v int) (r *VCarouselBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) HideControls(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":hide-controls", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) HideDelimiters(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":hide-delimiters", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) Interval(v int) (r *VCarouselBuilder) {
	b.tag.Attr(":interval", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) Light(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) Mandatory(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) Max(v int) (r *VCarouselBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) Multiple(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) NextIcon(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":next-icon", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) PrevIcon(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":prev-icon", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) Reverse(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) Touch(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr("touch", v)
	return b
}

func (b *VCarouselBuilder) Touchless(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":touchless", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) Value(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(":value", v)
	return b
}

func (b *VCarouselBuilder) Vertical(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) Class(names ...string) (r *VCarouselBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCarouselBuilder) ClassIf(name string, add bool) (r *VCarouselBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCarouselBuilder) On(name string, value string) (r *VCarouselBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCarouselBuilder) Bind(name string, value string) (r *VCarouselBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCarouselBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
