// SPDX-License-Identifier: Unlicense OR MIT

package gelook

import (
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image/color"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

type WingUItheme struct {
	T        *material.Theme
	Shaper   text.Shaper
	TextSize unit.Value
	Color    struct {
		Primary color.RGBA
		Text    color.RGBA
		Hint    color.RGBA
		InvText color.RGBA
	}
	Colors map[string]string
	Fonts  map[string]text.Typeface
	Icons  map[string]*widget.Icon

	scrollBarSize int
}

func register(fnt text.Font, ttf []byte) {
	//face, err := opentype.Parse(ttf)
	//if err != nil {
	//	panic(fmt.Sprintf("failed to parse font: %v", err))
	//}
	//fnt.Typeface = "Go"
	//font.Register(fnt, face)
}

func NewWingUItheme() *WingUItheme {
	t := &WingUItheme{
		T: material.NewTheme(gofont.Collection()),
		//Shaper: font.Default(),
	}
	t.Colors = NewWingUIcolors()
	t.TextSize = unit.Sp(16)
	t.Icons = NewWingUIicons()
	return t
}

func (t *WingUItheme) ChangeLightDark() {
	light := t.Colors["Light"]
	dark := t.Colors["Dark"]
	lightGray := t.Colors["LightGrayIII"]
	darkGray := t.Colors["DarkGrayII"]
	t.Colors["Light"] = dark
	t.Colors["Dark"] = light
	t.Colors["LightGrayIII"] = darkGray
	t.Colors["DarkGrayII"] = lightGray
}

// mulAlpha scales all color components by alpha/255.
func mulAlpha(c color.RGBA, alpha uint8) color.RGBA {
	a := uint16(alpha)
	return color.RGBA{
		A: uint8(uint16(c.A) * a / 255),
		R: uint8(uint16(c.R) * a / 255),
		G: uint8(uint16(c.G) * a / 255),
		B: uint8(uint16(c.B) * a / 255),
	}
}
