package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VAutocompleteBuilder struct {
	tag           *h.HTMLTagBuilder
	selectedItems interface{}
	items         interface{}
}

func VAutocomplete(children ...h.HTMLComponent) (r *VAutocompleteBuilder) {
	r = &VAutocompleteBuilder{
		tag: h.Tag("vw-autocomplete").Children(children...),
	}
	r.Multiple(true)

	return
}
func (b *VAutocompleteBuilder) AllowOverflow(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":allow-overflow", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) AppendIcon(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VAutocompleteBuilder) AppendOuterIcon(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("append-outer-icon", v)
	return b
}

func (b *VAutocompleteBuilder) Attach(v interface{}) (r *VAutocompleteBuilder) {
	b.tag.Attr(":attach", v)
	return b
}

func (b *VAutocompleteBuilder) AutoSelectFirst(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":auto-select-first", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Autofocus(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) BackgroundColor(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VAutocompleteBuilder) Box(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":box", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) BrowserAutocomplete(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("browser-autocomplete", v)
	return b
}

func (b *VAutocompleteBuilder) CacheItems(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":cache-items", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Chips(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":chips", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) ClearIcon(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("clear-icon", v)
	return b
}

func (b *VAutocompleteBuilder) Clearable(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Color(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VAutocompleteBuilder) Counter(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":counter", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Dark(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) DeletableChips(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":deletable-chips", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Dense(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Disabled(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) DontFillMaskBlanks(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":dont-fill-mask-blanks", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Error(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) ErrorCount(v int) (r *VAutocompleteBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) ErrorMessages(v ...string) (r *VAutocompleteBuilder) {
	if len(v) == 0 {
		return b
	}

	b.tag.Attr(":error-messages", h.JSONString(v))
	return b
}

func (b *VAutocompleteBuilder) Flat(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) FullWidth(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Height(v int) (r *VAutocompleteBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) HideDetails(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) HideNoData(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":hide-no-data", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) HideSelected(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":hide-selected", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Hint(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VAutocompleteBuilder) ItemAvatar(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("item-avatar", v)
	return b
}

func (b *VAutocompleteBuilder) ItemDisabled(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("item-disabled", v)
	return b
}

func (b *VAutocompleteBuilder) ItemText(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("item-text", v)
	return b
}

func (b *VAutocompleteBuilder) ItemValue(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("item-value", v)
	return b
}

func (b *VAutocompleteBuilder) Items(v interface{}) (r *VAutocompleteBuilder) {
	b.items = v
	return b
}

func (b *VAutocompleteBuilder) Label(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VAutocompleteBuilder) Light(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Loading(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Mask(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("mask", v)
	return b
}

func (b *VAutocompleteBuilder) MenuProps(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("menu-props", v)
	return b
}

func (b *VAutocompleteBuilder) Messages(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VAutocompleteBuilder) Multiple(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) NoDataText(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("no-data-text", v)
	return b
}

func (b *VAutocompleteBuilder) NoFilter(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":no-filter", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) OpenOnClear(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":open-on-clear", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Outline(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":outline", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) PersistentHint(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Placeholder(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VAutocompleteBuilder) Prefix(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("prefix", v)
	return b
}

func (b *VAutocompleteBuilder) PrependIcon(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VAutocompleteBuilder) PrependInnerIcon(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("prepend-inner-icon", v)
	return b
}

func (b *VAutocompleteBuilder) Readonly(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) ReturnMaskedValue(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":return-masked-value", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) ReturnObject(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":return-object", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Reverse(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Rules(v []string) (r *VAutocompleteBuilder) {
	b.tag.Attr(":rules", v)
	return b
}

func (b *VAutocompleteBuilder) SearchInput(v interface{}) (r *VAutocompleteBuilder) {
	b.tag.Attr(":search-input", v)
	return b
}

func (b *VAutocompleteBuilder) SingleLine(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) SmallChips(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":small-chips", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Solo(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":solo", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) SoloInverted(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":solo-inverted", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Success(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) SuccessMessages(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VAutocompleteBuilder) Suffix(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("suffix", v)
	return b
}

func (b *VAutocompleteBuilder) Type(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VAutocompleteBuilder) ValidateOnBlur(v bool) (r *VAutocompleteBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VAutocompleteBuilder) Value(v interface{}) (r *VAutocompleteBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VAutocompleteBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if b.items == nil {
		b.items = b.selectedItems
	}
	b.tag.Attr(":items", b.items)
	b.tag.Attr(":selected-items", b.selectedItems)
	return b.tag.MarshalHTML(ctx)
}
