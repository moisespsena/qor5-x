package presets

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
)

type ModelBuilder struct {
	p            *Builder
	model        interface{}
	modelType    reflect.Type
	uriName      string
	label        string
	fieldLabels  []string
	placeholders []string
	listing      *ListingBuilder
	editing      *EditingBuilder
	detailing    *DetailingBuilder
}

func NewModelBuilder(p *Builder, model interface{}) (r *ModelBuilder) {
	r = &ModelBuilder{p: p, model: model}
	r.modelType = reflect.TypeOf(model)
	if r.modelType.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("model %#+v must be pointer", model))
	}
	modelstr := r.modelType.String()
	r.label = strcase.ToCamel(modelstr[strings.LastIndex(modelstr, "."):])
	r.newListing()
	r.newDetailing()
	r.newEditing()
	r.inspectModel()
	return
}

type SearchParams struct {
	KeywordColumns []string
	Keyword        string
	Params         url.Values
}

type Searcher func(obj interface{}, params *SearchParams) (r interface{}, err error)
type Fetcher func(obj interface{}, id string) (r interface{}, err error)
type Updater func(obj interface{}, id string, fieldName string, value interface{}) (err error)
type Saver func(obj interface{}, id string) (err error)

func (b *ModelBuilder) newModel() (r interface{}) {
	return reflect.New(b.modelType.Elem()).Interface()
}

func (b *ModelBuilder) newModelArray() (r interface{}) {
	return reflect.New(reflect.SliceOf(b.modelType)).Interface()
}

func (b *ModelBuilder) inspectModel() {
	v := reflect.ValueOf(b.model)

	for v.Elem().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v = v.Elem()

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		//fmt.Println(f.Name, f.Type)
		ft := b.p.fieldTypeByType(f.Type)
		b.listing.Field(f.Name).ComponentFunc(ft.listingCompFunc)
		b.detailing.Field(f.Name).ComponentFunc(ft.detailingCompFunc)
		b.editing.Field(f.Name).ComponentFunc(ft.editingCompFunc)
	}
}

func (b *ModelBuilder) newListing() (r *ListingBuilder) {
	b.listing = &ListingBuilder{filtering: &FilteringBuilder{}, mb: b}
	if b.p.dataOperator != nil {
		b.listing.Searcher(b.p.dataOperator.Search)
	}
	return
}

func (b *ModelBuilder) newEditing() (r *EditingBuilder) {
	b.editing = &EditingBuilder{mb: b}
	if b.p.dataOperator != nil {
		b.editing.Fetcher(b.p.dataOperator.Fetch)
		b.editing.Saver(b.p.dataOperator.Save)
	}
	return
}

func (b *ModelBuilder) newDetailing() (r *DetailingBuilder) {
	b.detailing = &DetailingBuilder{mb: b}
	if b.p.dataOperator != nil {
		b.detailing.Fetcher(b.p.dataOperator.Fetch)
	}
	return
}

func (b *ModelBuilder) URIName(v string) (r *ModelBuilder) {
	b.uriName = v
	return b
}

func (b *ModelBuilder) Label(v string) (r *ModelBuilder) {
	b.label = v
	return b
}

func (b *ModelBuilder) Labels(vs ...string) (r *ModelBuilder) {
	b.fieldLabels = append(b.fieldLabels, vs...)
	return b
}

func (b *ModelBuilder) Placeholders(vs ...string) (r *ModelBuilder) {
	b.placeholders = append(b.placeholders, vs...)
	return b
}

func (b *ModelBuilder) getComponentFuncField(field *FieldBuilder) (r *Field) {
	r = &Field{
		Name:  field.name,
		Label: b.getLabel(field),
	}
	return
}

func (b *ModelBuilder) getLabel(field *FieldBuilder) (r string) {
	if len(field.label) > 0 {
		return field.label
	}

	for i := 0; i < len(b.fieldLabels)-1; i = i + 2 {
		if b.fieldLabels[i] == field.name {
			return b.fieldLabels[i+1]
		}
	}

	return field.name
}
