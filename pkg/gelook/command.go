// SPDX-License-Identifier: Unlicense OR MIT

package gelook

//type Command struct {
//	Com      interface{}
//	ComID    string
//	Category string
//	Out      func()
//	Time     time.Time
//}
//
//func (t *DuoUItheme) Command(name string) *Command {
//	return &Command{
//		ComID: name,
//	}
//}
//
//func (p Command) Layout(g *layout.Context, f func()) {
//	layout.Flex{}.Layout(*g,
//		layout.Flexed(1, func(gtx C) D {
//			in := layout.UniformInset(unit.Dp(0))
//			in.Layout(*g, func(gtx C) D {
//				cs := g.Constraints
//				DuoUIdrawRectangle(g, cs.Max.X, cs.Max.Y, "ffacacac", [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
//				layout.UniformInset(unit.Dp(0)).Layout(*g, func(gtx C) D {
//					cs := g.Constraints
//					DuoUIdrawRectangle(g, cs.Max.X, cs.Max.Y, "ffcfcfcf", [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
//					f()
//				})
//			})
//		}),
//	)
//}
