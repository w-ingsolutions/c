package gelook

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel"
	"image"
	"strconv"
)

type WingUIcounter struct {
	increase     *iconBtn
	decrease     *iconBtn
	reset        *iconBtn
	input        material.EditorStyle
	pageFunction func()
	Font         text.Font
	TextSize     unit.Value
	TxColor      string
	BgColor      string
	BtnBgColor   string
	shaper       text.Shaper
}

func (t *WingUItheme) WingUIcounter(cc *gel.DuoUIcounter, pageFunction func()) WingUIcounter {
	return WingUIcounter{
		// ToDo Replace theme's buttons with counter exclusive buttons, set icons for increase/decrease
		increase:     iconButton(t.T, t.Icons["counterPlusIcon"], cc.CounterIncrease, t.Colors["Primary"]),
		decrease:     iconButton(t.T, t.Icons["counterMinusIcon"], cc.CounterDecrease, t.Colors["Primary"]),
		input:        material.Editor(t.T, cc.CounterInput, "0"),
		pageFunction: pageFunction,
		Font: text.Font{
			Typeface: t.Fonts["Primary"],
		},
		TxColor:    t.Colors["Light"],
		BgColor:    t.Colors["Dark"],
		BtnBgColor: t.Colors["Primary"],
		TextSize:   unit.Dp(float32(18)),
		shaper:     t.Shaper,
	}
}

func (c WingUIcounter) Layout(cc *gel.DuoUIcounter, g layout.Context, th *material.Theme, label, value string) func(gtx C) layout.Dimensions {
	return func(gtx C) D {
		cc.CounterInput.SetText(value)
		hmin := g.Constraints.Min.X
		vmin := g.Constraints.Min.Y
		//txColor := c.TxColor
		bgColor := c.BgColor
		return layout.Stack{Alignment: layout.Center}.Layout(g,
			layout.Expanded(func(gtx C) D {
				rr := float32(g.Px(unit.Dp(0)))
				clip.Rect{
					Rect: f32.Rectangle{Max: f32.Point{
						X: float32(g.Constraints.Min.X),
						Y: float32(g.Constraints.Min.Y),
					}},
					NE: rr, NW: rr, SE: rr, SW: rr,
				}.Op(g.Ops).Add(g.Ops)
				return fill(g, HexARGB(bgColor))
			}),
			layout.Stacked(func(gtx C) D {
				g.Constraints.Min.X = hmin
				g.Constraints.Min.Y = vmin
				return layout.Center.Layout(g, func(gtx C) D {
					return layout.Flex{
						Spacing:   layout.SpaceAround,
						Alignment: layout.Middle,
					}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							for cc.CounterDecrease.Clicked() {
								cc.Decrease()
								c.pageFunction()
							}
							return c.decrease.Layout(gtx)
						}),
						layout.Rigid(func(gtx C) D {
							return layout.Center.Layout(gtx, func(gtx C) D {
								return layout.Inset{
									Top:    unit.Dp(0),
									Right:  unit.Dp(16),
									Bottom: unit.Dp(0),
									Left:   unit.Dp(16),
								}.Layout(gtx, func(gtx C) D {
									return layout.Flex{
										Axis:      layout.Vertical,
										Spacing:   layout.SpaceAround,
										Alignment: layout.Middle,
									}.Layout(gtx,
										layout.Rigid(func(gtx C) D {
											//paint.ColorOp{Color: HexARGB(c.TxColor)}.Add(gtx.Ops)
											//return widget.Label{
											//	Alignment: text.Middle,
											//}.Layout(gtx, c.shaper, c.Font, unit.Dp(8), label)
											return material.Body1(th, label).Layout(gtx)
											//return layout.Dimensions{}
										}),
										layout.Rigid(func(gtx C) D {
											c.input.Font.Typeface = c.Font.Typeface
											c.input.Color = HexARGB(c.TxColor)
											for _, e := range cc.CounterInput.Events() {
												switch e.(type) {
												case widget.ChangeEvent:
													if i, err := strconv.Atoi(cc.CounterInput.Text()); err == nil {
														cc.Value = i
													}
												}
											}
											return c.input.Layout(gtx)
											paint.ColorOp{Color: HexARGB(c.TxColor)}.Add(gtx.Ops)
											return material.Body1(th, value).Layout(gtx)
											//{
											//	Alignment: text.Middle,
											//}.Layout(gtx, c.shaper, c.Font, unit.Dp(12), value)
											//return layout.Dimensions{}
										}))
								})
							})
						}),
						//layout.Flexed(0.2, func() {
						//	for cc.CounterReset.Clicked(gtx) {
						//		cc.Reset()
						//		c.pageFunction()
						//	}
						//	c.reset.Layout(gtx, cc.CounterReset)
						//}),
						//layout.Flexed(0.2, func(gtx C) D {
						//
						//	//paint.ColorOp{Color: HexARGB(c.TxColor)}.Add(gtx.Ops)
						//	//return widget.Label{
						//	//	Alignment: text.Middle,
						//	//}.Layout(gtx, c.shaper, c.Font, unit.Dp(8), label)
						//}),
						//layout.Rigid(func(gtx C) D {
						//	c.input.Font.Typeface = c.Font.Typeface
						//	c.input.Color = HexARGB(c.TxColor)
						//	for _, e := range cc.CounterInput.Events() {
						//		switch e.(type) {
						//		case widget.ChangeEvent:
						//			if i, err := strconv.Atoi(cc.CounterInput.Text()); err == nil {
						//				cc.Value = i
						//			}
						//		}
						//	}
						//	return c.input.Layout(gtx)
						//	// paint.ColorOp{Color: HexARGB(c.TxColor)}.Add(gtx.Ops)
						//	// gel.Label{
						//	//	Alignment: text.Middle,
						//	// }.Layout(gtx, c.shaper, c.Font, unit.Dp(12), value)
						//}),
						layout.Rigid(func(gtx C) D {
							for cc.CounterIncrease.Clicked() {
								cc.Increase()
								c.pageFunction()
							}
							return c.increase.Layout(gtx)
						}))
				})
			}),
		)

	}
}

type iconBtn struct {
	theme  *material.Theme
	button *widget.Clickable
	icon   *widget.Icon
	word   string
	bg     string
}

func iconButton(t *material.Theme, i *widget.Icon, c *widget.Clickable, bg string) *iconBtn {
	return &iconBtn{
		theme:  t,
		button: c,
		icon:   i,
		bg:     bg,
	}
}

func (b iconBtn) Layout(gtx layout.Context) layout.Dimensions {
	btn := material.ButtonLayout(b.theme, b.button)
	btn.CornerRadius = unit.Dp(0)
	btn.Background = HexARGB(b.bg)
	return btn.Layout(gtx, func(gtx C) D {
		return layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx C) D {
			iconAndLabel := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}
			layIcon := layout.Rigid(func(gtx C) D {
				return layout.Inset{}.Layout(gtx, func(gtx C) D {
					var d D
					if b.icon != nil {
						size := gtx.Px(unit.Dp(24))
						b.icon.Layout(gtx, unit.Px(float32(size)))
						d = layout.Dimensions{
							Size: image.Point{X: size, Y: size},
						}
					}
					return d
				})
			})

			return iconAndLabel.Layout(gtx, layIcon)
		})
	})
}
