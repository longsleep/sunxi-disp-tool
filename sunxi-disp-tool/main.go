// Copyright 2016 Simon Eisenmann <simon@longsleep.org>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	This implementation is based on information from the A64 BSP 1.2 Linux
	kernel source.

	The original disp2 framebuffer driver is fucked up and needs a patch like
	https://github.com/jernejsk/OpenELEC-OPi2/blob/openelec-7.0/projects/H3/patches/linux/linux-79-fbdev-fixes.patch
	to properly display after resolution change. Thanks to Jernej Škrabec for
	pointing me to it.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/longsleep/sunxi-disp-tool/disp2"
)

const (
	DispDev = "/dev/disp"
)

type Args [4]uint64

type Disp struct {
	f *os.File
}

func NewDisp(fn string) (*Disp, error) {
	f, err := os.OpenFile(DispDev, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	return &Disp{
		f: f,
	}, nil
}

func (disp *Disp) Close() {
	disp.f.Close()
}

func (disp *Disp) ioctl(cmd uint32, args *Args) (uintptr, uintptr, syscall.Errno) {
	return syscall.Syscall(syscall.SYS_IOCTL, disp.f.Fd(), uintptr(cmd), uintptr(unsafe.Pointer(args)))
}

func (disp *Disp) Switch(screen int, outputType uint64, outputMode uint64) error {
	r1, r2, errno := disp.ioctl(disp2.DISP_DEVICE_SWITCH, &Args{uint64(screen), outputType, outputMode})
	if errno != 0 {
		fmt.Fprintf(os.Stderr, "error: switch %v %v %v\n", errno, r1, r2)
		return errno
	}

	return nil
}

func main() {
	disp, err := NewDisp(DispDev)
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
		os.Exit(2)
	}
}
