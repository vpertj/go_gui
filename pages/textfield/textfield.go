package textfield

import (
	alo "go_gui/applayout"
	"go_gui/icon"
	page "go_gui/pages"
	"image/color"
	"unicode"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// Page holds the state for a page demonstrating the features of
// the TextField component.
type Page struct {
	inputAlignment                                               layout.Alignment
	inputAlignmentEnum                                           widget.Enum
	nameInput, addressInput, priceInput, tweetInput, numberInput component.TextField
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
		Name: "文本字段功能",
		Icon: icon.EditIcon,
	}
}

func (p *Page) Layout(gtx C, th *material.Theme) D {
	p.List.Axis = layout.Vertical
	return material.List(th, &p.List).Layout(gtx, 1, func(gtx C, _ int) D {
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(
			gtx,
			layout.Rigid(func(gtx C) D {
				p.nameInput.Alignment = p.inputAlignment
				return p.nameInput.Layout(gtx, th, "姓名")
			}),
			layout.Rigid(func(gtx C) D {
				return alo.DefaultInset.Layout(gtx, material.Body2(th, "响应悬停事件.").Layout)
			}),
			layout.Rigid(func(gtx C) D {
				p.addressInput.Alignment = p.inputAlignment
				return p.addressInput.Layout(gtx, th, "地址")
			}),
			layout.Rigid(func(gtx C) D {
				return alo.DefaultInset.Layout(gtx, material.Body2(th, "当您单击以选择文本字段时，标签会正确动画化。").Layout)
			}),
			layout.Rigid(func(gtx C) D {
				p.priceInput.Prefix = func(gtx C) D {
					th := *th
					th.Palette.Fg = color.NRGBA{R: 100, G: 100, B: 100, A: 255}
					return material.Label(&th, th.TextSize, "$").Layout(gtx)
				}
				p.priceInput.Suffix = func(gtx C) D {
					th := *th
					th.Palette.Fg = color.NRGBA{R: 100, G: 100, B: 100, A: 255}
					return material.Label(&th, th.TextSize, ".00").Layout(gtx)
				}
				p.priceInput.SingleLine = true
				p.priceInput.Alignment = p.inputAlignment
				return p.priceInput.Layout(gtx, th, "价格")
			}),
			layout.Rigid(func(gtx C) D {
				return alo.DefaultInset.Layout(gtx, material.Body2(th, "可以有前缀和后缀元素.").Layout)
			}),
			layout.Rigid(func(gtx C) D {
				if err := func() string {
					for _, r := range p.numberInput.Text() {
						if !unicode.IsDigit(r) {
							return "必须仅包含数字"
						}
					}
					return ""
				}(); err != "" {
					p.numberInput.SetError(err)
				} else {
					p.numberInput.ClearError()
				}
				p.numberInput.SingleLine = true
				p.numberInput.Alignment = p.inputAlignment
				return p.numberInput.Layout(gtx, th, "数字")
			}),
			layout.Rigid(func(gtx C) D {
				return alo.DefaultInset.Layout(gtx, material.Body2(th, "可以验证.").Layout)
			}),
			layout.Rigid(func(gtx C) D {
				if p.tweetInput.TextTooLong() {
					p.tweetInput.SetError("字符过多")
				} else {
					p.tweetInput.ClearError()
				}
				p.tweetInput.CharLimit = 128
				p.tweetInput.Helper = "推文的字符数有限"
				p.tweetInput.Alignment = p.inputAlignment
				return p.tweetInput.Layout(gtx, th, "编辑 播报")
			}),
			layout.Rigid(func(gtx C) D {
				return alo.DefaultInset.Layout(gtx, material.Body2(th, "可以有字符计数器和帮助文本.").Layout)
			}),
			layout.Rigid(func(gtx C) D {
				if p.inputAlignmentEnum.Changed() {
					switch p.inputAlignmentEnum.Value {
					case layout.Start.String():
						p.inputAlignment = layout.Start
					case layout.Middle.String():
						p.inputAlignment = layout.Middle
					case layout.End.String():
						p.inputAlignment = layout.End
					default:
						p.inputAlignment = layout.Start
					}
					op.InvalidateOp{}.Add(gtx.Ops)
				}
				return alo.DefaultInset.Layout(
					gtx,
					func(gtx C) D {
						return layout.Flex{
							Axis: layout.Vertical,
						}.Layout(
							gtx,
							layout.Rigid(func(gtx C) D {
								return material.Body2(th, "Text Alignment-文本对齐方式").Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return layout.Flex{
									Axis: layout.Vertical,
								}.Layout(
									gtx,
									layout.Rigid(func(gtx C) D {
										return material.RadioButton(
											th,
											&p.inputAlignmentEnum,
											layout.Start.String(),
											"开始位置",
										).Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.RadioButton(
											th,
											&p.inputAlignmentEnum,
											layout.Middle.String(),
											"中间位置",
										).Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.RadioButton(
											th,
											&p.inputAlignmentEnum,
											layout.End.String(),
											"最后位置",
										).Layout(gtx)
									}),
								)
							}),
						)
					},
				)
			}),
			layout.Rigid(func(gtx C) D {
				return alo.DefaultInset.Layout(gtx, material.Body2(th, "此文本字段实现由 Jack Mordaunt 提供。谢谢杰克!").Layout)
			}),
		)
	})
}
