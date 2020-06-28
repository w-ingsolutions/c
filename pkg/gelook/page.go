// SPDX-License-Identifier: Unlicense OR MIT

package gelook

import (
	"image"

	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/text"
	"gioui.org/unit"
)

type WingUIpage struct {
	Title   string
	TxColor string
	// Font          text.Font
	shaper  text.Shaper
	Command func(gtx C) D

	Header        func(gtx C) D
	HeaderBgColor string
	HeaderPadding float32
	// header
	// header
	Border      float32
	BorderColor string
	Body        func(gtx C) D
	BodyBgColor string
	BodyPadding float32
	// body
	// body
	Footer        func(gtx C) D
	FooterBgColor string
	FooterPadding float32
	// footer
	// footer
}

func (t *WingUItheme) WingUIpage(p WingUIpage) *WingUIpage {
	return &WingUIpage{
		Title: p.Title,
		// Font: text.Font{
		// Size: t.TextSize.Scale(14.0 / 16.0),
		// },
		TxColor:       t.Colors["Dark"],
		shaper:        t.Shaper,
		Command:       p.Command,
		Header:        p.Header,
		HeaderBgColor: t.Colors["Primary"],
		HeaderPadding: p.HeaderPadding,
		Border:        p.Border,
		BorderColor:   p.BorderColor,
		Body:          p.Body,
		BodyBgColor:   p.BodyBgColor,
		BodyPadding:   p.BodyPadding,
		Footer:        p.Footer,
		FooterBgColor: t.Colors["Secondary"],
		FooterPadding: p.FooterPadding,
	}
}

func (p WingUIpage) Layout(g *layout.Context) {
	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(*g,
		layout.Rigid(pageElementLayout(g, layout.N, p.HeaderBgColor, p.HeaderPadding, p.Header)),
		layout.Flexed(1, func(gtx C) D {
			var dd D
			WingUIfill(g, p.BorderColor)
			layout.UniformInset(unit.Dp(p.Border)).Layout(*g, pageElementLayout(g, layout.N, p.BodyBgColor, p.BodyPadding, p.Body))
			return dd
		}),
		layout.Rigid(pageElementLayout(g, layout.N, p.FooterBgColor, p.FooterPadding, p.Footer)),
	)
}

func pageElementLayout(g *layout.Context, direction layout.Direction, background string, padding float32, elementContent func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		var dd D
		hmin := g.Constraints.Max.X
		vmin := g.Constraints.Min.Y
		layout.Stack{Alignment: layout.W}.Layout(*g,
			layout.Expanded(func(gtx C) D {
				var ddd D
				rr := float32(g.Px(unit.Dp(0)))
				clip.Rect{
					Rect: f32.Rectangle{Max: f32.Point{
						X: float32(g.Constraints.Min.X),
						Y: float32(g.Constraints.Min.Y),
					}},
					NE: rr, NW: rr, SE: rr, SW: rr,
				}.Op(gtx.Ops).Add(gtx.Ops)
				fill(gtx, HexARGB(background))
				pointer.Rect(image.Rectangle{Max: dd.Size}).Add(gtx.Ops)
				return ddd
			}),
			layout.Stacked(func(gtx C) D {
				var ddd D
				g.Constraints.Min.X = hmin
				g.Constraints.Min.Y = vmin
				direction.Layout(gtx, func(gtx C) D {
					var dddd D
					layout.Flex{}.Layout(gtx,
						layout.Flexed(1, func(gtx C) D {
							var ddddd D
							layout.Inset{
								Top:    unit.Dp(padding),
								Right:  unit.Dp(padding),
								Bottom: unit.Dp(padding),
								Left:   unit.Dp(padding),
							}.Layout(gtx, elementContent)
							return ddddd
						}))
					return dddd
				})
				return ddd
			}),
		)
		return dd
	}
}
