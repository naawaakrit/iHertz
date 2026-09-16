// Copyright (c) 2026 Naawaakrit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package main

import (
	"fmt"
	"os"

	cpuinfo "ihertz/cpu"
	ui "ihertz/ui" //ชื่อไฟล์ "ชื่อ MOD / ชื่อ โฟรเดอร์" **แล้วดึงมาใช้
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--apply-cpu" {
		if err := cpuinfo.ApplyCPUSettingsFromArgs(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	ui.CreateWindow()
}
