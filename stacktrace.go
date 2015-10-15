// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package stacktrace

import (
	"bytes"
	"fmt"
	"runtime"
)

const (
	DEPTH = 20
	SKIP  = 2
)

func StackTrace(
	depth int,
) string {
	return StackTraceSkip(depth, SKIP)
}
func StackTraceSkip(
	depth int,
	skip int,
) string {
	if depth < 1 {
		depth = DEPTH
	}
	if skip < 0 {
		skip = 0
	}
	st := make([]uintptr, depth)
	// runtime.Callers(skip, []program_counter) int
	// skip: the number of stack frames to skip before recording
	// 0 identifying the frame for Callers itself
	// 1 identifying the caller of Callers
	// program_counter: fills the slice with function invocations on the calling goroutine's stack
	// returns: the number of written entries
	st = st[:runtime.Callers(skip, st[:])]
	if len(st) == 0 {
		return ""
	}
	buf := bytes.NewBuffer(nil)
	for _, pc := range st {
		if pc == 0 {
			continue
		}
		if f := runtime.FuncForPC(pc); f != nil {
			name := f.Name()
			file, line := f.FileLine(pc)
			// path/to/stacktrace.StackTrace
			//   /full/path/to/stacktrace.go:13 +0xABC123
			line-- // line must be decremented
			buf.WriteString(fmt.Sprintf("%s\n\t%s:%d +0x%x\n", name, file, line, pc))
			if file == "" || line < 0 {
				// empty line or negative number
				continue
			}
			//      Writeline(buf, file, line)
			Writeline(buf, file, line)
		}
	}
	return buf.String()
}
