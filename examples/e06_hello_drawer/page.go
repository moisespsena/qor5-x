package e06_hello_drawer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sunfmin/bran/ui"
	bo "github.com/sunfmin/branoverlay"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	drawerVisible bool
	Name          string
	NameError     string
}

func randStr(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s: %d", prefix, rand.Int31n(100))
}

func HelloDrawer(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	s := ctx.StateOrInit(&mystate{}).(*mystate)

	pr.Schema = Div(
		H1(s.Name),
		bo.Drawer(
			Button("Close").Attr("@click", "parent.close"),
			Div(Text(randStr("homeDrawer"))),
			bo.LazyLoader(ctx.Hub, "editPage", editPage, "param1").ParentVisible(),
			ui.Bind(Input("").Type("text").Value(s.Name)).FieldName("Name"),
			Label(s.NameError).Style("color:red"),
			ui.Bind(Button("Update")).OnClick(ctx.Hub, "update", update),
		).Trigger(
			A().Text("Edit").Href("#"),
		).Width(500).DefaultOpen(s.drawerVisible),
	)
	return
}

func update(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true
	s := ctx.State.(*mystate)
	if len(s.Name) < 10 {
		s.NameError = "name is too short"
		s.drawerVisible = true
		s.Name = ""
	} else {
		s.drawerVisible = false
	}
	return
}

func editPage(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	s := ctx.State.(*mystate)
	_ = s

	r.Schema = bo.Drawer(
		Button("Close").Attr("@click", "parent.close"),
		H1(ctx.Event.Params[0]),
		Div(Text(randStr("in editPage Drawer"))),
	).Trigger(
		A().Text("Open " + randStr("inner")).Href("#"),
	).Width(400)
	return
}
