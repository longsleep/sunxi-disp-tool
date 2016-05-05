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

func main() {
	disp, err := disp2.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)

	}
	defer disp.Close()

	// Parse commandline parameters.
	outputMode := flag.Int("mode", disp2.DISP_TV_MOD_1080P_60HZ, "HDMI output mode")
	flag.Parse()

	// For now only support switch command on first display as HDMI.
	err = disp.Switch(0, disp2.DISP_OUTPUT_TYPE_HDMI, uint64(*outputMode))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}
}
