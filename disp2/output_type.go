// Copyright 2016 Simon Eisenmann <simon@longsleep.org>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package disp2

// All display output types i could find. Some might make no sense depended on
// the board and the driver in the Kernel tree.
const (
	DISP_OUTPUT_TYPE_NONE = 0
	DISP_OUTPUT_TYPE_LCD  = 1
	DISP_OUTPUT_TYPE_TV   = 2
	DISP_OUTPUT_TYPE_HDMI = 4
	DISP_OUTPUT_TYPE_VGA  = 8
)
