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

package disp2

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/longsleep/sunxi-disp-tool/fb"
)

const (
	DispDev     = "/dev/disp"
	FbDevPrefix = "/dev/fb"
)

type ioArgs [4]uint64

type Disp2 struct {
	f *os.File
}

func New() (*Disp2, error) {
	f, err := os.OpenFile(DispDev, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	return &Disp2{
		f: f,
	}, nil
}

func (disp *Disp2) Close() {
	disp.f.Close()
}

func (disp *Disp2) ioctl(cmd uint32, args *ioArgs) (uintptr, uintptr, syscall.Errno) {
	return syscall.Syscall(syscall.SYS_IOCTL, disp.f.Fd(), uintptr(cmd), uintptr(unsafe.Pointer(args)))
}

func (disp *Disp2) Switch(screen int, outputType uint64, outputMode uint64) error {
	_, _, errno := disp.ioctl(DISP_DEVICE_SWITCH, &ioArgs{uint64(screen), outputType, outputMode})
	if errno != 0 {
		return fmt.Errorf("ioctl switch failed %v", errno)
	}

	height, _ := disp.GetScnHeight(screen)
	width, _ := disp.GetScnWidth(screen)

	if width > 0 && height > 0 {
		// Set fb size as well.
		err := disp.fbSet(screen, width, height)
		if err != nil {
			return fmt.Errorf("fbset failed %v", err)
		}
	} else {
		return fmt.Errorf("width or height is zero")
	}

	return nil
}

func (disp *Disp2) GetScnHeight(screen int) (uint32, error) {
	r1, _, errno := disp.ioctl(DISP_GET_SCN_HEIGHT, &ioArgs{uint64(screen)})
	if errno != 0 {
		return 0, fmt.Errorf("ioctl get_scn_height failed %v", errno)
	}

	return uint32(r1), nil
}

func (disp *Disp2) GetScnWidth(screen int) (uint32, error) {
	r1, _, errno := disp.ioctl(DISP_GET_SCN_WIDTH, &ioArgs{uint64(screen)})
	if errno != 0 {
		return 0, fmt.Errorf("ioctl get_scn_width failed %v", errno)
	}

	return uint32(r1), nil
}

func (disp *Disp2) fbSet(screen int, width, height uint32) error {
	set, err := fb.NewSet(fmt.Sprintf("%s%d", FbDevPrefix, screen))
	if err != nil {
		return err
	}
	defer set.Close()

	info, err := set.GetVarScreenInfo()
	if err != nil {
		return err
	}

	info.SetXRes(width, width)
	info.SetYRes(height, height*2)

	return set.SetVarScreenInfo(info)
}
