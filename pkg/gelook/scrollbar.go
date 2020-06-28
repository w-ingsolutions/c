package gelook

import (
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel"
	"image"
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
	//up           IconButton
	//down         IconButton
	container DuoUIcontainer
}

type ScrollBarSlider struct {
	container DuoUIcontainer
	//Icon      DuoUIicon
}

func (t *WingUItheme) ScrollBar(leftMargin int) *ScrollBar {
	slider := &ScrollBarSlider{
		container: t.DuoUIcontainer(0, t.Colors["Primary"]),
		//Icon:      *t.Icons["Grab"],
	}
	slider.container.CornerRadius = 8
	scrollbar := &ScrollBar{
		ColorBg:      t.Colors["DarkGrayII"],
		BorderRadius: [4]float32{},
		slider:       slider,
		//up:           t.IconButton(t.Icons["Up"]),
		//down:         t.IconButton(t.Icons["Down"]),
		container: t.DuoUIcontainer(0, t.Colors["Light"]),
	}
	scrollbar.container.PaddingLeft = leftMargin
	return scrollbar
}

func (p *DuoUIpanel) ScrollBarLayout(g *layout.Context, panel *gel.Panel) {
	p.ScrollBar.container.Layout(g, layout.Center, func(gtx C) D {
		var dd D
		layout.Flex{
			Axis: layout.Vertical,
		}.Layout(*g,
			layout.Rigid(func(gtx C) D {
				var ddd D
				for widgetButtonUp.Clicked() {
					if panel.PanelContentLayout.Position.First > 0 {
						panel.PanelContentLayout.Position.First = panel.PanelContentLayout.Position.First - 1
						panel.PanelContentLayout.Position.Offset = 0
					}
				}
				//p.ScrollBar.up.Padding = unit.Dp(0)
				//p.ScrollBar.up.Size = unit.Dp(16)
				//p.ScrollBar.up.Color = HexARGB("ffcfcfcf")
				//p.ScrollBar.up.Layout(gtx, widgetButtonUp)
				return ddd
			}),
			layout.Flexed(1, func(gtx C) D {
				var ddd D
				p.bodyLayout(g, panel)
				return ddd
			}),
			layout.Rigid(func(gtx C) D {
				var ddd D
				for widgetButtonDown.Clicked() {
					if panel.PanelContentLayout.Position.BeforeEnd {
						panel.PanelContentLayout.Position.First = panel.PanelContentLayout.Position.First + 1
						panel.PanelContentLayout.Position.Offset = 0
					}
				}
				//p.ScrollBar.down.Padding = unit.Dp(0)
				//p.ScrollBar.down.Size = unit.Dp(16)
				//p.ScrollBar.down.Color = HexARGB("ffcfcfcf")
				//p.ScrollBar.down.Layout(g, widgetButtonDown)
				return ddd
			}),
		)
		return dd
	})
}

func (p *DuoUIpanel) bodyLayout(g *layout.Context, panel *gel.Panel) {
	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(*g,
		layout.Rigid(func(gtx C) D {
			var dd D
			cs := g.Constraints
			pointer.Rect(
				image.Rectangle{Max: image.Point{X: cs.Max.X, Y: cs.Max.Y}},
			).Add(g.Ops)
			//pointer.InputOp{Key: panel.ScrollBar.Slider}.Add(g.Ops)
			layout.Center.Layout(*g, func(gtx C) D {
				var ddd D
				layout.Inset{
					Top: unit.Dp(float32(panel.PanelContentLayout.Position.First) * panel.ScrollUnit),
				}.Layout(gtx, func(gtx C) D {
					var dddd D
					gtx.Constraints.Min.X = panel.ScrollBar.Size
					gtx.Constraints.Min.Y = panel.ScrollBar.Slider.CursorHeight
					if panel.ScrollBar.Slider.CursorHeight < panel.ScrollBar.Size {
						panel.ScrollBar.Slider.CursorHeight = panel.ScrollBar.Size
					}
					p.ScrollBar.slider.container.Layout(g, layout.W, func(gtx C) D {
						var ddddd D
						//p.ScrollBar.slider.Icon.Color = HexARGB("ffcfcfcf")
						//p.ScrollBar.slider.Icon.Layout(g, unit.Px(float32(panel.ScrollBar.Size)))
						return ddddd
					})
					return dddd
				})
				panel.ScrollBar.Slider.Layout(g)
				return ddd
			})
			return dd
		}),
	)
}
