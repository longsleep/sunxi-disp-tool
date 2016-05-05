// Copyright 2016 Simon Eisenmann <simon@longsleep.org>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kernel

import (
	"io/ioutil"
	"strings"
)

const (
	ProcCmdLine = "/proc/cmdline"
)

func GetCmdlineParamValue(name string) (string, bool) {
	if raw, err := ioutil.ReadFile(ProcCmdLine); err == nil {
		params := strings.Split(string(raw), " ")

		for _, param := range params {
			if strings.HasPrefix(param, name) {
				value := strings.SplitN(param, "=", 2)[1]
				return value, true
			}
		}
	}

	return "", false
}
