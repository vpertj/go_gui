package about

import (
	"gioui.org/io/clipboard"
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
	eliasCopyButton, chrisCopyButtonGH, chrisCopyButtonLP widget.Clickable
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
		Name: "About this library",
		Icon: icon.OtherIcon,
	}
}

const (
	sponsorEliasURL          = "https://github.com/sponsors/eliasnaur"
	sponsorChrisURLGitHub    = "https://github.com/sponsors/whereswaldon"
	sponsorChrisURLLiberapay = "https://liberapay.com/whereswaldon/"
)

func (p *Page) Layout(gtx C, th *material.Theme) D {
	p.List.Axis = layout.Vertical
	return material.List(th, &p.List).Layout(gtx, 1, func(gtx C, _ int) D {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return alo.DefaultInset.Layout(gtx, material.Body1(th, `该库使用 https://gioui.org 实现 https://material.io 的材料设计组件。`).Layout)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return alo.DetailRow{}.Layout(gtx,
					material.Body1(th, "Elias Naur 可以在 GitHub 上获得赞助 "+sponsorEliasURL).Layout,
					func(gtx C) D {
						if p.eliasCopyButton.Clicked() {
							clipboard.WriteOp{
								Text: sponsorEliasURL,
							}.Add(gtx.Ops)
						}
						return material.Button(th, &p.eliasCopyButton, "复制赞助网址").Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return alo.DetailRow{}.Layout(gtx,
					material.Body1(th, "Chris Waldon 可以在 GitHub 上获得赞助，网址为 "+sponsorChrisURLGitHub+" 并在自由派上 "+sponsorChrisURLLiberapay).Layout,

					func(gtx C) D {
						if p.chrisCopyButtonGH.Clicked() {
							clipboard.WriteOp{Text: sponsorChrisURLGitHub}.Add(gtx.Ops)
						}
						if p.chrisCopyButtonLP.Clicked() {
							clipboard.WriteOp{Text: sponsorChrisURLLiberapay}.Add(gtx.Ops)
						}
						return alo.DefaultInset.Layout(gtx, func(gtx C) D {
							return layout.Flex{}.Layout(gtx,
								layout.Flexed(.5, material.Button(th, &p.chrisCopyButtonGH, "复制 GitHub 网址").Layout),
								layout.Flexed(.5, material.Button(th, &p.chrisCopyButtonLP, "复制自由支付网址").Layout),
							)
						})
					})
			}),
		)
	})
}
