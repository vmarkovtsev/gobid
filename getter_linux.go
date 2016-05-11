/*
MIT License

Copyright (c) 2016 Markovtsev Vadim / Mail.Ru Group

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package gobid

import (
	"debug/elf"
	"errors"
	"io"
)

// build with -ldflags "-B 0xdeadbeef"

func GetBuildIDFromFile(exe io.ReaderAt) ([]byte, error) {
  elffile, err := elf.NewFile(exe)
	if err != nil {
		return nil, err
	}
	bidsec := elffile.Section(".note.gnu.build-id")
	if bidsec == nil {
		return nil, errors.New("section .note.gnu.build-id was not found")
	}
	buildid, err := bidsec.Data()
	if err != nil {
		return nil, err
	}
	// first 16 bytes are service data and should be discarded
	return buildid[16:], nil
}
