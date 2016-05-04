// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package alidayu

import (
	"testing"

	"github.com/issue9/assert"
)

func TestSplitCode(t *testing.T) {
	a := assert.New(t)

	a.Equal(SplitCode("1111"), "1,1,1,1")
	a.Equal(SplitCode("1234"), "1,2,3,4")
}
