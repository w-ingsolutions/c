// SPDX-License-Identifier: Unlicense OR MIT

package gelook

import (
	"gioui.org/widget"
	"image/color"

	"github.com/w-ingsolutions/c/pkg/gelook/ico"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

func NewWingUIicons() (i map[string]*widget.Icon) {
	i = make(map[string]*widget.Icon)
	i["Checked"] = mustIcon(widget.NewIcon(icons.ToggleCheckBox))
	i["Unchecked"] = mustIcon(widget.NewIcon(icons.ToggleCheckBoxOutlineBlank))
	i["RadioChecked"] = mustIcon(widget.NewIcon(icons.ToggleRadioButtonChecked))
	i["RadioUnchecked"] = mustIcon(widget.NewIcon(icons.ToggleRadioButtonUnchecked))
	i["iconCancel"] = mustIcon(widget.NewIcon(icons.NavigationCancel))
	i["iconOK"] = mustIcon(widget.NewIcon(icons.NavigationCheck))
	i["iconClose"] = mustIcon(widget.NewIcon(icons.NavigationClose))
	i["foldIn"] = mustIcon(widget.NewIcon(icons.ContentRemove))
	i["minimize"] = mustIcon(widget.NewIcon(icons.NavigationExpandMore))
	i["zoom"] = mustIcon(widget.NewIcon(icons.NavigationExpandLess))
	i["logo"] = mustIcon(widget.NewIcon(ico.WingLogo))
	i["overviewIcon"] = mustIcon(widget.NewIcon(icons.ActionHome))
	i["sendIcon"] = mustIcon(widget.NewIcon(icons.ActionStarRate))
	i["receiveIcon"] = mustIcon(widget.NewIcon(icons.NavigationArrowDropDown))
	i["addressBookIcon"] = mustIcon(widget.NewIcon(icons.ActionBook))
	i["historyIcon"] = mustIcon(widget.NewIcon(icons.ActionHistory))
	i["closeIcon"] = mustIcon(widget.NewIcon(icons.NavigationClose))
	i["settingsIcon"] = mustIcon(widget.NewIcon(icons.ActionSettings))
	i["blocksIcon"] = mustIcon(widget.NewIcon(icons.ActionExplore))
	i["networkIcon"] = mustIcon(widget.NewIcon(icons.ActionFingerprint))
	i["traceIcon"] = mustIcon(widget.NewIcon(icons.ActionTrackChanges))
	// i["consoleIcon"] = mustIcon(widget.NewIcon(icons.ActionInput))
	i["helpIcon"] = mustIcon(widget.NewIcon(icons.NavigationArrowDropDown))
	i["counterPlusIcon"] = mustIcon(widget.NewIcon(icons.ImageExposurePlus1))
	i["counterMinusIcon"] = mustIcon(widget.NewIcon(icons.ImageExposureNeg1))
	i["CommunicationImportExport"] = mustIcon(widget.NewIcon(icons.CommunicationImportExport))
	i["NotificationNetworkCheck"] = mustIcon(widget.NewIcon(icons.NotificationNetworkCheck))
	i["NotificationSync"] = mustIcon(widget.NewIcon(icons.NotificationSync))
	i["NotificationSyncDisabled"] = mustIcon(widget.NewIcon(icons.NotificationSyncDisabled))
	i["NotificationSyncProblem"] = mustIcon(widget.NewIcon(icons.NotificationSyncProblem))
	i["NotificationVPNLock"] = mustIcon(widget.NewIcon(icons.NotificationVPNLock))
	i["network"] = mustIcon(widget.NewIcon(icons.NotificationWiFi))
	i["MapsLayers"] = mustIcon(widget.NewIcon(icons.MapsLayers))
	i["MapsLayersClear"] = mustIcon(widget.NewIcon(icons.MapsLayersClear))
	i["ImageTimer"] = mustIcon(widget.NewIcon(icons.ImageTimer))
	i["ImageRemoveRedEye"] = mustIcon(widget.NewIcon(icons.ImageRemoveRedEye))
	i["DeviceSignalCellular0Bar"] = mustIcon(widget.NewIcon(icons.DeviceSignalCellular0Bar))
	i["DeviceWidgets"] = mustIcon(widget.NewIcon(icons.DeviceWidgets))
	i["ActionTimeline"] = mustIcon(widget.NewIcon(icons.ActionTimeline))
	i["HardwareWatch"] = mustIcon(widget.NewIcon(icons.HardwareWatch))
	i["consoleIcon"] = mustIcon(widget.NewIcon(icons.HardwareKeyboardHide))
	i["DeviceSignalCellular0Bar"] = mustIcon(widget.NewIcon(icons.DeviceSignalCellular0Bar))
	i["HardwareWatch"] = mustIcon(widget.NewIcon(icons.HardwareWatch))
	i["EditorMonetizationOn"] = mustIcon(widget.NewIcon(icons.EditorMonetizationOn))
	i["Run"] = mustIcon(widget.NewIcon(icons.AVPlayArrow))
	i["Stop"] = mustIcon(widget.NewIcon(icons.AVStop))
	i["Pause"] = mustIcon(widget.NewIcon(icons.AVPause))
	i["Kill"] = mustIcon(widget.NewIcon(icons.NavigationCancel))
	i["Restart"] = mustIcon(widget.NewIcon(icons.NavigationRefresh))
	i["Grab"] = mustIcon(widget.NewIcon(icons.NavigationMenu))
	i["Up"] = mustIcon(widget.NewIcon(icons.NavigationArrowDropUp))
	i["Down"] = mustIcon(widget.NewIcon(icons.NavigationArrowDropDown))
	i["iconGrab"] = mustIcon(widget.NewIcon(icons.NavigationMenu))
	i["iconUp"] = mustIcon(widget.NewIcon(icons.NavigationArrowDropUp))
	i["iconDown"] = mustIcon(widget.NewIcon(icons.NavigationArrowDropDown))
	i["Copy"] = mustIcon(widget.NewIcon(icons.ContentContentCopy))
	i["Paste"] = mustIcon(widget.NewIcon(icons.ContentContentPaste))
	i["Sidebar"] = mustIcon(widget.NewIcon(icons.ActionChromeReaderMode))
	i["Filter"] = mustIcon(widget.NewIcon(icons.ContentFilterList))
	i["FilterAll"] = mustIcon(widget.NewIcon(icons.ActionDoneAll))
	i["FilterNone"] = mustIcon(widget.NewIcon(icons.ContentBlock))
	i["Build"] = mustIcon(widget.NewIcon(icons.ActionBuild))
	i["Folded"] = mustIcon(widget.NewIcon(icons.NavigationChevronRight))
	i["Unfolded"] = mustIcon(widget.NewIcon(icons.NavigationExpandMore))
	i["HideAll"] = mustIcon(widget.NewIcon(icons.NavigationUnfoldLess))
	i["ShowAll"] = mustIcon(widget.NewIcon(icons.NavigationUnfoldMore))
	i["HideItem"] = mustIcon(widget.NewIcon(icons.ActionVisibilityOff))
	i["ShowItem"] = mustIcon(widget.NewIcon(icons.ActionVisibility))
	i["TRC"] = mustIcon(widget.NewIcon(icons.ActionSearch))
	i["DBG"] = mustIcon(widget.NewIcon(icons.ActionBugReport))
	i["INF"] = mustIcon(widget.NewIcon(icons.ActionInfo))
	i["WRN"] = mustIcon(widget.NewIcon(icons.ActionHelp))
	i["CHK"] = mustIcon(widget.NewIcon(icons.AlertWarning))
	i["ERR"] = mustIcon(widget.NewIcon(icons.AlertError))
	i["FTL"] = mustIcon(widget.NewIcon(icons.ImageFlashOn))
	i["Delete"] = mustIcon(widget.NewIcon(icons.ActionDelete))
	i["Send"] = mustIcon(widget.NewIcon(icons.ContentSend))
	i["Screenshot"] = mustIcon(widget.NewIcon(icons.ContentSelectAll))
	i["ToggleOn"] = mustIcon(widget.NewIcon(icons.ToggleRadioButtonChecked))
	i["ToggleOff"] = mustIcon(widget.NewIcon(icons.ToggleRadioButtonUnchecked))
	return i
}

func mustIcon(ic *widget.Icon, err error) *widget.Icon {
	if err != nil {
		panic(err)
	}
	return ic
}

func rgb(c uint32) color.RGBA {
	return argb(0xff000000 | c)
}

func argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}

func fill(gtx layout.Context, col color.RGBA) layout.Dimensions {
	cs := gtx.Constraints
	d := cs.Min
	dr := f32.Rectangle{
		Max: f32.Point{X: float32(d.X), Y: float32(d.Y)},
	}
	paint.ColorOp{Color: col}.Add(gtx.Ops)
	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	return layout.Dimensions{Size: d}
}
