// Copyright 2016 Simon Eisenmann <simon@longsleep.org>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fb

const (
	FB_ACTIVATE_ALL = 64

	FBIOGET_VSCREENINFO = 0x4600
	FBIOPUT_VSCREENINFO = 0x4601
)

type BitField struct {
	offset    uint32
	length    uint32
	msb_right uint32
}

type VarScreenInfo struct {
	xres         uint32
	yres         uint32
	xres_virtual uint32
	yres_virtual uint32
	xoffset      uint32
	yoffset      uint32

	bits_per_pixel uint32
	grayscale      uint32

	red    BitField
	green  BitField
	blue   BitField
	transp BitField

	nonstd uint32

	activate uint32

	height uint32
	width  uint32

	accel_flags uint32

	pixclock     uint32
	left_margin  uint32
	right_margin uint32
	upper_margin uint32
	lower_margin uint32
	hsync_len    uint32
	vsync_len    uint32
	sync         uint32
	vmode        uint32
	reserved     [6]uint32
}

func (info *VarScreenInfo) SetXRes(xres, xres_virtual uint32) {
	info.xres = xres
	info.xres_virtual = xres_virtual
}

func (info *VarScreenInfo) SetYRes(yres, yres_virtual uint32) {
	info.yres = yres
	info.yres_virtual = yres_virtual
}
