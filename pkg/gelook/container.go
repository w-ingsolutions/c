// SPDX-License-Identifier: Unlicense OR MIT

package gelook

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/text"
	"gioui.org/unit"
)

type WingUIcontainer struct {
	// Color is the text color.
	Color         string
	Font          text.Font
	TextSize      unit.Value
	Background    string
	TxColor       string
	BgColor       string
	TxColorHover  string
	BgColorHover  string
	FullWidth     bool
	Width         int
	Height        int
	CornerRadius  int
	PaddingTop    int
	PaddingRight  int
	PaddingBottom int
	PaddingLeft   int
	shaper        text.Shaper
	link          bool
	hover         bool
}

func (t *WingUItheme) WingUIcontainer(padding int, background string) WingUIcontainer {
	return WingUIcontainer{
		Font: text.Font{
			Typeface: t.Fonts["Primary"],
		},
		// Color:      rgb(0xffffff),
		PaddingTop:    padding,
		PaddingRight:  padding,
		PaddingBottom: padding,
		PaddingLeft:   padding,
		Background:    background,
		TextSize:      t.TextSize.Scale(14.0 / 16.0),
		shaper:        t.Shaper,
	}
}

func (w WingUIcontainer) Layout(g layout.Context, direction layout.Direction, itemContent func(gtx C) D) layout.Dimensions {
	hmin := g.Constraints.Min.X
	vmin := g.Constraints.Min.Y
	if w.FullWidth {
		hmin = g.Constraints.Max.Y
	}
	return layout.Stack{Alignment: layout.W}.Layout(g,
		layout.Expanded(func(gtx C) D {
			rr := float32(gtx.Px(unit.Dp(float32(w.CornerRadius))))
			clip.Rect{
				Rect: f32.Rectangle{Max: f32.Point{
					X: float32(gtx.Constraints.Min.X),
					Y: float32(gtx.Constraints.Min.Y),
				}},
				NE: rr, NW: rr, SE: rr, SW: rr,
			}.Op(gtx.Ops).Add(gtx.Ops)
			return fill(gtx, HexARGB(w.Background))
			//pointer.Rect(image.Rectangle{Max: g.Dimensions.Size}).Add(g.Ops)
		}),
		layout.Stacked(func(gtx C) D {
			gtx.Constraints.Min.Y = hmin
			gtx.Constraints.Min.Y = vmin
			return direction.Layout(gtx, func(gtx C) D {
				return layout.Inset{
					Top:    unit.Dp(float32(w.PaddingTop)),
					Right:  unit.Dp(float32(w.PaddingRight)),
					Bottom: unit.Dp(float32(w.PaddingBottom)),
					Left:   unit.Dp(float32(w.PaddingLeft)),
				}.Layout(gtx, itemContent)
			})
		}),
	)
}
