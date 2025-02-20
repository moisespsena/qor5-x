package vuetifyx

import (
	v "github.com/qor5/x/v3/ui/vuetify"
	h "github.com/theplant/htmlgo"
)

type VXTreeDataTableBuilder struct {
	v.VTagBuilder[*VXTreeDataTableBuilder]
	items interface{}
	many  bool
}

func VXTreeDataTable(children ...h.HTMLComponent) *VXTreeDataTableBuilder {
	return v.VTag(&VXTreeDataTableBuilder{}, "vx-tree-data-table", children...)
}

func (b *VXTreeDataTableBuilder) Items(v interface{}) (r *VXTreeDataTableBuilder) {
	return b.Attr(":items", h.JSONString(v))
}

func (b *VXTreeDataTableBuilder) ItemsVar(v string) (r *VXTreeDataTableBuilder) {
	return b.Attr(":items", v)
}

func (b *VXTreeDataTableBuilder) Headers(v interface{}) (r *VXTreeDataTableBuilder) {
	return b.Attr(":headers", h.JSONString(v))
}

func (b *VXTreeDataTableBuilder) HeadersVar(v string) (r *VXTreeDataTableBuilder) {
	return b.Attr(":headers", v)
}

func (b *VXTreeDataTableBuilder) SettingsTitle(v string) *VXTreeDataTableBuilder {
	return b.Attr("settings-title", v)
}

func (b *VXTreeDataTableBuilder) SettingsColumnsTitle(v string) *VXTreeDataTableBuilder {
	return b.Attr("settings-columns-title", v)
}
