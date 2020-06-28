package gelook

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel"
)

var (
	widgetButtonUp   = new(widget.Clickable)
	widgetButtonDown = new(widget.Clickable)
)

type ScrollBar struct {
	ColorBg      string
	BorderRadius [4]float32
	OperateValue interface{}
	slider       *ScrollBarSlider
	up           material.IconButtonStyle
	down         material.IconButtonStyle
	container    WingUIcontainer
}

type ScrollBarSlider struct {
	container WingUIcontainer
	Icon      material.IconButtonStyle
}

func (t *WingUItheme) ScrollBar(leftMargin int) *ScrollBar {
	slider := &ScrollBarSlider{
		container: t.WingUIcontainer(0, t.Colors["Primary"]),
		//Icon:      *t.Icons["Grab"],
	}
	slider.container.CornerRadius = 8
	scrollbar := &ScrollBar{
		ColorBg:      t.Colors["DarkGrayII"],
		BorderRadius: [4]float32{},
		slider:       slider,
		up:           material.IconButton(t.T, widgetButtonUp, t.Icons["Up"]),
		down:         material.IconButton(t.T, widgetButtonDown, t.Icons["Down"]),
		container:    t.WingUIcontainer(0, t.Colors["Light"]),
	}
	scrollbar.container.PaddingLeft = leftMargin
	return scrollbar
}

func (p *WingUIpanel) ScrollBarLayout(g layout.Context, panel *gel.Panel) layout.Dimensions {
	return p.ScrollBar.container.Layout(g, layout.Center, func(gtx C) D {
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				for widgetButtonUp.Clicked() {
					if panel.PanelContentLayout.Position.First > 0 {
						panel.PanelContentLayout.Position.First = panel.PanelContentLayout.Position.First - 1
						panel.PanelContentLayout.Position.Offset = 0
					}
				}
				p.ScrollBar.up.Inset = layout.Inset{
					Top:    unit.Dp(16),
					Right:  unit.Dp(16),
					Bottom: unit.Dp(16),
					Left:   unit.Dp(16),
				}
				p.ScrollBar.up.Size = unit.Dp(16)
				p.ScrollBar.up.Color = HexARGB("ffcfcfcf")
				return p.ScrollBar.up.Layout(gtx)
			}),
			layout.Flexed(1, p.bodyLayout(gtx, panel)),
			layout.Rigid(func(gtx C) D {
				for widgetButtonDown.Clicked() {
					if panel.PanelContentLayout.Position.BeforeEnd {
						panel.PanelContentLayout.Position.First = panel.PanelContentLayout.Position.First + 1
						panel.PanelContentLayout.Position.Offset = 0
					}
				}
				p.ScrollBar.down.Inset = layout.Inset{
					Top:    unit.Dp(16),
					Right:  unit.Dp(16),
					Bottom: unit.Dp(16),
					Left:   unit.Dp(16),
				}
				p.ScrollBar.down.Size = unit.Dp(16)
				p.ScrollBar.down.Color = HexARGB("ffcfcfcf")
				return p.ScrollBar.down.Layout(g)
			}),
		)
	})
}

func (p *WingUIpanel) bodyLayout(g layout.Context, panel *gel.Panel) func(gtx C) layout.Dimensions {
	return func(gtx C) D {
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(g,
			layout.Rigid(func(gtx C) D {
				//cs := g.Constraints
				//pointer.Rect(
				//	image.Rectangle{Max: image.Point{X: cs.Max.X, Y: cs.Max.Y}},
				//).Add(g.Ops)
				//pointer.InputOp{Key: panel.ScrollBar.Slider}.Add(g.Ops)
				//return layout.Center.Layout(*g, func(gtx C) D {
				//return layout.Inset{
				//	Top: unit.Dp(float32(panel.PanelContentLayout.Position.First) * panel.ScrollUnit),
				//}.Layout(gtx, func(gtx C) D {
				//	gtx.Constraints.Min.X = panel.ScrollBar.Size
				//	gtx.Constraints.Min.Y = panel.ScrollBar.Slider.CursorHeight
				//	if panel.ScrollBar.Slider.CursorHeight < panel.ScrollBar.Size {
				//		panel.ScrollBar.Slider.CursorHeight = panel.ScrollBar.Size
				//	}
				//	p.ScrollBar.slider.container.Layout(layout.W, func(gtx C) D {
				//		//p.ScrollBar.slider.Icon.Color = HexARGB("ffcfcfcf")
				//		//p.ScrollBar.slider.Icon.Layout(g, unit.Px(float32(panel.ScrollBar.Size)))
				//	})
				//})
				//panel.ScrollBar.Slider.Layout(g)
				//})
				var d D
				return d
			}),
		)
	}
}
