// Copyright 2016 Simon Eisenmann <simon@longsleep.org>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fb

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

type Set struct {
	f *os.File
}

func NewSet(fn string) (*Set, error) {
	f, err := os.OpenFile(fn, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	return &Set{
		f: f,
	}, nil
}

func (set *Set) Close() {
	set.f.Close()
}

func (set *Set) GetVarScreenInfo() (*VarScreenInfo, error) {
	varScreenInfo := &VarScreenInfo{}
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, set.f.Fd(), FBIOGET_VSCREENINFO, uintptr(unsafe.Pointer(varScreenInfo)))
	if errno != 0 {
		return nil, fmt.Errorf("ioctl get_vscreeninfo failed %v", errno)
	}

	return varScreenInfo, nil
}

func (set *Set) SetVarScreenInfo(varScreenInfo *VarScreenInfo) error {
	varScreenInfo.activate = FB_ACTIVATE_ALL

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, set.f.Fd(), FBIOPUT_VSCREENINFO, uintptr(unsafe.Pointer(varScreenInfo)))
	if errno != 0 {
		return fmt.Errorf("ioctl put_vscreeninfo failed %v", errno)
	}

	return nil
}
