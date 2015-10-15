// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package stacktrace

import (
	"fmt"
	"testing"
)

func TestStackTrace(t *testing.T) {
	fmt.Println("TestStackTrace()")
	s := StackTrace(0)
	if s == "" {
		t.Fatal("StackTrace empty!")
	}
	fmt.Println(s)
	s = StackTraceSkip(DEPTH, 0)
	if s == "" {
		t.Fatal("StackTrace empty!")
	}
	fmt.Println(s)
}
