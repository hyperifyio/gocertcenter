// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package managers

import (
	"os"
)

// FileManager wraps up operations to file system for easier testing
// See fsutils for higher level utilities.
type File struct {
	file *os.File
}

// Close wraps up a call to f.file.Close
func (f *File) Close() error {
	return f.file.Close()
}

// Name wraps up a call to f.file.Name
func (f *File) Name() string {
	return f.file.Name()
}

// Write wraps up a call to f.file.Write
func (f *File) Write(b []byte) (int, error) {
	return f.file.Write(b)
}

func NewFile(file *os.File) *File {
	return &File{file}
}

var _ IFile = (*File)(nil)
