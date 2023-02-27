package appbar

import (
	"image/color"

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
// the AppBar component.
type Page struct {
	heartBtn, plusBtn, contextBtn          widget.Clickable
	exampleOverflowState, red, green, blue widget.Clickable
	bottomBar, customNavIcon               widget.Bool
	favorited                              bool
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
	return []component.AppBarAction{
		{
			OverflowAction: component.OverflowAction{
				Name: "喜欢点赞",
				Tag:  &p.heartBtn,
			},
			Layout: func(gtx layout.Context, bg, fg color.NRGBA) layout.Dimensions {
				if p.heartBtn.Clicked() {
					p.favorited = !p.favorited
				}
				btn := component.SimpleIconButton(bg, fg, &p.heartBtn, icon.HeartIcon)
				btn.Background = bg
				if p.favorited {
					btn.Color = color.NRGBA{R: 200, A: 255}
				} else {
					btn.Color = fg
				}
				return btn.Layout(gtx)
			},
		},
		component.SimpleIconAction(&p.plusBtn, icon.PlusIcon,
			component.OverflowAction{
				Name: "Create",
				Tag:  &p.plusBtn,
			},
		),
	}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{
		{
			Name: "示例 1",
			Tag:  &p.exampleOverflowState,
		},
		{
			Name: "示例 2",
			Tag:  &p.exampleOverflowState,
		},
	}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "应用栏功能",
		Icon: icon.HomeIcon,
	}
}

const (
	settingNameColumnWidth    = .3
	settingDetailsColumnWidth = 1 - settingNameColumnWidth
)

func (p *Page) Layout(gtx C, th *material.Theme) D {
	p.List.Axis = layout.Vertical
	return material.List(th, &p.List).Layout(gtx, 1, func(gtx C, _ int) D {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return alo.DefaultInset.Layout(gtx, material.Body1(th, `应用栏小组件提供一致的界面元素，用于触发导航和特定于页面的操作。

通过下面的控件，可以查看应用栏实现中可用的各种功能.`).Layout)
			}),
			layout.Rigid(func(gtx C) D {
				return alo.DetailRow{}.Layout(gtx, material.Body1(th, "上下文应用栏").Layout, func(gtx C) D {
					if p.contextBtn.Clicked() {
						p.Router.AppBar.SetContextualActions(
							[]component.AppBarAction{
								component.SimpleIconAction(&p.red, icon.HeartIcon,
									component.OverflowAction{
										Name: "主页",
										Tag:  &p.red,
									},
								),
							},
							[]component.OverflowAction{
								{
									Name: "foo",
									Tag:  &p.blue,
								},
								{
									Name: "bar",
									Tag:  &p.green,
								},
							},
						)
						p.Router.AppBar.ToggleContextual(gtx.Now, "上下文标题")
					}
					return material.Button(th, &p.contextBtn, "触发").Layout(gtx)
				})
			}),
			layout.Rigid(func(gtx C) D {
				return alo.DetailRow{}.Layout(gtx,
					material.Body1(th, "底部应用栏").Layout,
					func(gtx C) D {
						if p.bottomBar.Changed() {
							if p.bottomBar.Value {
								p.Router.ModalNavDrawer.Anchor = component.Bottom
								p.Router.AppBar.Anchor = component.Bottom
							} else {
								p.Router.ModalNavDrawer.Anchor = component.Top
								p.Router.AppBar.Anchor = component.Top
							}
							p.Router.BottomBar = p.bottomBar.Value
						}

						return material.Switch(th, &p.bottomBar, "使用底部应用栏").Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx C) D {
				return alo.DetailRow{}.Layout(gtx,
					material.Body1(th, "自定义导航图标").Layout,
					func(gtx C) D {
						if p.customNavIcon.Changed() {
							if p.customNavIcon.Value {
								p.Router.AppBar.NavigationIcon = icon.HomeIcon
							} else {
								p.Router.AppBar.NavigationIcon = icon.MenuIcon
							}
						}
						return material.Switch(th, &p.customNavIcon, "使用自定义导航图标").Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx C) D {
				return alo.DetailRow{}.Layout(gtx,
					material.Body1(th, "动画调整大小").Layout,
					material.Body2(th, "调整屏幕宽度，以查看应用栏操作折叠到溢出菜单中或从溢出菜单中显示（如果大小允许）。").Layout,
				)
			}),
			layout.Rigid(func(gtx C) D {
				return alo.DetailRow{}.Layout(gtx,
					material.Body1(th, "自定义操作按钮").Layout,
					material.Body2(th, "单击心形操作以查看自定义按钮行为。").Layout)
			}),
		)
	})
}
