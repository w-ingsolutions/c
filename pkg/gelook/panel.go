package gelook

import (
	"gioui.org/layout"
	"github.com/gioapp/gel"
)

type WingUIpanel struct {
	PanelObject interface{}
	ScrollBar   *ScrollBar
	container   WingUIcontainer
}

func (t *WingUItheme) WingUIpanel() WingUIpanel {
	return WingUIpanel{
		container: t.WingUIcontainer(0, t.Colors["Light"]),
	}
}
func (p *WingUIpanel) panelLayout(g *layout.Context, panel *gel.Panel, row func(i int, in interface{})) func(gtx C) D {
	return func(gtx C) D {
		var dd D
		visibleObjectsNumber := 0
		panel.PanelContentLayout.Layout(*g, panel.PanelObjectsNumber, func(gtx C, i int) D {
			var ddd D
			row(i, p.PanelObject)
			visibleObjectsNumber = visibleObjectsNumber + 1
			panel.VisibleObjectsNumber = visibleObjectsNumber
			return ddd
		})
		return dd
	}
}

func (p *WingUIpanel) Layout(g layout.Context, panel *gel.Panel, row func(i int, in interface{})) func(gtx C) D {
	return func(gtx C) D {
		return p.container.Layout(&g, layout.NW, func(gtx C) D {
			return layout.Flex{
				Axis:    layout.Horizontal,
				Spacing: layout.SpaceBetween,
			}.Layout(g,
				layout.Flexed(1, p.panelLayout(&g, panel, row)),
				layout.Rigid(func(gtx C) D {
					var d D
					if panel.PanelObjectsNumber > panel.VisibleObjectsNumber {
						//return	p.ScrollBarLayout(g, panel)
					}
					return d
				}),
			)
			//fmt.Println("scrollUnit:", panel.ScrollUnit)
			//fmt.Println("ScrollBar.Slider.Height:", panel.ScrollBar.Slider.Height)
			//fmt.Println("PanelObjectsNumber:", panel.PanelObjectsNumber)

			//panel.Layout(&g)
		})
	}
}
