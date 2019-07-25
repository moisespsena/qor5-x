package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListItemActionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListItemAction(children ...h.HTMLComponent) (r *VListItemActionBuilder) {
	r = &VListItemActionBuilder{
		tag: h.Tag("v-list-item-action").Children(children...),
	}
	return
}

func (b *VListItemActionBuilder) Children(children ...h.HTMLComponent) (r *VListItemActionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListItemActionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListItemActionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListItemActionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListItemActionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListItemActionBuilder) Class(names ...string) (r *VListItemActionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListItemActionBuilder) ClassIf(name string, add bool) (r *VListItemActionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListItemActionBuilder) On(name string, value string) (r *VListItemActionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListItemActionBuilder) Bind(name string, value string) (r *VListItemActionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListItemActionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
