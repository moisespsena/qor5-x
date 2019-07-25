package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListGroupBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListGroup(children ...h.HTMLComponent) (r *VListGroupBuilder) {
	r = &VListGroupBuilder{
		tag: h.Tag("v-list-group").Children(children...),
	}
	return
}
func (b *VListGroupBuilder) ActiveClass(v string) (r *VListGroupBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VListGroupBuilder) AppendIcon(v string) (r *VListGroupBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VListGroupBuilder) Color(v string) (r *VListGroupBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VListGroupBuilder) Disabled(v bool) (r *VListGroupBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VListGroupBuilder) Eager(v bool) (r *VListGroupBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VListGroupBuilder) Group(v string) (r *VListGroupBuilder) {
	b.tag.Attr("group", v)
	return b
}

func (b *VListGroupBuilder) NoAction(v bool) (r *VListGroupBuilder) {
	b.tag.Attr(":no-action", fmt.Sprint(v))
	return b
}

func (b *VListGroupBuilder) PrependIcon(v string) (r *VListGroupBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VListGroupBuilder) Ripple(v bool) (r *VListGroupBuilder) {
	b.tag.Attr(":ripple", fmt.Sprint(v))
	return b
}

func (b *VListGroupBuilder) SubGroup(v bool) (r *VListGroupBuilder) {
	b.tag.Attr(":sub-group", fmt.Sprint(v))
	return b
}

func (b *VListGroupBuilder) Value(v bool) (r *VListGroupBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VListGroupBuilder) Children(children ...h.HTMLComponent) (r *VListGroupBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListGroupBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListGroupBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListGroupBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListGroupBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListGroupBuilder) Class(names ...string) (r *VListGroupBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListGroupBuilder) ClassIf(name string, add bool) (r *VListGroupBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListGroupBuilder) On(name string, value string) (r *VListGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListGroupBuilder) Bind(name string, value string) (r *VListGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListGroupBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
