// Copyright (c) 2022, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parseval

import (
	"strconv"

	"github.com/go-pogo/errors"
)

// Bool tries to parse Value as a bool with strconv.ParseBool.
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Any other value returns an error.
func (v Value) Bool() (bool, error) {
	x, err := strconv.ParseBool(string(v))
	return x, errors.WithKind(err, errKind(err))
}

// BoolVar sets the value p points to using Bool.
func (v Value) BoolVar(p *bool) (err error) {
	*p, err = v.Bool()
	return
}
