package gelook

import (
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type DuoUIcounter struct {
	increase material.ButtonStyle
	decrease material.ButtonStyle
	reset    material.ButtonStyle
	//input        DuoUIeditor
	pageFunction func()
	Font         text.Font
	TextSize     unit.Value
	TxColor      string
	BgColor      string
	shaper       text.Shaper
}

func (t *WingUItheme) DuoUIcounter(pageFunction func()) DuoUIcounter {
	return DuoUIcounter{
		// ToDo Replace theme's buttons with counter exclusive buttons, set icons for increase/decrease
		//increase: t.DuoUIbutton("", "", "", t.Colors["Light"], "", t.Colors["Dark"], "counterPlusIcon", t.Colors["Primary"], 0, 24, 32, 32, 0, 0, 0, 0),
		//decrease: t.DuoUIbutton("", "", "", t.Colors["Light"], "", t.Colors["Dark"], "counterMinusIcon", t.Colors["Primary"], 0, 24, 32, 32, 0, 0, 0, 0),
		// reset:        t.DuoUIbutton(t.Font.Secondary, "RESET", t.Colors["Primary"], t.Colors["Light"], t.Colors["Light"], t.Colors["Primary"], "", "", 12, 0, 0, 48, 48, 0),
		//input:        t.DuoUIeditor("", "DocText", "DocBg", 5),
		pageFunction: pageFunction,
		Font: text.Font{
			Typeface: t.Fonts["Primary"],
		},
		TxColor:  t.Colors["Light"],
		BgColor:  t.Colors["Dark"],
		TextSize: unit.Dp(float32(18)),
		shaper:   t.Shaper,
	}
}

//
//func (c DuoUIcounter) Layout(g *layout.Context, cc *gel.DuoUIcounter, label, value string) {
//	cc.CounterInput.SetText(value)
//	hmin := g.Constraints.Min.X
//	vmin := g.Constraints.Min.Y
//	// txColor := c.TxColor
//	bgColor := c.BgColor
//	layout.Stack{Alignment: layout.Center}.Layout(g,
//		layout.Expanded(func() {
//			rr := float32(g.Px(unit.Dp(0)))
//			clip.Rect{
//				Rect: f32.Rectangle{Max: f32.Point{
//					X: float32(g.Constraints.Min.X),
//					Y: float32(g.Constraints.Min.Y),
//				}},
//				NE: rr, NW: rr, SE: rr, SW: rr,
//			}.Op(g.Ops).Add(g.Ops)
//			fill(g, HexARGB(bgColor))
//		}),
//		layout.Stacked(func() {
//			g.Constraints.Min.X = hmin
//			g.Constraints.Min.Y = vmin
//			layout.Center.Layout(g, func() {
//				layout.Flex{
//					Spacing:   layout.SpaceAround,
//					Alignment: layout.Middle,
//				}.Layout(gtx,
//					layout.Rigid(func() {
//						for cc.CounterDecrease.Clicked(gtx) {
//							cc.Decrease()
//							c.pageFunction()
//						}
//						c.decrease.IconLayout(gtx, cc.CounterDecrease)
//					}),
//					layout.Rigid(func() {
//						layout.Center.Layout(gtx, func() {
//							layout.Inset{
//								Top:    unit.Dp(0),
//								Right:  unit.Dp(16),
//								Bottom: unit.Dp(0),
//								Left:   unit.Dp(16),
//							}.Layout(gtx, func() {
//								layout.Flex{
//									Axis:      layout.Vertical,
//									Spacing:   layout.SpaceAround,
//									Alignment: layout.Middle,
//								}.Layout(gtx,
//									layout.Rigid(func() {
//										paint.ColorOp{Color: HexARGB(c.TxColor)}.Add(gtx.Ops)
//										gel.Label{
//											Alignment: text.Middle,
//										}.Layout(gtx, c.shaper, c.Font, unit.Dp(8), label)
//									}),
//									layout.Rigid(func() {
//										c.input.Font.Typeface = c.Font.Typeface
//										c.input.Color = HexARGB(c.TxColor)
//										c.input.Layout(gtx, cc.CounterInput)
//										for _, e := range cc.CounterInput.Events(gtx) {
//											switch e.(type) {
//											case gel.ChangeEvent:
//												if i, err := strconv.Atoi(cc.CounterInput.Text()); err == nil {
//													cc.Value = i
//												}
//											}
//										}
//										// paint.ColorOp{Color: HexARGB(c.TxColor)}.Add(gtx.Ops)
//										// gel.Label{
//										//	Alignment: text.Middle,
//										// }.Layout(gtx, c.shaper, c.Font, unit.Dp(12), value)
//									}))
//							})
//						})
//					}),
//					// layout.Flexed(0.2, func() {
//					//	//for cc.CounterReset.Clicked(gtx) {
//					//	//	cc.Reset()
//					//	//	c.pageFunction()
//					//	//}
//					//	//c.reset.Layout(gtx, cc.CounterReset)
//					// }),
//					layout.Rigid(func() {
//						for cc.CounterIncrease.Clicked(gtx) {
//							cc.Increase()
//							c.pageFunction()
//						}
//						c.increase.IconLayout(gtx, cc.CounterIncrease)
//					}))
//			})
//		}),
//	)
//
//}
