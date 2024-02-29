// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package managers

import (
	"os"
	"path/filepath"
)

// OSFileManager wraps up operations to file system for easier testing
// See fsutils for higher level utilities.
type OSFileManager struct{}

// ReadFile wraps up a call to os.ReadFile
func (f *OSFileManager) ReadFile(fileName string) ([]byte, error) {
	return os.ReadFile(filepath.Clean(fileName))
}

// MkdirAll wraps up a call to os.MkdirAll
func (f *OSFileManager) MkdirAll(dir string, dirPerms os.FileMode) error {
	return os.MkdirAll(dir, dirPerms)
}

// CreateTemp wraps up a call to os.CreateTemp
func (f *OSFileManager) CreateTemp(dir, pattern string) (File, error) {
	file, err := os.CreateTemp(dir, pattern)
	return NewFile(file), err
}

// Remove wraps up a call to os.Remove
func (f *OSFileManager) Remove(name string) error {
	return os.Remove(name)
}

// Rename wraps up a call to os.Rename
func (f *OSFileManager) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

// Chmod wraps up a call to os.Chmod
func (f *OSFileManager) Chmod(file string, mode os.FileMode) error {
	return os.Chmod(file, mode)
}

func NewFileManager() *OSFileManager {
	return &OSFileManager{}
}

var _ FileManager = (*OSFileManager)(nil)
