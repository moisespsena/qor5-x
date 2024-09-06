package vuetifyx

import (
	v "github.com/qor5/x/v3/ui/vuetify"
	h "github.com/theplant/htmlgo"
)

type VXLinkageSelectBuilber struct {
	v.VTagBuilder[*VXLinkageSelectBuilber]
}

func VXLinkageSelect() *VXLinkageSelectBuilber {
	return v.VTag(&VXLinkageSelectBuilber{}, "vx-linkageselect")
}

type LinkageSelectItem struct {
	ID          string
	Name        string
	ChildrenIDs []string
}

func (b *VXLinkageSelectBuilber) Items(vs ...[]*LinkageSelectItem) (r *VXLinkageSelectBuilber) {
	b.Attr(":items", vs)
	return b
}

func (b *VXLinkageSelectBuilber) Labels(vs ...string) (r *VXLinkageSelectBuilber) {
	b.Attr(":labels", vs)
	return b
}

func (b *VXLinkageSelectBuilber) ErrorMessages(vs ...[]string) (r *VXLinkageSelectBuilber) {
	b.Attr(":error-messages", vs)
	return b
}

func (b *VXLinkageSelectBuilber) Disabled(v bool) (r *VXLinkageSelectBuilber) {
	b.Attr(":disabled", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectBuilber) SelectOutOfOrder(v bool) (r *VXLinkageSelectBuilber) {
	b.Attr(":select-out-of-order", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectBuilber) Chips(v bool) (r *VXLinkageSelectBuilber) {
	b.Attr(":chips", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectBuilber) Row(v bool) (r *VXLinkageSelectBuilber) {
	b.Attr(":row", h.JSONString(v))
	return b
}
