# Sunxi disp tool

This utility allows to control a sunxi disp2 device from user space. The code
is organized in Go modules, so it can easily be integrated.

## Usage

```
Usage: sunxi-disp-tool [<global-options>] <command> [<args>]
Global options:
  -screen int
          Screen ID (default 0)
Available commands are:
  switch   Switch HDMI output mode
  init     Initialize HDMI output mode from Kernel args

```

### Examples

```bash
sunxi-disp-tool switch -mode 0xa
sunxi-disp-tool switch -mode 0x5
sunxi-disp-tool switch -name 720p
sunxi-disp-tool switch -name 1080p
sunxi-disp-tool switch -name 1020x720p60
```


## Available HDMI output names

The following names are supported by this tool, and are mapped to there
corresponding output mode number values. These names can either be passed with
the `-name` parameter of the `switch` command or via Kernel argument using the
`init` command.

```
720p50
720p60
720p
1080i50
1080i60
1080p24
1080p50
1080p60
1080p
2160p24
2160p25
2160p30
2160p
```

The values without a frequency all map to the highest supported frequency for
the selected vertical solution. The parser also supports formats like
`1280x720p` and splits every value by `x` to be consistent with other existing
Sunxi documentation.


## Available HDMI output modes (-mode parameter of `switch` command)

The following modes are supported by this tool. If the modes actually work
depends on your HDMI device, the Sunxi board and the Kernel.

```
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
```

--
Simon Eisenmann - mailto:simon@longsleep.org
