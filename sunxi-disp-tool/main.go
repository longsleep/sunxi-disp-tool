// Copyright 2016 Simon Eisenmann <simon@longsleep.org>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/longsleep/sunxi-disp-tool/disp2"
	"github.com/longsleep/sunxi-disp-tool/kernel"
)

func usage() {
	fmt.Println("Usage: sunxi-disp-tool <command> [<args>]")
	fmt.Println("Available commands are:")
	fmt.Println("  switch   Switch HDMI output mode")
	fmt.Println("  init     Initialize HDMI output mode from Kernel args")
}

func main() {
	switchCommand := flag.NewFlagSet("switch", flag.ExitOnError)
	outputMode := switchCommand.Int("mode", disp2.DISP_TV_MOD_1080P_60HZ, "HDMI output mode")

	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	initKernelArg := initCommand.String("kernelarg", "disp.screen0_output_mode", "HDMI output Kernel arg name")
	initDefaultOutputMode := initCommand.Int("mode", disp2.DISP_TV_MOD_1080P_60HZ, "HDMI output mode default")

	if len(os.Args) == 1 {
		usage()
		return
	}

	switch os.Args[1] {
	case "switch":
		switchCommand.Parse(os.Args[2:])
	case "init":
		initCommand.Parse(os.Args[2:])
	case "-h":
		usage()
		return
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(1)
	}

	disp, err := disp2.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)

	}
	defer disp.Close()

	if switchCommand.Parsed() {
		err = disp.Switch(0, disp2.DISP_OUTPUT_TYPE_HDMI, uint64(*outputMode))
	} else if initCommand.Parsed() {
		if *initKernelArg != "" {
			boot, ok := kernel.GetCmdlineParamValue(*initKernelArg)
			if ok {
				resolutions := strings.Split(boot, ":")
				mode := disp2.GetTVModFromString(resolutions...)
				if mode != disp2.DISP_TV_MODE_UNKOWN {
					*initDefaultOutputMode = mode
				}
			}
		}

		err = disp.Switch(0, disp2.DISP_OUTPUT_TYPE_HDMI, uint64(*initDefaultOutputMode))
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(3)
	}
}
