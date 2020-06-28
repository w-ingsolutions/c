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
func (p *WingUIpanel) panelLayout(g *layout.Context, panel *gel.Panel, row func(i int, in interface{}) D) func(gtx C) D {
	return func(gtx C) D {
		visibleObjectsNumber := 0
		return panel.PanelContentLayout.Layout(*g, panel.PanelObjectsNumber, func(gtx C, i int) D {
			visibleObjectsNumber = visibleObjectsNumber + 1
			panel.VisibleObjectsNumber = visibleObjectsNumber
			return row(i, p.PanelObject)
		})
	}
}

func (p *WingUIpanel) Layout(g layout.Context, panel *gel.Panel, row func(i int, in interface{}) D) layout.Dimensions {
	//return p.container.Layout(g, layout.NW)
	panel.Layout(&g)
	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceBetween,
	}.Layout(g,
		layout.Flexed(1, p.panelLayout(&g, panel, row)),
		layout.Rigid(func(gtx C) D {
			var s layout.Dimensions
			if panel.PanelObjectsNumber > panel.VisibleObjectsNumber {
				s = p.ScrollBarLayout(g, panel)
			}
			return s
		}),
	)
	//fmt.Println("scrollUnit:", panel.ScrollUnit)
	//fmt.Println("ScrollBar.Slider.Height:", panel.ScrollBar.Slider.Height)
	//fmt.Println("PanelObjectsNumber:", panel.PanelObjectsNumber)

}
