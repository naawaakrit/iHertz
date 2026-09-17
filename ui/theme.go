// Copyright (c) 2026 Naawaakrit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyTheme struct{}

func (m MyTheme) Color(name fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	if v == theme.VariantDark {
		switch name {
		case theme.ColorNameBackground:
			return color.NRGBA{13, 17, 23, 255}
		case theme.ColorNameForeground:
			return color.NRGBA{245, 247, 250, 255}
		case theme.ColorNameButton:
			return color.NRGBA{31, 41, 55, 255}
		case theme.ColorNamePressed:
			return color.NRGBA{0, 210, 235, 255}
		case theme.ColorNameHover:
			return color.NRGBA{255, 255, 255, 24}
		case theme.ColorNameDisabledButton:
			return color.NRGBA{55, 65, 81, 255}
		case theme.ColorNameDisabled:
			return color.NRGBA{107, 114, 128, 255}
		case theme.ColorNameFocus:
			return color.NRGBA{0, 210, 235, 90}
		case theme.ColorNamePrimary:
			return color.NRGBA{0, 210, 235, 255}
		case theme.ColorNameInputBackground:
			return color.NRGBA{24, 32, 43, 255}
		case theme.ColorNamePlaceHolder:
			return color.NRGBA{148, 163, 184, 255}
		case theme.ColorNameMenuBackground:
			return color.NRGBA{22, 27, 34, 255}
		case theme.ColorNameOverlayBackground:
			return color.NRGBA{13, 17, 23, 255}
		case theme.ColorNameShadow:
			return color.NRGBA{0, 210, 235, 90}
		case theme.ColorNameError:
			return color.NRGBA{248, 113, 113, 255}
		case theme.ColorNameSuccess:
			return color.NRGBA{74, 222, 128, 255}
		case theme.ColorNameWarning:
			return color.NRGBA{251, 191, 36, 255}
		}
	} else {
		switch name {
		case theme.ColorNameBackground:
			return color.NRGBA{248, 250, 252, 255}
		case theme.ColorNameForeground:
			return color.NRGBA{15, 23, 42, 255}
		case theme.ColorNameButton:
			return color.NRGBA{226, 232, 240, 255}
		case theme.ColorNamePressed:
			return color.NRGBA{0, 151, 173, 255}
		case theme.ColorNameHover:
			return color.NRGBA{15, 23, 42, 24}
		case theme.ColorNameDisabledButton:
			return color.NRGBA{203, 213, 225, 255}
		case theme.ColorNameDisabled:
			return color.NRGBA{100, 116, 139, 255}
		case theme.ColorNameFocus:
			return color.NRGBA{0, 151, 173, 90}
		case theme.ColorNamePrimary:
			return color.NRGBA{0, 151, 173, 255}
		case theme.ColorNameInputBackground:
			return color.NRGBA{255, 255, 255, 255}
		case theme.ColorNamePlaceHolder:
			return color.NRGBA{71, 85, 105, 255}
		case theme.ColorNameMenuBackground:
			return color.NRGBA{255, 255, 255, 255}
		case theme.ColorNameOverlayBackground:
			return color.NRGBA{241, 245, 249, 255}
		case theme.ColorNameShadow:
			return color.NRGBA{0, 151, 173, 70}
		case theme.ColorNameError:
			return color.NRGBA{220, 38, 38, 255}
		case theme.ColorNameSuccess:
			return color.NRGBA{22, 163, 74, 255}
		case theme.ColorNameWarning:
			return color.NRGBA{217, 119, 6, 255}
		}
	}
	return theme.DefaultTheme().Color(name, v)
}

// ต้องมีครบ
func (m MyTheme) Font(s fyne.TextStyle) fyne.Resource {
	return myFont
	//return theme.DefaultTheme().Font(s)
}
func (m MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}
func (m MyTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {

	// 📏 spacing / ระยะ
	case theme.SizeNamePadding: // → ระยะห่างทั่วไป (margin/padding ของ widget)
		return 4
	case theme.SizeNameSeparatorThickness: // → ความหนาเส้นคั่น
		return 1

	// 🖼️ ไอคอน / scrollbar
	case theme.SizeNameInlineIcon: // → ขนาด icon ในปุ่ม/ข้อความ /dialog
		return 19

	case theme.SizeNameScrollBar: // → ความกว้าง scrollbar ปกติ
		return 12
	case theme.SizeNameScrollBarSmall: // → scrollbar แบบเล็ก
		return 3

	// 🔤 ขนาดตัวอักษร
	case theme.SizeNameText: // → ข้อความปกติ
		return 14
	case theme.SizeNameHeadingText: // → หัวข้อใหญ่
		return 20
	case theme.SizeNameSubHeadingText: // → หัวข้อรอง
		return 16
	case theme.SizeNameCaptionText: // → ตัวเล็ก (caption/คำอธิบาย)
		return 12

	// 🧾 input
	case theme.SizeNameInputBorder: // → ความหนาขอบ input
		return 1
	}
	return theme.DefaultTheme().Size(name)
}
