// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package stacktrace

import (
	"bufio"
	"bytes"
	"os"
)

func Writeline(
	buf *bytes.Buffer,
	file string,
	line int,
) {
	if buf == nil {
		return
	}
	// prints SINGLE LINE of source code ONLY!
	// if the source code spreads multiple lines it won't be printed fully
	// this also ONLY works if the source code exists in the current folder
	f, err := os.Open(file)
	if err != nil {
		// source code doesn't exist
		return
	}
	r := bufio.NewReader(f)
	var bs []byte
	for i := 0; i <= line; i++ {
		for {
			ln, isPrefix, err := r.ReadLine()
			if err != nil {
				// ReadLine either returns a non-nil line or it returns an error, never both.
				break
			}
			if i == line {
				bs = append(bs, ln...)
			}
			if !isPrefix {
				// If the line was too long for the buffer then isPrefix is set and the beginning of the line is returned.
				// The rest of the line will be returned from future calls.
				// isPrefix will be false when returning the last fragment of the line.
				break
			}
		}
	}
	if len(bs) > 0 {
		// found
		buf.WriteString("\t\t")
		buf.Write(bytes.TrimSpace(bs))
		buf.WriteString("\n")
	}
	return
}
