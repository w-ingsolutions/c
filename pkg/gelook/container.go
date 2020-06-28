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

func (t *WingUItheme) DuoUIcontainer(padding int, background string) WingUIcontainer {
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

func (d WingUIcontainer) Layout(g *layout.Context, direction layout.Direction, itemContent func(gtx C) D) {
	hmin := g.Constraints.Min.X
	vmin := g.Constraints.Min.Y
	if d.FullWidth {
		hmin = g.Constraints.Max.Y
	}
	layout.Stack{Alignment: layout.W}.Layout(*g,
		layout.Expanded(func(gtx C) D {
			var dd D
			rr := float32(g.Px(unit.Dp(float32(d.CornerRadius))))
			clip.Rect{
				Rect: f32.Rectangle{Max: f32.Point{
					X: float32(g.Constraints.Min.X),
					Y: float32(g.Constraints.Min.Y),
				}},
				NE: rr, NW: rr, SE: rr, SW: rr,
			}.Op(g.Ops).Add(g.Ops)
			fill(*g, HexARGB(d.Background))
			//pointer.Rect(image.Rectangle{Max: g.Dimensions.Size}).Add(g.Ops)
			return dd
		}),
		layout.Stacked(func(gtx C) D {
			var dd D
			gtx.Constraints.Min.Y = hmin
			gtx.Constraints.Min.Y = vmin
			direction.Layout(gtx, func(gtx C) D {
				var ddd D
				layout.Inset{
					Top:    unit.Dp(float32(d.PaddingTop)),
					Right:  unit.Dp(float32(d.PaddingRight)),
					Bottom: unit.Dp(float32(d.PaddingBottom)),
					Left:   unit.Dp(float32(d.PaddingLeft)),
				}.Layout(gtx, itemContent)
				return ddd
			})
			return dd

		}),
	)
}
