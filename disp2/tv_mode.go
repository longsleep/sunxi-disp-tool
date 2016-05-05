// Copyright 2016 Simon Eisenmann <simon@longsleep.org>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package disp2

// All display output modes i could find. Actual support for those depend on
// the chip and the driver in the Kernel tree in use.
const (
	DISP_TV_MOD_480I             = 0
	DISP_TV_MOD_576I             = 1
	DISP_TV_MOD_480P             = 2
	DISP_TV_MOD_576P             = 3
	DISP_TV_MOD_720P_50HZ        = 4
	DISP_TV_MOD_720P_60HZ        = 5
	DISP_TV_MOD_1080I_50HZ       = 6
	DISP_TV_MOD_1080I_60HZ       = 7
	DISP_TV_MOD_1080P_24HZ       = 8
	DISP_TV_MOD_1080P_50HZ       = 9
	DISP_TV_MOD_1080P_60HZ       = 0xa
	DISP_TV_MOD_1080P_24HZ_3D_FP = 0x17
	DISP_TV_MOD_720P_50HZ_3D_FP  = 0x18
	DISP_TV_MOD_720P_60HZ_3D_FP  = 0x19
	DISP_TV_MOD_1080P_25HZ       = 0x1a
	DISP_TV_MOD_1080P_30HZ       = 0x1b
	DISP_TV_MOD_PAL              = 0xb
	DISP_TV_MOD_PAL_SVIDEO       = 0xc
	DISP_TV_MOD_NTSC             = 0xe
	DISP_TV_MOD_NTSC_SVIDEO      = 0xf
	DISP_TV_MOD_PAL_M            = 0x11
	DISP_TV_MOD_PAL_M_SVIDEO     = 0x12
	DISP_TV_MOD_PAL_NC           = 0x14
	DISP_TV_MOD_PAL_NC_SVIDEO    = 0x15
	DISP_TV_MOD_3840_2160P_30HZ  = 0x1c
	DISP_TV_MOD_3840_2160P_25HZ  = 0x1d
	DISP_TV_MOD_3840_2160P_24HZ  = 0x1e
	DISP_TV_MODE_NUM             = 0x1f

	DISP_TV_MODE_UNKOWN = 0xffff
)

func GetTVModFromString(args ...string) int {
	for _, arg := range args {
		switch arg {
		case "EDID":
			// Not implemented yet.
		case "720p50":
			return DISP_TV_MOD_720P_50HZ
		case "720p":
			fallthrough
		case "720p60":
			return DISP_TV_MOD_720P_60HZ
		case "1080i50":
			return DISP_TV_MOD_1080I_50HZ
		case "1080i":
			fallthrough
		case "1080i60":
			return DISP_TV_MOD_1080I_60HZ
		case "1080p24":
			return DISP_TV_MOD_1080P_24HZ
		case "1080p50":
			return DISP_TV_MOD_1080P_50HZ
		case "1080p":
			fallthrough
		case "1080p60":
			return DISP_TV_MOD_1080P_60HZ
		case "2160p30":
			return DISP_TV_MOD_3840_2160P_30HZ
		case "2160p25":
			return DISP_TV_MOD_3840_2160P_25HZ
		case "2160p24":
			return DISP_TV_MOD_3840_2160P_24HZ
		}
	}

	return DISP_TV_MODE_UNKOWN
}
