// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package managers

import (
	"os"
	"path/filepath"
)

// FileManager wraps up operations to file system for easier testing
// See fsutils for higher level utilities.
type FileManager struct{}

// ReadFile wraps up a call to os.ReadFile
func (f *FileManager) ReadFile(fileName string) ([]byte, error) {
	return os.ReadFile(filepath.Clean(fileName))
}

// MkdirAll wraps up a call to os.MkdirAll
func (f *FileManager) MkdirAll(dir string, dirPerms os.FileMode) error {
	return os.MkdirAll(dir, dirPerms)
}

// CreateTemp wraps up a call to os.CreateTemp
func (f *FileManager) CreateTemp(dir, pattern string) (IFile, error) {
	file, err := os.CreateTemp(dir, pattern)
	return NewFile(file), err
}

// Remove wraps up a call to os.Remove
func (f *FileManager) Remove(name string) error {
	return os.Remove(name)
}

// Rename wraps up a call to os.Rename
func (f *FileManager) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

// Chmod wraps up a call to os.Chmod
func (f *FileManager) Chmod(file string, mode os.FileMode) error {
	return os.Chmod(file, mode)
}

func NewFileManager() FileManager {
	return FileManager{}
}

var _ IFileManager = (*FileManager)(nil)
