package discloser

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"go_gui/icon"
	page "go_gui/pages"
)

// TreeNode is a simple tree implementation that holds both
// display data and the state for Discloser widgets. In
// practice, you'll often want to separate the state from
// the data being presented.
type TreeNode struct {
	Text     string
	Children []TreeNode
	component.DiscloserState
}

type (
	C = layout.Context
	D = layout.Dimensions
)

// Page holds the state for a page demonstrating the features of
// the AppBar component.
type Page struct {
	TreeNode
	widget.List
	*page.Router
	CustomDiscloserState component.DiscloserState
}

// New constructs a Page with the provided router.
func New(router *page.Router) *Page {
	return &Page{
		Router: router,
		TreeNode: TreeNode{
			Text: "展开我",
			Children: []TreeNode{
				{
					Text: "可以扩展我...",
					Children: []TreeNode{
						{
							Text: "...嵌套到任意深度.",
						},
						{
							Text: "还有一些类型可用于自定义分离器的外观和感觉:",
							Children: []TreeNode{
								{
									Text: "• DiscloserStyle 允许您提供自己的控件，而不是此处使用的默认三角形.",
								},
								{
									Text: "• DiscloserArrowStyle 允许您更改此处使用的三角形的表示形式，例如更改其颜色、大小、左/右锚定或边距.",
								},
							},
						},
					},
				},
			},
		},
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
		Name: "Disclosers",
		Icon: icon.VisibilityIcon,
	}
}

// LayoutTreeNode recursively lays out a tree of widgets described by
// TreeNodes.
func (p *Page) LayoutTreeNode(gtx C, th *material.Theme, tn *TreeNode) D {
	if len(tn.Children) == 0 {
		return layout.UniformInset(unit.Dp(2)).Layout(gtx,
			material.Body1(th, tn.Text).Layout)
	}
	children := make([]layout.FlexChild, 0, len(tn.Children))
	for i := range tn.Children {
		child := &tn.Children[i]
		children = append(children, layout.Rigid(
			func(gtx C) D {
				return p.LayoutTreeNode(gtx, th, child)
			}))
	}
	return component.SimpleDiscloser(th, &tn.DiscloserState).Layout(gtx,
		material.Body1(th, tn.Text).Layout,
		func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
		})
}

// LayoutCustomDiscloser demonstrates how to create a custom control for
// a discloser.
func (p *Page) LayoutCustomDiscloser(gtx C, th *material.Theme) D {
	return component.Discloser(th, &p.CustomDiscloserState).Layout(gtx,
		func(gtx C) D {
			var l material.LabelStyle
			l = material.Body1(th, "+")
			if p.CustomDiscloserState.Visible() {
				l.Text = "-"
			}
			l.Font.Variant = "Mono"
			return layout.UniformInset(unit.Dp(2)).Layout(gtx, l.Layout)
		},
		material.Body1(th, "自定义控件").Layout,
		material.Body2(th, "此控件仅包含 9 行代码.").Layout,
	)
}

func (p *Page) Layout(gtx C, th *material.Theme) D {
	p.List.Axis = layout.Vertical
	return material.List(th, &p.List).Layout(gtx, 2, func(gtx C, index int) D {
		return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
			if index == 0 {
				return p.LayoutTreeNode(gtx, th, &p.TreeNode)
			}
			return p.LayoutCustomDiscloser(gtx, th)
		})
	})
}
