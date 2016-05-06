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
	fmt.Println("Usage: sunxi-disp-tool [<global-options>] <command> [<args>]")
	fmt.Println("Global options:")
	fmt.Println("  -screen int")
	fmt.Println("          Screen ID (default 0)")
	fmt.Println("Available commands are:")
	fmt.Println("  switch   Switch HDMI output mode")
	fmt.Println("  init     Initialize HDMI output mode from Kernel args")
}

func main() {
	globalOptions := flag.NewFlagSet("global", flag.ContinueOnError)
	screenID := globalOptions.Int("screen", 0, "Screen ID")

	switchCommand := flag.NewFlagSet("switch", flag.ExitOnError)
	outputMode := switchCommand.Int("mode", disp2.DISP_TV_MOD_1080P_60HZ, "Set HDMI output mode by number")
	outputModeName := switchCommand.String("name", "", "Set HDMI output mode by name")

	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	initKernelArg := initCommand.String("kernelarg", "disp.screen0_output_mode", "Set HDMI output mode from Kernel arg")

	globalOptions.Parse(os.Args[1:])
	args := globalOptions.Args()

	if len(args) == 0 {
		usage()
		return
	}

	switch args[0] {
	case "switch":
		switchCommand.Parse(args[1:])
	case "init":
		initCommand.Parse(args[1:])
	case "-h":
		usage()
		return
	default:
		fmt.Printf("%q is not valid command.\n", args[0])
		os.Exit(1)
	}

	disp, err := disp2.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)

	}
	defer disp.Close()

	if switchCommand.Parsed() {
		if *outputModeName != "" {
			mode := disp2.GetTVModFromString(*outputModeName)
			if mode != disp2.DISP_TV_MODE_UNKOWN {
				*outputMode = mode
			}
		}

		err = disp.Switch(*screenID, disp2.DISP_OUTPUT_TYPE_HDMI, uint64(*outputMode))
	} else if initCommand.Parsed() {
		mode := disp2.DISP_TV_MODE_UNKOWN
		if *initKernelArg != "" {
			boot, ok := kernel.GetCmdlineParamValue(*initKernelArg)
			if ok {
				resolutions := strings.Split(boot, ":")
				mode = disp2.GetTVModFromString(resolutions...)
				if mode == disp2.DISP_TV_MODE_UNKOWN {
					// No or invalid mode set via kerne parameter. Do nothing.
					return
				}
			}
		}

		err = disp.Switch(*screenID, disp2.DISP_OUTPUT_TYPE_HDMI, uint64(mode))
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(3)
	}
}
