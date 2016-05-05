// Copyright 2016 Simon Eisenmann <simon@longsleep.org>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/longsleep/sunxi-disp-tool/disp2"
)

func usage() {
	fmt.Println("Usage: sunxi-disp-tool <command> [<args>]")
	fmt.Println("Available commands are:")
	fmt.Println("  switch   Switch HDMI output mode")
}

func main() {
	switchCommand := flag.NewFlagSet("switch", flag.ExitOnError)
	outputMode := switchCommand.Int("mode", disp2.DISP_TV_MOD_1080P_60HZ, "HDMI output mode")

	if len(os.Args) == 1 {
		usage()
		return
	}

	switch os.Args[1] {
	case "switch":
		switchCommand.Parse(os.Args[2:])
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
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(3)
	}
}
