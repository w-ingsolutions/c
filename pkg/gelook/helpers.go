package gelook

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func DuoUIdrawRectangle(g layout.Context, w, h int, color string, borderRadius [4]float32, padding [4]float32) D {
	in := layout.Inset{
		Top:    unit.Dp(padding[0]),
		Right:  unit.Dp(padding[1]),
		Bottom: unit.Dp(padding[2]),
		Left:   unit.Dp(padding[3]),
	}
	return in.Layout(g, func(gtx C) D {
		var dd D
		square := f32.Rectangle{
			Max: f32.Point{
				X: float32(w),
				Y: float32(h),
			},
		}
		paint.ColorOp{Color: HexARGB(color)}.Add(g.Ops)
		clip.Rect{Rect: square,
			NE: borderRadius[0], NW: borderRadius[1], SE: borderRadius[2], SW: borderRadius[3]}.Op(g.Ops).Add(g.Ops) // HLdraw
		paint.PaintOp{Rect: square}.Add(g.Ops)
		dd = layout.Dimensions{Size: image.Point{X: w, Y: h}}
		return dd
	})
}

func HexARGB(s string) (c color.RGBA) {
	_, _ = fmt.Sscanf(s, "%02x%02x%02x%02x", &c.A, &c.R, &c.G, &c.B)
	return
}

func WingUIfill(gtx layout.Context, col string) {
	cs := gtx.Constraints
	d := image.Point{X: cs.Min.X, Y: cs.Min.Y}
	dr := f32.Rectangle{
		Max: f32.Point{X: float32(d.X), Y: float32(d.Y)},
	}
	paint.ColorOp{Color: HexARGB(col)}.Add(gtx.Ops)
	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	//gtx.Dimensions = layout.Dimensions{Size: d}
}

func (t *WingUItheme) WingUIline(g layout.Context, verticalPadding, horizontalPadding float32, size int, color string) layout.Dimensions {
	return layout.Inset{
		Top:    unit.Dp(verticalPadding),
		Right:  unit.Dp(horizontalPadding),
		Bottom: unit.Dp(verticalPadding),
		Left:   unit.Dp(horizontalPadding),
	}.Layout(g, func(gtx C) D {
		return DuoUIdrawRectangle(g, g.Constraints.Max.X, size, color, [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
	})
}

func toPointF(p image.Point) f32.Point {
	return f32.Point{X: float32(p.X), Y: float32(p.Y)}
}
