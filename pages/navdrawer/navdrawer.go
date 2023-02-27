package navdrawer

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	alo "go_gui/applayout"
	"go_gui/icon"
	page "go_gui/pages"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// Page holds the state for a page demonstrating the features of
// the NavDrawer component.
type Page struct {
	nonModalDrawer widget.Bool
	widget.List
	*page.Router
}

// New constructs a Page with the provided router.
func New(router *page.Router) *Page {
	return &Page{
		Router: router,
	}
}

var _ page.Page = &Page{}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "导航抽屉特点",
		Icon: icon.SettingsIcon,
	}
}

func (p *Page) Layout(gtx C, th *material.Theme) D {
	p.List.Axis = layout.Vertical
	return material.List(th, &p.List).Layout(gtx, 1, func(gtx C, _ int) D {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return alo.DefaultInset.Layout(gtx, material.Body1(th, `导航抽屉小部件为导航提供了一致的界面元素。

通过下面的控件，您可以查看我们的导航抽屉实现中可用的各种功能。`).Layout)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return alo.DetailRow{}.Layout(gtx,
					material.Body1(th, "使用非模态抽屉").Layout,
					func(gtx C) D {
						if p.nonModalDrawer.Changed() {
							p.Router.NonModalDrawer = p.nonModalDrawer.Value
							if p.nonModalDrawer.Value {
								p.Router.NavAnim.Appear(gtx.Now)
							} else {
								p.Router.NavAnim.Disappear(gtx.Now)

							}
						}
						return material.Switch(th, &p.nonModalDrawer, "使用非模式导航抽屉").Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return alo.DetailRow{}.Layout(gtx,
					material.Body1(th, "拖动以关闭").Layout,
					material.Body2(th, "您可以通过向左拖动来关闭模式导航抽屉.").Layout)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return alo.DetailRow{}.Layout(gtx,
					material.Body1(th, "触摸稀松布以关闭").Layout,
					material.Body2(th, "您可以关闭模态导航抽屉，触摸右侧半透明稀松布中的任何位置.").Layout)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return alo.DetailRow{}.Layout(gtx,
					material.Body1(th, "底部内容锚定").Layout,
					material.Body2(th, "如果在应用栏设置中切换对底部应用栏的支持，则导航抽屉内容将锚定到抽屉区域的底部而不是顶部.").Layout)
			}),
		)
	})
}
